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
				"Hola, responde la API: API1 en la VM3, desarrollada por el estudiante Ingrid Garcia con carnet: %s",
				carnet,
			),
		})
	})

	// Llamada a API2
	app.Get("/api1/"+carnet+"/llamar-api2", func(c *fiber.Ctx) error {
		resp, err := http.Get("http://vm3:8082/") // Endpoint raíz de API2
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": fmt.Sprintf("No se pudo contactar a API2: %v", err),
			})
		}
		defer resp.Body.Close()

		body, _ := ioutil.ReadAll(resp.Body)

		// Devolver directamente el JSON de API2
		return c.Send(body)
	})

	// Llamada a API3
	app.Get("/api1/"+carnet+"/llamar-api3", func(c *fiber.Ctx) error {
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

	fmt.Println("API1 corriendo en puerto 8081...")
	app.Listen("0.0.0.0:8081")
}
