package options

import (
	"github.com/cortezaproject/corteza/server/codegen/schema"
)

attachment: schema.#optionsGroup & {
	handle: "attachment"

	options: {
		avatar_max_file_size: {
			type: "int64"
			defaultGoExpr: "3000000"
			description:  "Avatar image maximum upload size, default value is 3MB"
		}
		avatar_initials_font_path: {
			defaultValue: "./auth/assets/public/fonts/poppins/Poppins-Regular.ttf"
			description:  "Avatar initials font file path"
			env:          "AVATAR_INITIALS_FONT_PATH"
		}
		avatar_initials_background_color: {
			defaultValue: "#000000"
			description:  "Avatar initials background color"
			env:          "AVATAR_INITIALS_BACKGROUND_COLOR"
		}
		avatar_initials_color: {
			defaultValue: "#ffffff"
			description:  "Avatar initials text color"
			env:          "AVATAR_INITIALS_COLOR"
		}
	}
}
