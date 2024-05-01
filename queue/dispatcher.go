package queue

import (
	"encoding/json"
	"log"

	"github.com/hibiken/asynq"
)

func Dispatch(t BaseTask, options ...asynq.Option) error {
	payload, err := json.Marshal(t)
	if err != nil {
		return err
	}
	task := asynq.NewTask(t.GetEventName(), payload, options...)
	info, err := GetClient().Enqueue(task)
	if err != nil {
		return err
	}
	log.Printf("enqueued task: id=%s queue=%s", info.ID, info.Queue)
	return nil
}
