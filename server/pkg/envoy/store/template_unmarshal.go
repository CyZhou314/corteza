package store

import (
	"github.com/cyzhou314/corteza/server/pkg/envoy"
	"github.com/cyzhou314/corteza/server/pkg/envoy/resource"
	"github.com/cyzhou314/corteza/server/system/types"
)

func newTemplate(t *types.Template) *template {
	return &template{
		t: t,
	}
}

func (u *template) MarshalEnvoy() ([]resource.Interface, error) {
	return envoy.CollectNodes(
		resource.NewTemplate(u.t),
	)
}
