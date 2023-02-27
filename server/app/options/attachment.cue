package options

import (
	"github.com/cortezaproject/corteza/server/codegen/schema"
)

attachment: schema.#optionsGroup & {
	handle: "attachment"

	options: {
		avatar_max_file_size: {
			type: "int64"
			description:  "Avatar image maximum upload size, default value is 3MB"
		}
		avatar_initials_font_path: {
			defaultValue: "./auth/assets/public/fonts/poppins/Poppins-Regular.ttf"
			description:  "Avatar initials font file path"
			env:          "AVATAR_INITIALS_FONT_PATH"
		}
	}
}
