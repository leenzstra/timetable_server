package timetable

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

	routes := app.Group("/timetable")
	routes.Get("/groups", h.GetGroups)
	routes.Get("/timetables/:group_name/:timetable_type", h.GetGroupTimetable)
	routes.Get("/sessions/:group_name", h.GetGroupSession)
	routes.Get("/types/:group_name", h.GetTimetableTypes)
}
