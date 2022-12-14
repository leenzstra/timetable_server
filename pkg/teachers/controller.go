package teachers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/leenzstra/timetable_server/common/db"
)

type handler struct {
	DB *db.Database
}

func RegisterRoutes(app *fiber.App, db *db.Database) {
	h := &handler{
		DB: db,
	}

	routes := app.Group("/teachers")
	routes.Get("/:filter", h.GetTeachers)
	routes.Get("/group/:group_name", h.GetGroupTeachers)
	routes.Get("/", h.GetTeachers)
	routes.Get("/eval/:id", h.GetTeacherEval)
	routes.Post("/set_mark", h.SetMark)
}
