package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	// Cargar plantilla de HTML
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Servir archivos estáticos (CSS, imágenes)
	app.Static("/static", "./static")

	// Ruta principal (Home)
	app.Get("/", func(c *fiber.Ctx) error {
		// Pasar datos a la plantilla
		return c.Render("index", fiber.Map{
			"Title": "Mi Aplicación con Fiber",
		})
	})

	// Iniciar servidor en el puerto 3000
	app.Listen(":3000")
}
