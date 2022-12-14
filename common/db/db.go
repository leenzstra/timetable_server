package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/leenzstra/timetable_server/common/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database struct {
	Gorm *gorm.DB
}

func New(login, pass, host string, port int) *Database {
	url := fmt.Sprintf("host=%s user=%s password=%s dbname=timetabledb port=%d sslmode=disable", host, login, pass, port)
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

	return &Database{Gorm: db}
}

func (d *Database) GetTeachers(filter string) ([]*models.Teacher, error) {
	var response []*models.Teacher

	err := d.Gorm.Raw("SELECT * FROM teachers WHERE LOWER(fio) like LOWER(?);", filter+"%").Scan(&response).Error
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (d *Database) GetGroups() ([]*models.Group, error) {
	var response []*models.Group

	err := d.Gorm.Raw("SELECT DISTINCT ON (group_name) * FROM groups;").Scan(&response).Error
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (d *Database) GetTeacherEvaluations(teacherId int) ([]*models.TeacherEvaluation, error) {
	var response []*models.TeacherEvaluation

	err := d.Gorm.Joins("JOIN teachers ON teachers.id = teacher_evaluations.teacher_id").Where("teachers.id = ?", teacherId).Find(&response).Error
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (d *Database) GetGroupTimetableTypes(groupName string) ([]*models.Group, error) {
	var response []*models.Group

	err := d.Gorm.Where(&models.Group{GroupName: groupName}).Find(&response).Error
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (d *Database) GetGroupSession(groupName string) (*models.Session, error) {
	var response *models.Session

	// возможно GroupId
	err := d.Gorm.Model(&models.Group{GroupName: groupName}).Association("Sessions").Find(&response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (d *Database) GetGroupTimetable(groupName, tableType string) ([]*models.Timetable, error) {
	var response []*models.Timetable
	var group = models.Group{}

	// возможно GroupId
	err := d.Gorm.Where(&models.Group{GroupName: groupName, TableKind: tableType}).Find(&group).Error
	if err != nil {
		return nil, err
	}

	// log.Println(group)
	err = d.Gorm.Model(&models.Group{ID: group.ID}).Association("Timetables").Find(&response)
	if err != nil {
		return nil, err
	}
	// log.Println(response)

	return response, nil
}

func (d *Database) GetTeacherById(teacherId int) (*models.Teacher, error) {
	var response *models.Teacher

	err := d.Gorm.Where("id = ?", teacherId).Find(&response).Error
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (d *Database) GetGroupByName(groupName string) (*models.Group, error) {
	var response *models.Group

	err := d.Gorm.Raw("SELECT * FROM groups WHERE LOWER(group_name) = LOWER(?) LIMIT 1;", groupName).Scan(&response).Error
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (d *Database) CreateTeacherEvaluation(teacherId int, mark int, comment string) error {
	err := d.Gorm.Create(&models.TeacherEvaluation{TeacherId: teacherId, Mark: mark, Comment: sql.NullString{String: comment, Valid: true}}).Error
	return err
}

