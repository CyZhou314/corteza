package options

import (
	"github.com/cyzhou314/corteza/server/codegen/schema"
)

environment: schema.#optionsGroup & {
	handle: "environment"
	options: {
		environment: {
			defaultValue: "production"
			env:          "ENVIRONMENT"
		}
	}
	title: "Environment"
}
