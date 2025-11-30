package main

import (
    "log"
    "net/http"
)

func main() {
    // Initialize application components here

    // Start the HTTP server
    log.Println("Starting FruitUS server...")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatalf("Could not start server: %s\n", err)
    }
}