package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
	defer cancel()

	fmt.Println("Mengirim HTTP request...")

	req, err := http.NewRequestWithContext(ctx, "GET", "https://httpbin.org/delay/5", nil)
	if err != nil {
		fmt.Println("Error creating request: ", err)
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request: ", err)
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	fmt.Println("Response status: ", string(body))
}