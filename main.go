package main

import (
	"backend/api"
)

func main() {
	server := api.Init()
	err := server.Run(":7001")
	if err != nil {
		panic(err)
	}
}
