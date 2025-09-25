package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"cron/jobs"
)

func main() {
	url := os.Getenv("PING_URL")
	if url == "" {
		log.Fatal("PING_URL is not set")
	}

	go func() {
		for {
			jobs.Ping(url)
			time.Sleep(10 * time.Minute)
		}
	}()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Ping job is running")
	})
	log.Printf("Listening on port %s...", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
