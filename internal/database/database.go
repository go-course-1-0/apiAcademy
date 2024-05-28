package database

import (
	"apiAcademy/internal/database/models"
	"database/sql"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// use .env
const (
	DBHost     = "localhost" // 127.0.0.1
	DBPort     = 5432
	DBUser     = "postgres"
	DBPassword = "postgres"
	DBName     = "academy_db"
)

func UsualConnect(host, port, user, password, dbname string) (*sql.DB, error) {
	// DSN - data source name
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, err
}

func GormConnect() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Dushanbe", DBHost, DBPort, DBUser, DBPassword, DBName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func GormAutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.Admin{},
		&models.Teacher{},
		&models.Course{},
		&models.Group{},
		&models.Student{},
		&models.Lesson{},
	)
}
