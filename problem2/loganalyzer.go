package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func analyzeLogs(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	var totalRequests int
	var error404Count int
	ipMap := make(map[string]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		totalRequests++

		fields := strings.Fields(line)
		if len(fields) < 9 {
			continue
		}

		ip := fields[0]
		statusCode := fields[8]

		ipMap[ip]++
		if statusCode == "404" {
			error404Count++
		}
	}

	fmt.Println("Log Analysis Report")
	fmt.Printf("Total Requests: %d\n", totalRequests)
	fmt.Printf("404 Errors:     %d\n", error404Count)
	fmt.Println("Top IP Addresses:")

	count := 0
	for ip, freq := range ipMap {
		fmt.Printf("- %s: %d requests\n", ip, freq)
		count++
		if count >= 3 {
			break
		}
	}
}

func main() {
	logFile := "access.log"
	analyzeLogs(logFile)
}
