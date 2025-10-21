package main

import (
    "fmt"
    "log"
    "net"
    "net/http"
    "io/ioutil"
)

func main() {
    // Simple HTTP endpoint that simula recibir desde Rust y forwardear a writers via gRPC (placeholder)
    http.HandleFunc("/forward", func(w http.ResponseWriter, r *http.Request) {
        body, _ := ioutil.ReadAll(r.Body)
        fmt.Printf("go_client received: %s\n", string(body))
        // Aquí, en producción, llamarías al gRPC writer (Kafka/Rabbit) o publicas en broker
        w.Write([]byte("{"status":"forwarded"}"))
    })

    ln, err := net.Listen("tcp", ":9000")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("go_client listening on :9000")
    http.Serve(ln, nil)
}
