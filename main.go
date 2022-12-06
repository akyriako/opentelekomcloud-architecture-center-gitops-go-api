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

	env, ok := os.LookupEnv("ENV_NAME")
	if !ok {
		env = "UNKNOWN"
	}

	fmt.Fprintf(w, "Hello, Open Telekom Cloud!\n\nHostname: %s\nTime: %v\nEnvironment: %s", hostname, time.Now().Local(), env)
}

func startServer() {
	http.HandleFunc("/", index)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	startServer()
}
