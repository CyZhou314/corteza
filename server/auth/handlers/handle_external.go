package handlers

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/cyzhou314/corteza/server/auth/external"
	"github.com/cyzhou314/corteza/server/auth/request"
	"github.com/cyzhou314/corteza/server/auth/settings"
	"github.com/cyzhou314/corteza/server/pkg/api"
	"github.com/cyzhou314/corteza/server/pkg/auth"
	"github.com/cyzhou314/corteza/server/system/types"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

func copyProviderToContext(r *http.Request) *http.Request {
	return r.WithContext(context.WithValue(r.Context(), "provider", chi.URLParam(r, "provider")))
}

func (h AuthHandlers) externalInit(w http.ResponseWriter, r *http.Request) {
	r = copyProviderToContext(r)
	h.Log.Info("starting external authentication flow")

	beginUserAuth(w, r, external.NewDefaultExternalHandler())
}

func (h *AuthHandlers) externalCallback(w http.ResponseWriter, r *http.Request) {
	r = copyProviderToContext(r)
	h.Log.Info("external authentication callback")

	if user, err := completeUserAuth(w, r, external.NewDefaultExternalHandler()); err != nil {
		h.Log.Error("failed to complete user auth", zap.Error(err))
		h.handleFailedExternalAuth(w, r, err)
	} else {
		h.handleSuccessfulExternalAuth(w, r, *user)
	}
}

// Handles authentication via external auth providers of
// unknown an user + appending authentication on external providers
// to a current user
func (h *AuthHandlers) handleSuccessfulExternalAuth(w http.ResponseWriter, r *http.Request, cred types.ExternalAuthUser) {
	h.Log.Info("login successful", zap.String("provider", cred.Provider))

	switch cred.Provider {
	case "saml":
		h.handleSuccessfulExternalSAMLAuth(w, r, cred)

	default:
		h.handleSuccessfulExternalGenericAuth(w, r, cred)
	}
}

// handleSuccessfulExternalSAMLAuth is a special handler just for the SAML authentication
func (h *AuthHandlers) handleSuccessfulExternalSAMLAuth(w http.ResponseWriter, r *http.Request, cred types.ExternalAuthUser) {
	var (
		user *types.User
		err  error
		ctx  = r.Context()
	)

	// Try to login/sign-up external user
	if user, err = h.AuthService.External(ctx, cred); err != nil {
		api.Send(w, r, err)
		return
	}

	h.handle(func(req *request.AuthReq) error {
		req.AuthUser = request.NewAuthUser(
			h.Settings,
			user,

			// external logins are never permanent!
			false,
		)

		// auto-complete EmailOTP and TOTP when authenticating via external identity provider
		req.AuthUser.CompleteEmailOTP()
		req.AuthUser.CompleteTOTP()

		req.AuthUser.Save(req.Session)

		handleSuccessfulAuth(req)
		return nil
	})(w, r)
}

// handleSuccessfulExternalGenericAuth is a generic external auth handler
// which can be used with generic auth providers to also support external API integrations
func (h *AuthHandlers) handleSuccessfulExternalGenericAuth(w http.ResponseWriter, r *http.Request, cred types.ExternalAuthUser) {
	var (
		user *types.User
		err  error
		ctx  = r.Context()
	)

	handleErr := func(err error) {
		h.handleFailedExternalAuth(w, r, err)
	}

	// Get the provider config so we can correctly handle the provided values
	p := h.getProviderConfig(cred.Provider, h.Settings.Providers)
	if p == nil {
		handleErr(fmt.Errorf("credentials for provider %s are not registered in the system", cred.Provider))
		return
	}

	// For later, the request's auth user (no big deal if there isn't one)
	au := request.GetAuthUser(h.SessionManager.Get(r))

	// Check if we're using it as an identity provider; if so, use it to authenticate
	// If no providers are specified, assume identity (legacy support)
	if len(p.Usage) == 0 || p.HasUsage(types.ExternalProviderUsageIdentity) {
		// Try to login/sign-up external user
		if user, err = h.AuthService.External(ctx, cred); err != nil {
			handleErr(err)
			return
		}

		au = request.NewAuthUser(
			h.Settings,
			user,

			// external logins are never permanent!
			false,
		)

		// If that's that, cut the flow here
		if !p.HasUsage(types.ExternalProviderUsageAPI) {
			h.handle(func(req *request.AuthReq) error {
				req.AuthUser = au

				// auto-complete EmailOTP and TOTP when authenticating via external identity provider
				req.AuthUser.CompleteEmailOTP()
				req.AuthUser.CompleteTOTP()

				req.AuthUser.Save(req.Session)

				handleSuccessfulAuth(req)
				return nil
			})(w, r)
			return
		}
	}

	// Check if we're using it for an API integration; if so, note the access tokens
	if p.HasUsage(types.ExternalProviderUsageAPI) {
		if au.User == nil {
			handleErr(fmt.Errorf("could not add credentials for user: not authenticated"))
			return
		}
		ctx = auth.SetIdentityToContext(ctx, au.User)

		// Look for existing
		cc, err := h.CredentialsService.List(ctx, au.User.ID)
		if err != nil {
			handleErr(fmt.Errorf("couldn't fetch user credentials: %w", err))
			return
		}

		// Find the existing one
		kind := fmt.Sprintf("access-%s", cred.Provider)
		var current *types.Credential
		for _, c := range cc {
			if c.Kind == kind && c.OwnerID == au.User.ID {
				current = c
				break
			}
		}

		// Update
		if current != nil {
			current.Credentials = cred.AccessToken
			_, err = h.CredentialsService.Update(ctx, current)
			if err != nil {
				handleErr(fmt.Errorf("couldn't update user credentials: %w", err))
				return
			}
		} else {
			_, err = h.CredentialsService.Create(ctx, &types.Credential{
				Label:       fmt.Sprintf("access token %s %s", cred.Provider, au.User.Email),
				Kind:        kind,
				OwnerID:     au.User.ID,
				Credentials: cred.AccessToken,
			})
			if err != nil {
				handleErr(fmt.Errorf("couldn't create user credentials: %w", err))
				return
			}
		}
	}

	h.handle(func(req *request.AuthReq) error {
		req.AuthUser = au

		handleSuccessfulAuth(req)
		return nil
	})(w, r)
}

func (h AuthHandlers) getProviderConfig(handle string, set []settings.Provider) *settings.Provider {
	for _, s := range set {
		if s.Handle == handle {
			return &s
		}
	}

	return nil
}

func (h AuthHandlers) handleFailedExternalAuth(w http.ResponseWriter, _ *http.Request, err error) {
	if strings.Contains(err.Error(), "Error processing your OAuth request: Invalid oauth_verifier parameter") {
		// Just take user through the same loop again
		w.Header().Set("Location", GetLinks().Profile)
		w.WriteHeader(http.StatusSeeOther)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprintf(w, "SSO Error: %v", err.Error())
}

func beginUserAuth(w http.ResponseWriter, r *http.Request, eh external.ExternalAuthHandler) {
	eh.BeginUserAuth(w, r)
}

func completeUserAuth(w http.ResponseWriter, r *http.Request, eh external.ExternalAuthHandler) (u *types.ExternalAuthUser, err error) {
	return eh.CompleteUserAuth(w, r)
}
