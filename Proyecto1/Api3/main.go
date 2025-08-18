package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func main() {
	carnet := "201700399"
	app := fiber.New()

	// Endpoint raíz
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"mensaje": fmt.Sprintf(
				"Hola, responde la API: API3 en la VM2, desarrollada por el estudiante Ingrid Lorena Garcia Yantuche con carnet: %s",
				carnet,
			),
		})
	})

	// Llamada a API1
	app.Get("/api3/"+carnet+"/llamar-api1", func(c *fiber.Ctx) error {
		resp, err := http.Get("http://vm3:8081/") // API1 en VM3:8081
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": fmt.Sprintf("No se pudo contactar a API1: %v", err),
			})
		}
		defer resp.Body.Close()

		body, _ := ioutil.ReadAll(resp.Body)
		return c.Send(body)
	})

	// Llamada a API2
	app.Get("/api3/"+carnet+"/llamar-api2", func(c *fiber.Ctx) error {
		resp, err := http.Get("http://vm3:8082/") // API2 en VM3:8082
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": fmt.Sprintf("No se pudo contactar a API2: %v", err),
			})
		}
		defer resp.Body.Close()

		body, _ := ioutil.ReadAll(resp.Body)
		return c.Send(body)
	})

	fmt.Println("API3 corriendo en puerto 8083...")
	app.Listen("0.0.0.0:8083")
}
