package main

import (
	"log"
	"time"
)

func main() {
	log.Println("Daemon iniciado")
	go StartCronJob()
	for {
		ProcessKernelData()
		time.Sleep(20 * time.Second)
	}
}
