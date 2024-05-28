package models

import (
	"github.com/paraparadox/datetime"
	"time"
)

type Group struct {
	ID        int           `json:"id" gorm:"primaryKey" fake:"-"`
	CourseID  int           `json:"courseID" gorm:"not null;" fake:"{number:1,10}"`
	TeacherID int           `json:"teacherID" gorm:"not null;" fake:"{number:1,10}"`
	Title     string        `json:"title" gorm:"not null" fake:"{sentence}"`
	Start     datetime.Date `json:"start" gorm:"default:null" fake:"-"`
	Finish    datetime.Date `json:"finish" gorm:"default:null" fake:"-"`
	CreatedAt time.Time     `json:"createdAt" gorm:"default:now()" fake:"-"`
	UpdatedAt time.Time     `json:"updatedAt" gorm:"default:now()" fake:"-"`

	// group belongs to course
	Course *Course `json:"course,omitempty" fake:"-"`
	// group belongs to teacher
	Teacher *Teacher `json:"teacher,omitempty" fake:"-"`

	// group has many students
	Students []Student `json:"students,omitempty" fake:"-"`
	// group has many lessons
	Lessons []Lesson `json:"lessons,omitempty" fake:"-"`
}
