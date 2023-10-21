package main

import (
	"fmt"

	api "github.com/artisademi/go-bank/pkg/api"
)

func main() {
	fmt.Println()
	server := api.NewAPIServer(":8000")
	fmt.Println("Server running at port :8000")
	server.Run()
}
