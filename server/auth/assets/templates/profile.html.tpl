{{ template "inc_header.html.tpl" set . "activeNav" "profile" }}
<div class="card-body p-0">
	<h4 class="card-title p-3 border-bottom">{{ tr "profile.template.title" }}</h4>

	<form
		method="POST"
		action="{{ links.Profile }}"
        enctype="multipart/form-data"
		class="p-3"
	>
		{{ .csrfField }}

		{{ if .form.error }}
		<div
            data-test-id="error"
            class="text-danger mb-4 font-weight-bold"
            role="alert"
        >
			{{ .form.error }}
		</div>
		{{ end }}

        <div class="mb-3">
            <label for="profileFormHandle">{{ tr "profile.template.form.avatar.label" }}</label>
            <div class="d-block">
                <div class="d-flex justify-content-between">
                    <div>
                        <img style="height: 4rem; width: 4rem;" class="rounded-circle" src="{{ .form.avatarUrl }}" alt="Profile Photo">
                    </div>
                    <div>
                        <button
                            name="avatar-delete"
                            value="avatar-delete"
                            class="btn btn-danger"
                        >
                        {{ tr "profile.template.form.avatar.delete" }}
                        </button>

                        <label for="avatar" class="p-2 ml-3 bg-light text-dark rounded">
                        {{ tr "profile.template.form.avatar.update" }}
                        </label>
                        <input id="avatar" name="avatar" type="file" class="sr-only" accept="image/*">
                    </div>
                </div>
            </div>
        </div>


        <div class="mb-3">
            <label for="avatarInitials">{{ tr "profile.template.form.avatar-initial.label" }}</label>
            <div class="d-block">
                <div class="d-flex justify-content-start align-items-center">
                    <div>
                        <input class="col-10 form-control" type="text" id="avatarInitials" name="avatar-initials" value="{{ .form.avatarInitial }}">
                    </div>

                    <div>
                        <div
                            style="height: 4rem; width: 4rem; color: {{ .form.initialTextColor }}; background-color: {{ .form.initialBgColor }} ;" 
                            class="d-flex justify-content-center align-items-center rounded-circle"
                        >
                            <span style="font-size: 1.5rem; line-height: 2rem; letter-spacing: 0.05em;">{{ .form.avatarInitial }}</span>
                        </div>
                    </div>
                </div>
            </div>
        </div>

        <div class="mb-3">
            <label for="initialColor">{{ tr "profile.template.form.avatar-initial.color" }}</label>
            <input type="color" id="initialColor" class="col-6 form-control" value="{{ .form.initialTextColor }}" name="initial-color">
        </div>

        <div class="mb-3">    
            <label for="customColor">{{ tr "profile.template.form.avatar-initial.background-color" }}</label>
            <input type="color" id="customColor" class="col-6 form-control" value="{{ .form.initialBgColor }}" name="initial-bg">
        </div>

        <div class="mb-3">
            <label for="profileFormEmail">{{ tr "profile.template.form.email.label" }}</label>
            <input
                data-test-id="input-email"
                type="email"
                class="form-control"
                name="email"
                id="profileFormEmail"
                placeholder="email@domain.ltd"
                autocomplete="username"
                readonly
                value="{{ .form.email }}"
                aria-label="{{ tr "profile.template.form.email.label" }}"
            >
            <div>
                {{ if .emailConfirmationRequired }}
                <div class="form-text text-danger">
                	{{ tr "profile.template.form.email.resend-confirmation-link" "link" links.PendingEmailConfirmation }}
                </div>
                {{ end }}
            </div>
        </div>

		<div class="mb-3">
			<label for="profileFormName">{{ tr "profile.template.form.name.label" }}</label>
            <input
                data-test-id="input-name"
                type="text"
                class="form-control"
                name="name"
                id="profileFormName"
                placeholder="{{ tr "profile.template.form.name.placeholder" }}"
                value="{{ .form.name }}"
                autocomplete="name"
                aria-label="{{ tr "profile.template.form.name.label" }}"
            >
		</div>

		<div class="mb-3">
			<label for="profileFormHandle">{{ tr "profile.template.form.handle.label" }}</label>
            <input
                data-test-id="input-handle"
                type="text"
                class="form-control handle-mask"
                name="handle"
                id="profileFormHandle"
                placeholder="{{ tr "profile.template.form.handle.placeholder" }}"
                value="{{ .form.handle }}"
                autocomplete="handle"
                aria-label="{{ tr "profile.template.form.handle.label" }}"
            >
		</div>


		<div class="mb-3">
			<label for="profileFormPreferredLanguage">{{ tr "profile.template.form.preferred-language.label" }}</label>
			<select
                data-test-id="select-language"
                class="form-control"
				name="preferredLanguage"
                id="profileFormPreferredLanguage"
                aria-label="{{ tr "profile.template.form.preferred-language.label" }}"
                value="{{ .form.preferredLanguage }}"
			>
			{{ $prefLang := .form.preferredLanguage }}
			{{ range .languages }}
				<option
					value="{{ .Tag }}"
					{{ if eq $prefLang .Tag.String }}selected{{ end }}
				>
					{{ .LocalizedName }} ({{ .Name }})
				</option>
			{{ end }}
			</select>
		</div>

        <div>
            <button
                data-test-id="button-submit"
                type="submit"
                class="btn btn-primary btn-block btn-lg"
            >
                {{ tr "profile.template.form.buttons.submit" }}
            </button>
        </div>
	</form>
</div>
{{ template "inc_footer.html.tpl" . }}
