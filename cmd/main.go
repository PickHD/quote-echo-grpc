package main

import (
	"log"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalln("Failed to construct new zap logger:", err)
	}

	if err := godotenv.Load(); err != nil {
		logger.Fatal("Failed to load .env with godotenv", zap.Error(err))
	}
}
