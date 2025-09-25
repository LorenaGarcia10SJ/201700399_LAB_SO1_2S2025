package main

import "log"

func ProcessKernelData() {
	cont := ReadProc("/proc/continfo_so1_201700399")
	sys := ReadProc("/proc/sysinfo_so1_201700399")
	log.Println("Datos contenedores:", cont)
	log.Println("Datos sistema:", sys)
}
