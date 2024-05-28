package models

import "time"

type Teacher struct {
	ID        int       `json:"id" gorm:"primaryKey" fake:"-"`
	FullName  string    `json:"fullName" gorm:"not null" fake:"{name}"`
	Subject   string    `json:"subject" gorm:"not null" fake:"{noun}"`
	CreatedAt time.Time `json:"createdAt" gorm:"default:now()" fake:"-"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"default:now()" fake:"-"`

	// teacher has many groups
	Groups []Group `json:"groups,omitempty" fake:"-"`
}
