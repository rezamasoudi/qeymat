package api

import (
	"fmt"
	"net/http"
	"qeymat/api/handlers"
	"qeymat/config"
	"time"
)

func Run(config config.Api) error {

	mux := http.NewServeMux()

	// handlers
	mux.HandleFunc("POST /auth/register", handlers.Register)

	httpServer := http.Server{
		Addr:         fmt.Sprintf(":%d", config.Port),
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
		Handler:      mux,
	}

	fmt.Printf("server stasrted: %s \n", config.Url)
	err := httpServer.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}
