package queue

import (
	"context"

	"github.com/hibiken/asynq"
)

type BaseTask interface {
	ProcessTask(context.Context, *asynq.Task) error
	GetEventName() string
}
