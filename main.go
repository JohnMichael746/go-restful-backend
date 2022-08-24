package main

import (
	"fmt"

	"github.com/AlxPatidar/go-restful-api/app"
	"github.com/AlxPatidar/go-restful-api/config"
)

func main() {
	fmt.Println("Go Mongo Restful API")
	config.Init()
	app.Start()
}
