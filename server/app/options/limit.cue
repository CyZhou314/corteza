package options

import (
	"github.com/cyzhou314/corteza/server/codegen/schema"
)

limit: schema.#optionsGroup & {
	handle: "limit"
	options: {
		system_users: {
			type:        "int"
			description: "Maximum number of valid (not deleted, not suspended) users"
		}
	}
	title: "Limits"
}
