package main

import (
	"apiAcademy/internal/database"
	"apiAcademy/internal/database/models"
	"apiAcademy/internal/database/seeder"
	"fmt"
	"log/slog"
	"os"
)

func main() {
	fmt.Println("Academy API Seeder\n")
	fmt.Println("Connecting to DB...")

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	db, err := database.GormConnect()
	if err != nil {
		logger.Error("cannot connect to DB via gorm", "err", err.Error())
		return
	}

	db.Migrator().DropTable(&models.Lesson{})
	db.Migrator().DropTable(&models.Student{})
	db.Migrator().DropTable(&models.Group{})
	db.Migrator().DropTable(&models.Course{})
	db.Migrator().DropTable(&models.Teacher{})
	db.Migrator().DropTable(&models.Admin{})

	if err := database.GormAutoMigrate(db); err != nil {
		logger.Error("cannot automigrate models", "err", err.Error())
		return
	}

	// Admins
	if err := seeder.SeedAdmins(db, 10); err != nil {
		logger.Error("cannot seed admins", "err", err.Error())
	}

	// Teachers
	if err := seeder.SeedTeachers(db, 10); err != nil {
		logger.Error("cannot seed teachers", "err", err.Error())
	}

	// Courses
	if err := seeder.SeedCourses(db, 10); err != nil {
		logger.Error("cannot seed courses", "err", err.Error())
	}

	// Groups
	if err := seeder.SeedGroups(db, 10); err != nil {
		logger.Error("cannot seed groups", "err", err.Error())
	}

	// Students
	if err := seeder.SeedStudents(db, 100); err != nil {
		logger.Error("cannot seed students", "err", err.Error())
	}

	// Lessons
	if err := seeder.SeedLessons(db, 10); err != nil {
		logger.Error("cannot seed lessons", "err", err.Error())
	}
}
