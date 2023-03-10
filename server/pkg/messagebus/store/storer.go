package store

import (
	"github.com/cyzhou314/corteza/server/pkg/messagebus/types"
)

type (
	Storer interface {
		SetStore(types.QueueStorer)
		GetStore() types.QueueStorer
	}
)
