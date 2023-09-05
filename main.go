package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/gofiber/template/html/v2"
)

func main() {
	engine := html.New("./views", ".html")
	engine.Debug(true)
	engine.Reload(true)
	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "layouts/main",
	})
	app.Use(logger.New())

	contacts := ContactsHandler{
		contacts: NewContactsStore(),
	}
	app.Get("/", contacts.Redirect)
	app.Get("/contacts", contacts.Index)
	app.Get("/contacts/new", contacts.New)
	app.Post("/contacts/new", contacts.Create)
	app.Get("/contacts/:id", contacts.Show)
	app.Get("/contacts/:id/edit", contacts.Edit)
	app.Post("/contacts/:id/edit", contacts.Update)
	app.Post("/contacts/:id/delete", contacts.Delete)

	app.Static("/static", "./static")

	app.Listen(":8000")
}
