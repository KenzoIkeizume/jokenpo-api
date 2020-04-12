package main

import (
	"jokenpo-api/config"
	controller "jokenpo-api/controllers"
)

func main() {
	config.ReadConfig()
	controller.AppController()
}
