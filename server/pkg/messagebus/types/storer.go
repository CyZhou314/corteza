package types

import (
	"context"

	"github.com/cyzhou314/corteza/server/pkg/eventbus"
)

type (
	QueueStorer interface {
		SearchQueues(context.Context, QueueFilter) ([]QueueDb, QueueFilter, error)
		CreateQueueMessage(context.Context, QueueMessage) error
		ProcessQueueMessage(context.Context, uint64, QueueMessage) error
	}

	QueueEventBuilder interface {
		CreateQueueEvent(string, []byte) eventbus.Event
	}

	QueueServicer interface {
		QueueStorer
		QueueEventBuilder
	}
)
