package main

import (
	"github.com/front-go/gateway/internal/api"
	"github.com/front-go/gateway/internal/service"
	"log"
	"net/http"
)

func main() {
	srv := service.NewService()

	handler := api.NewHandler(srv)

	http.HandleFunc("/", handler.Handle)
	http.HandleFunc("/test", handler.Handle)
	http.HandleFunc("/{q}", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusNotFound)
	})

	err := http.ListenAndServe(":3131", nil)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
