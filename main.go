package main

import (
	"fmt"
)

func main() {
	port := 3000
	_, err := startWebServer(port, 2)
	fmt.Println(err)
}

func startWebServer(port int, numberOfRetries int) (int, error) {
	fmt.Println("Starting server ...")
	fmt.Println("Server started on port", port)
	fmt.Println("Number of retries", numberOfRetries)

	return port, nil
}
