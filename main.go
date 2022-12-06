package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func index(w http.ResponseWriter, r *http.Request) {
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, "OTC-Hello-Server: %s\nTime: %v", hostname, time.Now().Local())
}

func startServer() {
	http.HandleFunc("/", index)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	startServer()
}
