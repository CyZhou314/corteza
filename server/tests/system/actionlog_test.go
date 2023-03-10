package system

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/cyzhou314/corteza/server/pkg/actionlog"
	"github.com/cyzhou314/corteza/server/pkg/id"
	"github.com/cyzhou314/corteza/server/store"
	"github.com/cyzhou314/corteza/server/system/service"
	"github.com/cyzhou314/corteza/server/system/types"
	"github.com/cyzhou314/corteza/server/tests/helpers"
	jsonpath "github.com/steinfletcher/apitest-jsonpath"
)

func (h helper) clearActionLog() {
	h.noError(store.TruncateActionlogs(context.Background(), service.DefaultStore))
}

func (h helper) repoMakeActionLog() *actionlog.Action {
	var res = &actionlog.Action{
		ID:        id.Next(),
		Timestamp: time.Now(),
		ActorID:   id.Next(),
		Resource:  types.ComponentRbacResource(),
		Action:    "lookup",
	}

	h.a.NoError(store.CreateActionlog(context.Background(), service.DefaultStore, res))

	return res
}

func TestActionLogList(t *testing.T) {
	h := newHelper(t)
	h.clearActionLog()

	helpers.AllowMe(h, types.ComponentRbacResource(), "action-log.read")

	h.repoMakeActionLog()

	h.apiInit().
		Get("/actionlog/").
		Header("Accept", "application/json").
		Expect(t).
		Status(http.StatusOK).
		Assert(helpers.AssertNoErrors).
		Assert(jsonpath.Len(`$.response.set`, 1)).
		End()
}

func TestActionLogListForbidden(t *testing.T) {
	h := newHelper(t)
	h.clearActionLog()

	h.repoMakeActionLog()

	h.apiInit().
		Get("/actionlog/").
		Header("Accept", "application/json").
		Expect(t).
		Status(http.StatusOK).
		Assert(helpers.AssertError("access_denied")).
		End()
}
