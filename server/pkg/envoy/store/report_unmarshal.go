package store

import (
	"github.com/cyzhou314/corteza/server/pkg/envoy"
	"github.com/cyzhou314/corteza/server/pkg/envoy/resource"
	"github.com/cyzhou314/corteza/server/system/types"
)

func newReport(wf *types.Report, ux *userIndex) *report {
	return &report{
		rp: wf,
		ss: wf.Sources,
		pp: wf.Blocks,

		ux: ux,
	}
}

func (awf *report) MarshalEnvoy() ([]resource.Interface, error) {
	rs := resource.NewReport(awf.rp)
	syncUserStamps(rs.Userstamps(), awf.ux)

	for _, s := range awf.ss {
		rs.AddReportSource(s)
	}

	for _, p := range awf.pp {
		rs.AddReportBlock(p)
	}

	return envoy.CollectNodes(
		rs,
	)
}
