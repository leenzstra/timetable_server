package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/leenzstra/timetable_server/common/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Init(login, pass, host string) *gorm.DB {
	url := fmt.Sprintf("host=%s user=%s password=%s dbname=timetabledb port=5432 sslmode=disable",host, login, pass)
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{Logger: logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
		SlowThreshold:             time.Second, // Slow SQL threshold
		LogLevel:                  logger.Info, // Log level
		IgnoreRecordNotFoundError: false,       // Ignore ErrRecordNotFound error for logger
		Colorful:                  true,        // Disable color
	})})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.TeacherEvaluation{})

	return db
}
