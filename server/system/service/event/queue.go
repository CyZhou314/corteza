package event

import (
	"github.com/cyzhou314/corteza/server/pkg/eventbus"
	"github.com/cyzhou314/corteza/server/system/types"
)

var _ = eventbus.ConstraintMaker

// Match returns false if given conditions do not match event & resource internals
func (res queueBase) Match(c eventbus.ConstraintMatcher) bool {
	return queueMatch(res.payload, c)
}

func queueMatch(r *types.QueueMessage, c eventbus.ConstraintMatcher) bool {
	switch c.Name() {
	case "payload.queue":
		return c.Match(r.Queue)
	}

	return false
}
