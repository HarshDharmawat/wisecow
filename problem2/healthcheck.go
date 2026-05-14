package main

import (
	"fmt"
	"net/http"
	"time"
)

const targetURL = "https://accuknox.com/"

func checkAppStatus(url string) {
	client := http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		fmt.Printf("[DOWN] Application %s is unavailable: %v\n", url, err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		fmt.Printf("[UP] Application %s is functioning correctly. Status: %d\n", url, resp.StatusCode)
	} else {
		fmt.Printf("[DOWN] Application %s returned status: %d\n", url, resp.StatusCode)
	}
}

func main() {
	fmt.Println("Starting Application Health Check...")
	checkAppStatus(targetURL)
}
