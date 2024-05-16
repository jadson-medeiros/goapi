package main

import "github.com/jadson-medeiros/goapi/configs"

func main() {
	config, _ := configs.LoadConfig(".")
	println(config.DBDriver)
}