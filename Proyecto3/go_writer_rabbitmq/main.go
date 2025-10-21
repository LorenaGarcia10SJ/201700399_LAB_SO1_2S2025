package main

import (
    "fmt"
    "time"
)

func main() {
    fmt.Println("go_writer_rabbitmq started - placeholder")
    for {
        fmt.Println("RabbitMQ writer idle - waiting for messages...")
        time.Sleep(10 * time.Second)
    }
}
