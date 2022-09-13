package models

import (
	"gorm.io/datatypes"
)

type Tabler interface {
	TableName() string
}

type Group struct {
	ID        uint
	Faculty   string
	Direction string
	GroupName string
	GroupUrl  string
	TableKind string
	Timetables []Timetable
	Sessions []Session
}

func (g Group) TableName() string {
	return "groups"
}

type Timetable struct {
	ID        uint
	GroupID   uint
	Day       string
	WeekNum   uint
	TableJson datatypes.JSON
}

func (t Timetable) TableName() string {
	return "timetables"
}

type Session struct {
	ID       uint
	GroupID  uint
	Addition string
	Exams    datatypes.JSON
}

func (s Session) TableName() string {
	return "session"
}
