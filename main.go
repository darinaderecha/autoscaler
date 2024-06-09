package main

import (
	"scaler/api"
)

func main() {
	server := api.NewServer()

	if err := server.Start(":8080"); err != nil {
		panic(err)
	}

}
