package yaml

import (
	"github.com/cyzhou314/corteza/server/pkg/envoy/resource"
	"github.com/cyzhou314/corteza/server/system/types"
)

type (
	template struct {
		res *types.Template
		ts  *resource.Timestamps
		us  *resource.Userstamps

		envoyConfig   *resource.EnvoyConfig
		encoderConfig *EncoderConfig

		rbac rbacRuleSet
	}
	templateSet []*template
)

func (nn templateSet) configureEncoder(cfg *EncoderConfig) {
	for _, n := range nn {
		n.encoderConfig = cfg
	}
}
