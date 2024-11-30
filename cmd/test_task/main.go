package main

import (
	"github.com/yanasirina/umbrella_corp_test_task/internal/pkg/app"
	"log"
)

func main() {
	server, err := app.New()
	if err != nil {
		log.Fatal(err)
	}

	err = server.Run()
	if err != nil {
		log.Fatal(err)
	}
}
