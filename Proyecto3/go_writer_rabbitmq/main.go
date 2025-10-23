package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("📩 Mensaje recibido en Writer RabbitMQ (simulado)")
	fmt.Fprintln(w, "RabbitMQ Writer OK")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("🐇 Writer RabbitMQ corriendo en :8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
