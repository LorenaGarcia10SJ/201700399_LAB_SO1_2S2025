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

	// Endpoint raíz, devuelve JSON
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"mensaje": fmt.Sprintf(
				"Hola, responde la API: API2 en la VM3, desarrollada por el estudiante Ingrid Lorena Garcia Yantuche con carnet: %s",
				carnet,
			),
		})
	})

	// Llamada a API1
	app.Get("/api2/"+carnet+"/llamar-api1", func(c *fiber.Ctx) error {
		resp, err := http.Get("http://vm3:8081/") // Endpoint raíz de API1
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": fmt.Sprintf("No se pudo contactar a API1: %v", err),
			})
		}
		defer resp.Body.Close()

		body, _ := ioutil.ReadAll(resp.Body)

		// Devolver directamente el JSON de API1
		return c.Send(body)
	})

	// Llamada a API3
	app.Get("/api2/"+carnet+"/llamar-api3", func(c *fiber.Ctx) error {
		resp, err := http.Get("http://vm2:8083/") // Endpoint raíz de API3
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": fmt.Sprintf("No se pudo contactar a API3: %v", err),
			})
		}
		defer resp.Body.Close()

		body, _ := ioutil.ReadAll(resp.Body)

		// Devolver directamente el JSON de API3
		return c.Send(body)
	})

	fmt.Println("API2 corriendo en puerto 8082...")
	app.Listen("0.0.0.0:8082")
}
