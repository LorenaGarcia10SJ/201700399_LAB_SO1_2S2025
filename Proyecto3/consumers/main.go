package main

import (
    "fmt"
    "time"
)

func main() {
    fmt.Println("consumer started - simulating storing to Valkey")
    for {
        fmt.Println("consumer: polling message, storing to Valkey (simulated)")
        time.Sleep(8 * time.Second)
    }
}
