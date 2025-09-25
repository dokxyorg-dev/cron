package main

import (
	"log"
	"os"
	"time"

	"cron/jobs"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	url := os.Getenv("PING_URL")
	if url == "" {
		log.Fatal("PING_URL is not set")
	}

	for {
		jobs.Ping(url)
		time.Sleep(10 * time.Minute)
	}
}
