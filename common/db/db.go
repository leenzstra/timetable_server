package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Init(login, pass string) *gorm.DB {
	url := fmt.Sprintf("host=localhost user=%s password=%s dbname=timetabledb port=5432 sslmode=disable", login, pass)
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{Logger: logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
		SlowThreshold:             time.Second, // Slow SQL threshold
		LogLevel:                  logger.Info, // Log level
		IgnoreRecordNotFoundError: false,       // Ignore ErrRecordNotFound error for logger
		Colorful:                  true,        // Disable color
	})})

	if err != nil {
		log.Fatalln(err)
	}

	// db.AutoMigrate(&models.Book{})

	return db
}
