package main

import (
	"github.com/innv8/ambulance-hail-api/controllers"
	"github.com/joho/godotenv"
)

var server = controllers.Server{}

func main() {
	_ = godotenv.Load()
	server.Init()
	server.Run()

}

