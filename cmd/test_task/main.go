package main

import (
	"github.com/yanasirina/umbrella_corp_test_task/internal/pkg/app"
)

func main() {
	server, _ := app.New()
	server.Run()
}
