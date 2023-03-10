package locale

import (
	"github.com/cyzhou314/corteza/server/pkg/xss"
)

func SanitizeMessage(in string) string {
	return xss.RichText(in)
}
