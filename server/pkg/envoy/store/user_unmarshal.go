package store

import (
	"github.com/cyzhou314/corteza/server/pkg/envoy"
	"github.com/cyzhou314/corteza/server/pkg/envoy/resource"
	"github.com/cyzhou314/corteza/server/system/types"
)

func newUser(u *types.User) *user {
	return &user{
		u: u,
	}
}

func (u *user) MarshalEnvoy() ([]resource.Interface, error) {
	return envoy.CollectNodes(
		resource.NewUser(u.u),
	)
}
