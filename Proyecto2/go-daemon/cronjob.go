package main

import (
	"log"
	"os/exec"
)

func StartCronJob() {
	log.Println("Ejecutando cronjob para contenedores")
	cmd := exec.Command("/bin/bash", "../bash/crear_contenedores.sh")
	err := cmd.Run()
	if err != nil {
		log.Println("Error al ejecutar cronjob:", err)
	}
}
