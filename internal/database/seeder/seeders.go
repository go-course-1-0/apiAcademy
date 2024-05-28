package seeder

import (
	"apiAcademy/internal/database/models"
	"github.com/brianvoe/gofakeit/v7"
	"gorm.io/gorm"
)

func SeedAdmins(db *gorm.DB, number int) error {
	for i := 0; i < number; i++ {
		var admin models.Admin

		//admin.FullName = gofakeit.Name()
		//admin.Email = gofakeit.Email()
		//admin.Password = gofakeit.Password(true, true, true, true, false, 8)

		if err := gofakeit.Struct(&admin); err != nil {
			return err
		}

		if err := db.Create(&admin).Error; err != nil {
			return err
		}
	}

	return nil
}

func SeedTeachers(db *gorm.DB, number int) error {
	for i := 0; i < number; i++ {
		var model models.Teacher

		if err := gofakeit.Struct(&model); err != nil {
			return err
		}

		if err := db.Create(&model).Error; err != nil {
			return err
		}
	}

	return nil
}

func SeedCourses(db *gorm.DB, number int) error {
	for i := 0; i < number; i++ {
		var model models.Course

		if err := gofakeit.Struct(&model); err != nil {
			return err
		}

		if err := db.Create(&model).Error; err != nil {
			return err
		}
	}

	return nil
}

func SeedGroups(db *gorm.DB, number int) error {
	for i := 0; i < number; i++ {
		var model models.Group

		if err := gofakeit.Struct(&model); err != nil {
			return err
		}

		if err := db.Create(&model).Error; err != nil {
			return err
		}
	}

	return nil
}

func SeedStudents(db *gorm.DB, number int) error {
	for i := 0; i < number; i++ {
		var model models.Student

		if err := gofakeit.Struct(&model); err != nil {
			return err
		}

		if err := db.Create(&model).Error; err != nil {
			return err
		}
	}

	return nil
}

func SeedLessons(db *gorm.DB, number int) error {
	for i := 0; i < number; i++ {
		var model models.Lesson

		if err := gofakeit.Struct(&model); err != nil {
			return err
		}

		if err := db.Create(&model).Error; err != nil {
			return err
		}
	}

	return nil
}
