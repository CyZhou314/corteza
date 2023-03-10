package system

import (
	"context"
	discoveryType "github.com/cyzhou314/corteza/server/pkg/discovery/types"
	"github.com/cyzhou314/corteza/server/pkg/id"
	"github.com/cyzhou314/corteza/server/store"
	"github.com/cyzhou314/corteza/server/system/service"
	"testing"
)

func (h helper) clearActivityLog() {
	h.noError(store.TruncateResourceActivitys(context.Background(), service.DefaultStore))
}

func (h helper) repoMakeActivityLog() *discoveryType.ResourceActivity {
	var res = &discoveryType.ResourceActivity{
		ID:             id.Next(),
		ResourceID:     id.Next(),
		ResourceType:   "compose:record",
		ResourceAction: "create",
	}

	h.a.NoError(store.CreateResourceActivity(context.Background(), service.DefaultStore, res))

	return res
}

func TestCreateActivityLog(t *testing.T) {
	h := newHelper(t)
	h.clearActionLog()

	h.repoMakeActivityLog()
}
