package main

import (
	"log"
	"os"
	"time"

	"cron/jobs"
)

func main() {
	url := os.Getenv("PING_URL")
	if url == "" {
		log.Fatal("PING_URL is not set")
	}

	for {
		jobs.Ping(url)
		time.Sleep(10 * time.Minute)
	}
}
