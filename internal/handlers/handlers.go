package handlers

import (
	"gorm.io/gorm"
	"log/slog"
)

type Handlers struct {
	DB     *gorm.DB
	Logger *slog.Logger
}

func NewHandlers(db *gorm.DB, logger *slog.Logger) *Handlers {
	return &Handlers{
		DB:     db,
		Logger: logger,
	}
}
