package store

import (
	"github.com/cyzhou314/corteza/server/pkg/envoy"
	"github.com/cyzhou314/corteza/server/pkg/envoy/resource"
	"github.com/cyzhou314/corteza/server/system/types"
)

func newRole(rl *types.Role) *role {
	return &role{
		rl: rl,
	}
}

func (rl *role) MarshalEnvoy() ([]resource.Interface, error) {
	return envoy.CollectNodes(
		resource.NewRole(rl.rl),
	)
}
