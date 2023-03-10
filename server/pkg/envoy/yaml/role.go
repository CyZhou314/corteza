package yaml

import (
	"github.com/cyzhou314/corteza/server/pkg/envoy/resource"
	"github.com/cyzhou314/corteza/server/system/types"
)

type (
	role struct {
		res *types.Role
		ts  *resource.Timestamps

		envoyConfig   *resource.EnvoyConfig
		encoderConfig *EncoderConfig

		// module's RBAC rules
		rbac rbacRuleSet
	}
	roleSet []*role
)

func (nn roleSet) configureEncoder(cfg *EncoderConfig) {
	for _, n := range nn {
		n.encoderConfig = cfg
	}
}
