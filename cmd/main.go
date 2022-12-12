package main

import (
	"flag"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"

	"github.com/leenzstra/timetable_server/common/config"
	"github.com/leenzstra/timetable_server/common/db"
	_ "github.com/leenzstra/timetable_server/docs"
	"github.com/leenzstra/timetable_server/pkg/teachers"
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
	isProd :=flag.Bool("prod", false, "")
	flag.Parse()

	log.Printf("isProd %v", *isProd)

	c, err := config.LoadConfig(*isProd)

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	app := fiber.New()
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(cors.New(cors.ConfigDefault))
	// api := app.Group("/api")

	db := db.Init(c.PostgresLogin, c.PostgresPass, c.PostgresHost)

	app.Get("/swagger/*", swagger.HandlerDefault)
	timetable.RegisterRoutes(app, db)
	teachers.RegisterRoutes(app, db)

	app.Listen(":3000")
}
