package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <endpoint>")
		fmt.Println("Example: go run main.go /comment")
		return
	}

	endpoint := os.Args[1]
	baseURL := "http://localhost:3000" // Juice Shop URL

	// Open the payloads file
	file, err := os.Open("payloads.txt")
	if err != nil {
		fmt.Println("Error opening payloads file:", err)
		return
	}
	defer file.Close()

	fmt.Printf("Testing Stored XSS on endpoint: %s%s\n\n", baseURL, endpoint)

	// Read payloads and test each one
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		payload := scanner.Text()
		encodedPayload := strings.ReplaceAll(payload, " ", "%20")
		testURL := baseURL + endpoint + encodedPayload

		// Send HTTP GET request
		resp, err := http.Get(testURL)
		if err != nil {
			fmt.Printf("Error testing payload: %s\n", err)
			continue
		}
		defer resp.Body.Close()

		fmt.Printf("Tested: %s | Status: %d\n", testURL, resp.StatusCode)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading payloads file:", err)
	}
}
