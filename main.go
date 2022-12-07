package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
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

	fmt.Fprintf(w, "Hello, Open Telekom Cloud!\n\nHostname: %s\nTime: %v\nEnvironment: %s\n\n", hostname, time.Now().Local(), env)
}

func startServer(port int) {
	log.Printf("Listening and Serve at :%d", port)

	http.HandleFunc("/", index)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

func main() {
	httpPort, ok := os.LookupEnv("HTTP_PORT")
	if !ok {
		httpPort = "8080"
	}

	if port, err := strconv.Atoi(httpPort); err == nil {
		startServer(port)
	}
}
