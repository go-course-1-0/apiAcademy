package models

import "time"

type Admin struct {
	ID        int       `json:"id" gorm:"primaryKey" fake:"-"`
	FullName  string    `json:"fullName" gorm:"not null" fake:"{name}"`
	Email     string    `json:"email" gorm:"not null;unique" fake:"{email}"`
	Password  string    `json:"-" gorm:"not null" fake:"{password}"`
	CreatedAt time.Time `json:"createdAt" gorm:"default:now()" fake:"-"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"default:now()" fake:"-"`
}
