package main

import (
	"github.com/akshanshgusain/Video-Calling-App/internal/server"
	"log"
)

func main() {
	if err := server.Run(); err != nil {
		log.Fatal(err.Error())
	}
}
