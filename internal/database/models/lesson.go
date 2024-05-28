package models

import (
	"github.com/paraparadox/datetime"
	"time"
)

type Lesson struct {
	ID        int           `json:"id" gorm:"primaryKey" fake:"-"`
	GroupID   int           `json:"groupID" gorm:"not null;" fake:"{number:1,10}"`
	DayOfWeek time.Weekday  `json:"dayOfWeek" gorm:"not null" fake:"{number:0,6}"`
	Time      datetime.Time `json:"time" gorm:"default:null" fake:"-"`
	CreatedAt time.Time     `json:"createdAt" gorm:"default:now()" fake:"-"`
	UpdatedAt time.Time     `json:"updatedAt" gorm:"default:now()" fake:"-"`

	// lesson belongs to group
	Group *Group `json:"group,omitempty" fake:"-"`
}
