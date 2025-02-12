package main

import (
	"github.com/front-go/gateway/internal/api"
	"github.com/front-go/gateway/internal/client/auth"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

func main() {
	authClient := auth.NewClient()

	handler := api.NewHandler(authClient)

	router := chi.NewRouter()

	api.AttachHandlers(router, handler)

	err := http.ListenAndServe(":3131", router)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
