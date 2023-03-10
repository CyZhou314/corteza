package yaml

import (
	"github.com/cyzhou314/corteza/server/compose/types"
	"github.com/cyzhou314/corteza/server/pkg/envoy/resource"
)

type (
	composeRecord struct {
		id     string
		values map[string]string
		ts     *resource.Timestamps
		us     *resource.Userstamps
		config *resource.EnvoyConfig

		cfg *EncoderConfig

		refModule    string
		refNamespace string

		rbac rbacRuleSet
	}
	composeRecordSet []*composeRecord

	composeRecordValues struct {
		rvs types.RecordValueSet
	}
)

func (nn composeRecordSet) configureEncoder(cfg *EncoderConfig) {
	for _, n := range nn {
		n.cfg = cfg
	}
}
