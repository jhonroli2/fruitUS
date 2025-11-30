package main

import (
    "log"
    "net/http"

    "fruitus/internal/mora/httpmora"
    "fruitus/pkg/config"
)

func main() {
    _, _ = config.LoadConfig()
    router := httpmora.NewRouter()
    log.Println("mora service starting on :8081")
    if err := http.ListenAndServe(":8081", router); err != nil {
        log.Fatal(err)
    }
}
