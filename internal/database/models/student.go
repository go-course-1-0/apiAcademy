package models

import (
	"github.com/paraparadox/datetime"
	"time"
)

type Student struct {
	ID          int           `json:"id" gorm:"primaryKey" fake:"-"`
	GroupID     int           `json:"groupID" gorm:"not null;" fake:"{number:1,10}"`
	FullName    string        `json:"fullName" gorm:"not null" fake:"{name}"`
	Email       string        `json:"email" gorm:"not null;unique" fake:"{email}"`
	Password    string        `json:"-" gorm:"not null" fake:"{password}"`
	Phone       string        `json:"phone" gorm:"not null;unique" fake:"{phone}"`
	Avatar      string        `json:"avatar" gorm:"default:null" fake:"-"`
	DateOfBirth datetime.Date `json:"dateOfBirth" gorm:"not null" fake:"-"`
	CreatedAt   time.Time     `json:"createdAt" gorm:"default:now()" fake:"-"`
	UpdatedAt   time.Time     `json:"updatedAt" gorm:"default:now()" fake:"-"`

	// email and password needed
	// student belongs to group
	Group *Group `json:"group,omitempty" fake:"-"`
}

// Avatar feature for students:
// 1. Additional field in Student struct (representing the link to the avatar file) #we choose this one
// 2. Create new struct (for example StudentPhoto), that will hold the link to the file and implement one-to-one relation between Student and new struct
