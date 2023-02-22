package handlers

import (
	"fmt"
	"github.com/cortezaproject/corteza/server/auth/request"
	"github.com/cortezaproject/corteza/server/system/service"
	"github.com/cortezaproject/corteza/server/system/types"
	"go.uber.org/zap"
	"net/http"
)

func (h *AuthHandlers) profileForm(req *request.AuthReq) (err error) {
	req.Template = TmplProfile

	var (
		preferredLanguage string

		u         = req.AuthUser.User
		avatarUrl string
	)

	if langList := h.Locale.LocalizedList(req.Context()); len(langList) > 0 {
		req.Data["languages"] = langList

		preferredLanguage = langList[0].Tag.String()
		if u.Meta != nil && u.Meta.PreferredLanguage != "" {
			preferredLanguage = u.Meta.PreferredLanguage
		}
	}

	if u.Meta.AvatarID != 0 {
		avatarUrl = fmt.Sprintf("/api/system/attachment/avatar/%d/original/profile-photo-avatar", u.Meta.AvatarID)
	}

	if form := req.PopKV(); len(form) > 0 {
		req.Data["form"] = form
	} else {
		req.Data["form"] = map[string]string{
			"email":             u.Email,
			"handle":            u.Handle,
			"name":              u.Name,
			"preferredLanguage": preferredLanguage,
			"avatarUrl":         avatarUrl,
			"avatarInitial":     u.Meta.AvatarInitials,
			"initialTextColor":  u.Meta.AvatarInitialsTextColor,
			"initialBgColor":    u.Meta.AvatarInitialsBgColor,
		}
	}

	req.Data["emailConfirmationRequired"] = !u.EmailConfirmed && h.Settings.EmailConfirmationRequired
	req.Data["avatarEnabled"] = h.Settings.ProfilePhoto.EnabledAvatar
	req.Data["initialsEnabled"] = h.Settings.ProfilePhoto.EnabledInitials
	return nil
}

func (h *AuthHandlers) profileProc(req *request.AuthReq) error {
	req.RedirectTo = GetLinks().Profile
	req.SetKV(nil)

	var (
		u   = req.AuthUser.User
		att *types.Attachment
	)

	u.Handle = req.Request.PostFormValue("handle")
	u.Name = req.Request.PostFormValue("name")

	if pl := req.Request.PostFormValue("preferredLanguage"); pl != "" {
		if u.Meta == nil {
			u.Meta = &types.UserMeta{}
		}

		u.Meta.PreferredLanguage = pl
	}

	if req.Request.PostFormValue("avatar-delete") == "avatar-delete" {
		var (
			u = req.AuthUser.User
		)

		if err := h.Attachment.DeleteByID(req.Context(), u.Meta.AvatarID); err != nil {
			return err
		}

		u.Meta.AvatarID = 0
	} else {
		// get the file from the form
		file, header, err := req.Request.FormFile("avatar")
		if err != nil && err != http.ErrMissingFile {
			return err
		}

		if file != nil {
			defer file.Close()

			labels := map[string]string{
				"key": "profile-photo-avatar",
			}

			if u.Meta.AvatarID != 0 {
				if err = h.Attachment.DeleteByID(req.Context(), u.Meta.AvatarID); err != nil {
					return err
				}
			}

			att, err = h.Attachment.CreateAuthAttachment(
				req.Context(),
				header.Filename,
				header.Size,
				file,
				labels,
			)

			if err != nil {
				if service.AttachmentErrInvalidAvatarFileType().Is(err) || service.AttachmentErrInvalidAvatarFileSize().Is(err) {
					req.SetKV(map[string]string{
						"error": err.Error(),
					})
				}
				return err
			}

			u.Meta.AvatarID = att.ID
		}
	}

	// Assign initials
	u.Meta.AvatarInitials = req.Request.PostFormValue("avatar-initials")
	u.Meta.AvatarInitialsTextColor = req.Request.PostFormValue("initial-color")
	u.Meta.AvatarInitialsBgColor = req.Request.PostFormValue("initial-bg")

	// a little workaround to inject current user as authenticated identity into context
	// this way user service will pass us through.
	user, err := h.UserService.Update(req.Context(), u)

	if err == nil {
		err = h.AuthService.LoadRoleMemberships(req.Context(), user)
	}

	if err == nil {
		req.AuthUser.User = user
		req.AuthUser.Save(req.Session)

		t := translator(req, "auth")
		req.NewAlerts = append(req.NewAlerts, request.Alert{
			Type: "primary",
			Text: t("profile.alerts.profile-updated"),
		})

		req.RedirectTo = GetLinks().Profile
		return nil
	}

	switch {
	case
		service.UserErrInvalidID().Is(err),
		service.UserErrInvalidHandle().Is(err),
		service.UserErrInvalidEmail().Is(err),
		service.UserErrHandleNotUnique().Is(err),
		service.UserErrNotAllowedToUpdate().Is(err):
		req.SetKV(map[string]string{
			"error":  err.Error(),
			"email":  u.Email,
			"handle": u.Handle,
			"name":   u.Name,
		})

		t := translator(req, "auth")
		req.NewAlerts = append(req.NewAlerts, request.Alert{
			Type: "danger",
			Text: t("profile.alerts.profile-update-fail"),
		})

		h.Log.Warn("handled error", zap.Error(err))
		return nil

	default:
		h.Log.Error("unhandled error", zap.Error(err))
		return err
	}
}
