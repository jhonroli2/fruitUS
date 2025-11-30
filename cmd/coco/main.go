package main

import (
    "log"

    "fruitus/internal/coco/worker"
)

func main() {
    // Arrancamos el worker del servicio coco
    go worker.StartConsumer()
    log.Println("coco service running (worker started)")
    select {} // keep process alive for now
}
