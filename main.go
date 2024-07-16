package main

import "os"

func main() {
	port := os.Getenv("port")
	if port == "" {
		port = "8080"
	}

}
