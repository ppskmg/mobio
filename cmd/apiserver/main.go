package main

import (
	"log"
	"mobio/internal/app/apiserver"
)

func main() {
	config := apiserver.NewConfig()
	if err := apiserver.Start(config, true); err != nil {
		log.Fatal(err)
	}
}
