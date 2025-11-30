package main

import (
    "log"
    "net/http"

    "fruitus/internal/backoffice/httpbo"
)

func main() {
    router := httpbo.NewRouter()
    log.Println("backoffice service starting on :8082")
    if err := http.ListenAndServe(":8082", router); err != nil {
        log.Fatal(err)
    }
}
