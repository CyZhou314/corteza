package event

import "github.com/cyzhou314/corteza/server/pkg/eventbus"

// Match returns false if given conditions do not match event & resource internals
func (res authBase) Match(c eventbus.ConstraintMatcher) bool {
	return userMatch(res.user, c)
}
