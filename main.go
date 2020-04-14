package main

import (
	config "jokenpo-api/config"
	controller "jokenpo-api/controller"
)

func main() {
	config.ReadConfig()
	controller.AppController()
}
