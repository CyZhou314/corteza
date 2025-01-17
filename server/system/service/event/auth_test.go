package event

import (
	"github.com/cyzhou314/corteza/server/pkg/eventbus"
	"github.com/cyzhou314/corteza/server/system/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuthMatching(t *testing.T) {
	var (
		a   = assert.New(t)
		res = &authBase{
			user: &types.User{Handle: "user"},
		}

		cUsr = eventbus.MustMakeConstraint("user", "eq", "user")
	)

	a.True(res.Match(cUsr))
}
