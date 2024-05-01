package queue

import (
	"fmt"
	"qeymat/config"
	"qeymat/queue/tasks"

	"github.com/hibiken/asynq"
)

var client *asynq.Client

func RunServer(config config.Redis) error {
	srv := asynq.NewServer(
		asynq.RedisClientOpt{
			Addr:     fmt.Sprintf("%s:%d", config.Host, config.Port),
			Username: config.User,
			Password: config.Password,
			DB:       config.Database,
		},
		asynq.Config{
			// Specify how many concurrent workers to use
			Concurrency: 10,
		},
	)

	// mux maps a type to a handler
	mux := asynq.NewServeMux()

	crwler := &tasks.Crawler{}
	mux.Handle(crwler.GetEventName(), crwler)

	println("asyncq server started")
	return srv.Run(mux)
}

func CreateClient(config config.Redis) {
	client = asynq.NewClient(asynq.RedisClientOpt{
		Addr:     fmt.Sprintf("%s:%d", config.Host, config.Port),
		Username: config.User,
		Password: config.Password,
		DB:       config.Database,
	})
}

func CloseClient() {
	client.Close()
}

func GetClient() *asynq.Client {
	return client
}
