package main

import (
	"file-structure/internal/api"
	"file-structure/internal/service"
	"fmt"
)

func main() {
	srv := service.NewService()

	handler := api.NewHandler(srv)

	res := handler.Handle(10, 4)
	fmt.Println(res)
}
