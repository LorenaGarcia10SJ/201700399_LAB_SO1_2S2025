package main

import (
	"io/ioutil"
	"log"
)

func ReadProc(filePath string) string {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Println("Error leyendo /proc:", err)
		return ""
	}
	return string(data)
}
