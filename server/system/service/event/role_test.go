package event

import (
	"github.com/cyzhou314/corteza/server/pkg/eventbus"
	"github.com/cyzhou314/corteza/server/system/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoleMatching(t *testing.T) {
	var (
		a   = assert.New(t)
		res = &roleBase{
			role: &types.Role{Handle: "admin"},
		}

		cRol = eventbus.MustMakeConstraint("role", "eq", "admin")
	)

	a.True(res.Match(cRol))
}
