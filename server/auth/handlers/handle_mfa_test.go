package handlers

import (
	"context"
	"net/http"
	"net/url"
	"testing"

	"github.com/cyzhou314/corteza/server/auth/request"
	"github.com/cyzhou314/corteza/server/auth/settings"
	"github.com/cyzhou314/corteza/server/system/service"
	"github.com/cyzhou314/corteza/server/system/types"
	"github.com/stretchr/testify/require"
)

func Test_mfaProc(t *testing.T) {
	var (
		user = makeMockUser()

		req = &http.Request{}

		authService  authService
		authHandlers *AuthHandlers
		authReq      *request.AuthReq
	)

	service.CurrentSettings = &types.AppSettings{}

	tcc := []testingExpect{
		{
			name:    "Email: successful login",
			payload: map[string]string(nil),
			alerts:  []request.Alert{{Type: "primary", Text: "mfa.email.resent"}},
			link:    GetLinks().Profile,
			fn: func(_ *settings.Settings) {
				req.Form.Set("action", "verifyEmailOtp")
				req.PostForm.Add("code", "123456")

				authService = &authServiceMocked{
					validateEmailOTP: func(ctx context.Context, code string) (err error) {
						return nil
					},
				}
			},
		},
		{
			name:    "TOTP: successful login",
			payload: map[string]string(nil),
			alerts:  []request.Alert{{Type: "primary", Text: "mfa.topt.valid"}},
			link:    GetLinks().Mfa,
			fn: func(_ *settings.Settings) {
				req.Form.Set("action", "verifyTotp")
				req.PostForm.Add("code", "123456")

				authService = &authServiceMocked{
					validateTOTP: func(ctx context.Context, code string) (err error) {
						return nil
					},
				}
			},
		},
		{
			name:    "Email: disabled",
			payload: map[string]string{"emailOtpError": "multi factor authentication with email OTP is disabled"},
			alerts:  []request.Alert(nil),
			link:    GetLinks().Mfa,
			fn: func(_ *settings.Settings) {
				req.Form.Set("action", "verifyEmailOtp")

				authService = &authServiceMocked{
					validateEmailOTP: func(ctx context.Context, code string) (err error) {
						return service.AuthErrDisabledMFAWithEmailOTP()
					},
				}
			},
		},
		{
			name:    "Email: auth failed for disabled user",
			payload: map[string]string{"emailOtpError": "invalid username and password combination"},
			alerts:  []request.Alert(nil),
			link:    GetLinks().Mfa,
			fn: func(_ *settings.Settings) {
				req.Form.Set("action", "verifyEmailOtp")

				authService = &authServiceMocked{
					validateEmailOTP: func(ctx context.Context, code string) (err error) {
						return service.AuthErrFailedForUnknownUser()
					},
				}
			},
		},
		{
			name:    "Email: invalid token",
			payload: map[string]string{"emailOtpError": "invalid code"},
			alerts:  []request.Alert(nil),
			link:    GetLinks().Mfa,
			fn: func(_ *settings.Settings) {
				req.Form.Set("action", "verifyEmailOtp")
				req.PostForm.Add("code", "token_TOO_LONG")

				authService = &authServiceMocked{
					validateEmailOTP: func(ctx context.Context, code string) (err error) {
						return service.AuthErrInvalidEmailOTP()
					},
				}
			},
		},
		{
			name:    "Email: no token in credentials db",
			payload: map[string]string{"emailOtpError": "invalid code"},
			alerts:  []request.Alert(nil),
			link:    GetLinks().Mfa,
			fn: func(_ *settings.Settings) {
				req.Form.Set("action", "verifyEmailOtp")
				req.PostForm.Add("code", "123456")

				authService = &authServiceMocked{
					validateEmailOTP: func(ctx context.Context, code string) (err error) {
						return service.AuthErrInvalidEmailOTP()
					},
				}
			},
		},
		{
			name:    "TOTP: disabled",
			payload: map[string]string{"totpError": "multi factor authentication with TOTP is disabled"},
			alerts:  []request.Alert(nil),
			link:    GetLinks().Mfa,
			fn: func(_ *settings.Settings) {
				req.Form.Set("action", "verifyTotp")

				authService = &authServiceMocked{
					validateTOTP: func(ctx context.Context, code string) (err error) {
						return service.AuthErrDisabledMFAWithTOTP()
					},
				}
			},
		},
		{
			name:    "TOTP: auth failed for disabled user",
			payload: map[string]string{"totpError": "invalid username and password combination"},
			alerts:  []request.Alert(nil),
			link:    GetLinks().Mfa,
			fn: func(_ *settings.Settings) {
				req.Form.Set("action", "verifyTotp")

				authService = &authServiceMocked{
					validateTOTP: func(ctx context.Context, code string) (err error) {
						return service.AuthErrFailedForUnknownUser()
					},
				}
			},
		},
		{
			name:    "TOTP: invalid token",
			payload: map[string]string{"totpError": "invalid code"},
			alerts:  []request.Alert(nil),
			link:    GetLinks().Mfa,
			fn: func(_ *settings.Settings) {
				req.Form.Set("action", "verifyTotp")
				req.PostForm.Add("code", "token_TOO_LONG")

				authService = &authServiceMocked{
					validateTOTP: func(ctx context.Context, code string) (err error) {
						return service.AuthErrInvalidTOTP()
					},
				}
			},
		},
		{
			name:    "TOTP: no token in credentials db",
			payload: map[string]string{"totpError": "invalid code"},
			alerts:  []request.Alert(nil),
			link:    GetLinks().Mfa,
			fn: func(_ *settings.Settings) {
				req.Form.Set("action", "verifyTotp")
				req.PostForm.Add("code", "123456")

				authService = &authServiceMocked{
					validateTOTP: func(ctx context.Context, code string) (err error) {
						return service.AuthErrInvalidTOTP()
					},
				}
			},
		},
	}

	for _, tc := range tcc {
		t.Run(tc.name, func(t *testing.T) {
			rq := require.New(t)

			// reset from previous
			req.Form = url.Values{}
			req.PostForm = url.Values{}
			user.Meta = &types.UserMeta{}

			authSettings := &settings.Settings{}

			tc.fn(authSettings)

			authHandlers = prepareClientAuthHandlers(authService, authSettings)
			authReq = prepareClientAuthReq(authHandlers, req, user)

			err := authHandlers.mfaProc(authReq)

			rq.NoError(err)
			rq.Equal(tc.payload, authReq.GetKV())
			rq.Equal(tc.alerts, authReq.NewAlerts)
			rq.Equal(tc.link, authReq.RedirectTo)
		})
	}
}
