package renderer

import "github.com/cyzhou314/corteza/server/pkg/valuestore"

func envGetter() func(k string) any {
	return func(k string) any {
		return valuestore.Global().Env(k)
	}
}
