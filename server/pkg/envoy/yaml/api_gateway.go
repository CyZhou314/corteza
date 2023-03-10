package yaml

import (
	"github.com/cyzhou314/corteza/server/pkg/envoy/resource"
	"github.com/cyzhou314/corteza/server/system/types"
)

type (
	apiGateway struct {
		res     *types.ApigwRoute
		filters apiGwFilterSet

		ts *resource.Timestamps
		us *resource.Userstamps

		envoyConfig   *resource.EnvoyConfig
		encoderConfig *EncoderConfig

		rbac rbacRuleSet
	}
	apiGatewaySet []*apiGateway

	apiGwFilter struct {
		res *types.ApigwFilter

		ts *resource.Timestamps
		us *resource.Userstamps

		envoyConfig   *resource.EnvoyConfig
		encoderConfig *EncoderConfig
	}
	apiGwFilterSet []*apiGwFilter
)

func (nn apiGatewaySet) configureEncoder(cfg *EncoderConfig) {
	for _, n := range nn {
		n.encoderConfig = cfg
	}
}
