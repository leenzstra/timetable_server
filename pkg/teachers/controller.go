package teachers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(app *fiber.App, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	routes := app.Group("/teachers")
	routes.Get("/:filter", h.GetTeachers)
	routes.Get("/eval/:id", h.GetTeacherEval)
	routes.Post("/set_mark", h.SetMark)
}
