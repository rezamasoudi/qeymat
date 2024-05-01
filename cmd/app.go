package main

import (
	"fmt"
	"qeymat/api"
	"qeymat/config"
	mongo "qeymat/database"
	"qeymat/queue"
	"qeymat/queue/tasks"
	"time"

	"github.com/hibiken/asynq"
)

func main() {

	// load config
	err := config.LoadConfig()
	if err != nil {
		panic(err)
	}
	fmt.Println("application config loaded.")

	// create db connection
	err = mongo.CreateClient(config.GetConfig().Mongo)
	if err != nil {
		panic(fmt.Sprintf("eror on connect to mongo database: %v", err))
	}

	// disconnect database on finish appliation
	defer mongo.Disconnect()
	fmt.Println("mongo database connection stablished.")

	go func() {
		err := queue.RunServer(config.GetConfig().Redis)
		if err != nil {
			panic(fmt.Sprintf("error run asynq server: %v", err))
		}
	}()

	// create asyq client
	queue.CreateClient(config.GetConfig().Redis)
	defer queue.CloseClient()
	fmt.Println("asynq client created")

	// hit starting crawl
	queue.Dispatch(&tasks.Crawler{}, asynq.ProcessIn(10*time.Second))

	// start api server
	api.Run(config.GetConfig().Api)

}
