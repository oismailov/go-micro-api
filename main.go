package main

import (
	"fmt"
	r "github.com/oismailov/go-micro-api/router"
)

func main() {
	router := r.GetRouter()

	err := router.Run(r.GetPort())

	if err != nil {
		fmt.Println("unable to start service")
	}
}
