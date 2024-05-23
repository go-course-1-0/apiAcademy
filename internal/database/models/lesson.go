package models

import (
	"github.com/paraparadox/datetime"
	"time"
)

type Lesson struct {
	ID        int           `json:"id" gorm:"primaryKey"`
	GroupID   int           `json:"groupID" gorm:"not null;"`
	DayOfWeek time.Weekday  `json:"dayOfWeek" gorm:"not null"`
	Time      datetime.Time `json:"time" gorm:"default:null"`
	CreatedAt time.Time     `json:"createdAt" gorm:"default:now()"`
	UpdatedAt time.Time     `json:"updatedAt" gorm:"default:now()"`

	// lesson belongs to group
	Group *Group `json:"group,omitempty"`
}
