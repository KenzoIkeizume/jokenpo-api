package main

import (
	config "jokenpo-api/config"
	controller "jokenpo-api/controller"

	"flag"
)

func main() {
	flag.Parse()

	config.ReadConfig()
	controller.AppController()
}
