package worker

import "log"

// StartConsumer arranca un consumidor sencillo (simulado)
func StartConsumer() {
    log.Println("coco worker: consumer started")
    // Aquí se consumirían eventos externos por tenant
    select {}
}
