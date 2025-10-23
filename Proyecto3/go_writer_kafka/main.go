package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("Mensaje recibido en Writer Kafka")
	fmt.Fprintln(w, "Kafka Writer OK")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("☕ Writer Kafka corriendo en :8082")
	log.Fatal(http.ListenAndServe(":8082", nil))
}
