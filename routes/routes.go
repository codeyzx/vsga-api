package route

import (
	"vsga-api/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func RouteInit(r *fiber.App) {
	r.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect("/docs")
	})

	r.Get("/docs/*", swagger.HandlerDefault)

	r.Get("/employe", handler.EmployeHandlerGetAll)
	r.Get("/employe/:id", handler.EmployeHandlerGetById)
	r.Post("/employe", handler.EmployeHandlerCreate)
	r.Put("/employe/:id", handler.EmployeHandlerUpdate)
	r.Delete("/employe/:id", handler.EmployeHandlerDelete)
}
