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
	DateOfBirth datetime.Date `json:"dateOfBirth" gorm:"not null"`
	CreatedAt   time.Time     `json:"createdAt" gorm:"default:now()"`
	UpdatedAt   time.Time     `json:"updatedAt" gorm:"default:now()"`

	// student belongs to group
	Group *Group `json:"group,omitempty"`
}
