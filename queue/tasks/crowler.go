package tasks

import (
	"context"

	"github.com/hibiken/asynq"
)

type Crawler struct {
}

func (crawler *Crawler) GetEventName() string {
	return "start:crwal"
}

func (crawler *Crawler) ProcessTask(ctx context.Context, task *asynq.Task) error {
	println("processing crawler task")
	return nil
}
