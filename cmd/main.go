package main

import (
	"log"

	"github.com/joho/godotenv"

	"github.com/PickHD/quote-echo-grpc/internal/wire"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Failed to load .env with godotenv", err)
	}

	server, err := wire.InitializeServer("quote")
	if err != nil {
		log.Fatal(err)
	}

	server.Serve()
}
