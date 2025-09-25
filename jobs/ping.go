package jobs

import (
	"fmt"
	"net/http"
)

func Ping(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Ping failed:", err)
		return
	}
	defer resp.Body.Close()
	fmt.Println("Ping success, status:", resp.Status)
}
