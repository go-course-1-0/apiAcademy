package models

import (
	"github.com/paraparadox/datetime"
	"time"
)

type Student struct {
	ID          int           `json:"id" gorm:"primaryKey"`
	GroupID     int           `json:"groupID" gorm:"not null;"`
	FullName    string        `json:"fullName" gorm:"not null"`
	Phone       string        `json:"phone" gorm:"not null;unique"`
	Avatar      string        `json:"avatar" gorm:"default:null"`
	DateOfBirth datetime.Date `json:"dateOfBirth" gorm:"not null"`
	CreatedAt   time.Time     `json:"createdAt" gorm:"default:now()"`
	UpdatedAt   time.Time     `json:"updatedAt" gorm:"default:now()"`

	// student belongs to group
	Group *Group `json:"group,omitempty"`
}

// Avatar feature for students:
// 1. Additional field in Student struct (representing the link to the avatar file) #we choose this one
// 2. Create new struct (for example StudentPhoto), that will hold the link to the file and implement one-to-one relation between Student and new struct
