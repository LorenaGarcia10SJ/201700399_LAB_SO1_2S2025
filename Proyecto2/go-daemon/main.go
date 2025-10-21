package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

const (
	bajoConsumo  = "bajo_consumo"
	altoCPU      = "alto_cpu"
	altoRAM      = "alto_ram"
	procContInfo = "/proc/continfo_so1_201700399"
	procSysInfo  = "/proc/sysinfo_so1_201700399"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// Captura Ctrl+C para limpiar antes de salir
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\nDeteniendo daemon y limpiando contenedores...")
		limpiarContenedoresProyecto()
		os.Exit(0)
	}()

	// 1. Cargar módulo de kernel
	err := cargarModuloKernel()
	if err != nil {
		log.Println("Advertencia:", err)
	} else {
		fmt.Println("Módulo cargado exitosamente")
	}

	// 2. Cronjob interno: genera contenedores cada 90 segundos
	go cronjobGenerarContenedores()

	// 3. Loop principal cada 40 segundos
	for {
		time.Sleep(40 * time.Second)
		gestionarContenedores()
	}
}

// Ejecuta script para cargar módulo de kernel
func cargarModuloKernel() error {
	cmd := exec.Command("bash", "cargar_modulos.sh")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error cargando módulo: %v, %s", err, string(out))
	}
	return nil
}

// Cronjob interno en Go
func cronjobGenerarContenedores() {
	for {
		time.Sleep(90 * time.Second) // <-- evita saturar la máquina
		fmt.Println("Cronjob: Generando 10 contenedores aleatorios...")
		for i := 0; i < 10; i++ {
			tipo := rand.Intn(3)
			var imagen, nombre string
			switch tipo {
			case 0:
				imagen = bajoConsumo
				nombre = fmt.Sprintf("bajo_%d", rand.Intn(10000))
			case 1:
				imagen = altoCPU
				nombre = fmt.Sprintf("altoCPU_%d", rand.Intn(10000))
			case 2:
				imagen = altoRAM
				nombre = fmt.Sprintf("altoRAM_%d", rand.Intn(10000))
			}
			crearContenedor(nombre, imagen)
			time.Sleep(3 * time.Second) // <-- pausa entre contenedores
		}
	}
}

// Crear contenedor Docker
func crearContenedor(nombre, imagen string) {
	cmd := exec.Command("docker", "run", "-d", "--name", nombre, imagen)
	err := cmd.Run()
	if err != nil {
		log.Printf("Error creando contenedor %s: %v\n", nombre, err)
	} else {
		fmt.Printf("Contenedor %s (%s) creado\n", nombre, imagen)
	}
}

// Gestión de contenedores
func gestionarContenedores() {
	containers := listarContenedoresProyecto()
	countBajo, countAltoCPU, countAltoRAM := 0, 0, 0

	for _, c := range containers {
		if strings.HasPrefix(c, "bajo_") {
			countBajo++
		} else if strings.HasPrefix(c, "altoCPU_") {
			countAltoCPU++
		} else if strings.HasPrefix(c, "altoRAM_") {
			countAltoRAM++
		}
	}

	// Mantener 3 bajos
	for countBajo < 3 {
		nombre := fmt.Sprintf("bajo_%d", rand.Intn(10000))
		crearContenedor(nombre, bajoConsumo)
		countBajo++
	}

	// Mantener 2 altos
	for countAltoCPU+countAltoRAM < 2 {
		tipo := rand.Intn(2) // 0=CPU,1=RAM
		var nombre, imagen string
		if tipo == 0 {
			nombre = fmt.Sprintf("altoCPU_%d", rand.Intn(10000))
			imagen = altoCPU
			countAltoCPU++
		} else {
			nombre = fmt.Sprintf("altoRAM_%d", rand.Intn(10000))
			imagen = altoRAM
			countAltoRAM++
		}
		crearContenedor(nombre, imagen)
	}

	// Eliminar extras si hay demasiados
	for _, c := range containers {
		if strings.HasPrefix(c, "bajo_") && countBajo > 3 {
			eliminarContenedor(c)
			countBajo--
		} else if strings.HasPrefix(c, "altoCPU_") && countAltoCPU > 2 {
			eliminarContenedor(c)
			countAltoCPU--
		} else if strings.HasPrefix(c, "altoRAM_") && countAltoRAM > 2 {
			eliminarContenedor(c)
			countAltoRAM--
		}
	}
}

// Listar contenedores creados por este proyecto
func listarContenedoresProyecto() []string {
	cmd := exec.Command("docker", "ps", "-a", "--format", "{{.Names}}")
	out, err := cmd.Output()
	if err != nil {
		log.Println("Error listando contenedores:", err)
		return nil
	}
	var lista []string
	for _, c := range strings.Split(string(out), "\n") {
		c = strings.TrimSpace(c)
		if c != "" && (strings.HasPrefix(c, "bajo_") || strings.HasPrefix(c, "altoCPU_") || strings.HasPrefix(c, "altoRAM_")) {
			lista = append(lista, c)
		}
	}
	return lista
}

// Eliminar contenedor Docker
func eliminarContenedor(nombre string) {
	cmd := exec.Command("docker", "rm", "-f", nombre)
	err := cmd.Run()
	if err != nil {
		log.Printf("Error eliminando contenedor %s: %v\n", nombre, err)
	} else {
		fmt.Printf("Contenedor %s eliminado\n", nombre)
	}
}

// Limpiar todos los contenedores del proyecto
func limpiarContenedoresProyecto() {
	containers := listarContenedoresProyecto()
	for _, c := range containers {
		eliminarContenedor(c)
	}
}
