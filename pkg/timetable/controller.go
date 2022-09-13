package timetable

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

	routes := app.Group("/timetable")
	routes.Get("/groups", h.GetGroups)
	routes.Get("/timetables/:group_name", h.GetGroupTimetable)
	routes.Get("/sessions/:group_name", h.GetGroupSession)
}
