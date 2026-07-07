package main

import (
	"nursing-knowledge-ai/cmd/config"
	"nursing-knowledge-ai/infrastructure/http/server"
)

func main() {
	router := server.StartServer()

	router.Run(config.Port)
}
