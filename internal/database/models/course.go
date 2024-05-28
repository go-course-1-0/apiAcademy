package models

import "time"

type Course struct {
	ID        int       `json:"id" gorm:"primaryKey" fake:"-"`
	Title     string    `json:"title" gorm:"not null" fake:"{sentence}"`
	Duration  int       `json:"duration" gorm:"not null" fake:"{number:1,6}"`
	CreatedAt time.Time `json:"createdAt" gorm:"default:now()" fake:"-"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"default:now()" fake:"-"`

	// course has many groups
	Groups []Group `json:"groups,omitempty" fake:"-"`
}
