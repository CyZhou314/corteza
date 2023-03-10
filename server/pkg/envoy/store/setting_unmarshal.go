package store

import (
	"github.com/cyzhou314/corteza/server/pkg/envoy"
	"github.com/cyzhou314/corteza/server/pkg/envoy/resource"
	"github.com/cyzhou314/corteza/server/system/types"
)

func newSetting(res *types.SettingValue, ux *userIndex) *setting {
	return &setting{
		st: res,

		ux: ux,
	}
}

func (st *setting) MarshalEnvoy() ([]resource.Interface, error) {
	rs := resource.NewSetting(st.st)
	syncUserStamps(rs.Userstamps(), st.ux)

	return envoy.CollectNodes(
		rs,
	)
}
