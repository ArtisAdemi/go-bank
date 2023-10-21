package main

import (
	"fmt"
	"log"

	api "github.com/artisademi/go-bank/pkg/api"
	storage "github.com/artisademi/go-bank/pkg/storage"
)

func main() {
	store, err := storage.NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	server := api.NewAPIServer(":8000", store)
	fmt.Println("Server running at port :8000")
	server.Run()
}
