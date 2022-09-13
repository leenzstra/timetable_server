package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"

	_ "github.com/leenzstra/timetable_server/docs"
	"github.com/leenzstra/timetable_server/common/config"
	"github.com/leenzstra/timetable_server/common/db"
	"github.com/leenzstra/timetable_server/pkg/timetable"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:3000
// @BasePath  /

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	app := fiber.New()
	app.Use(logger.New())
	app.Use(recover.New())
	// api := app.Group("/api")

	db := db.Init(c.PostgresLogin, c.PostgresPass)

	app.Get("/swagger/*", swagger.HandlerDefault)
	timetable.RegisterRoutes(app, db)

	app.Listen(":3000")
}
