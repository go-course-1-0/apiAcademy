package main

import (
	"apiAcademy/internal/database"
	"apiAcademy/internal/handlers"
	"apiAcademy/internal/helpers"
	"github.com/gin-gonic/gin"
	"log/slog"
	"os"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	db, err := database.GormConnect()
	if err != nil {
		logger.Error("cannot connect to DB via gorm", "err", err.Error())
		return
	}

	if err := database.GormAutoMigrate(db); err != nil {
		logger.Error("cannot automigrate models", "err", err.Error())
		return
	}

	h := handlers.NewHandlers(db, logger)

	helpers.SetValidatorEngineToUseJSONTags()

	router := gin.Default()

	// #authorization
	// #middleware

	router.Static("/storage", "./storage")

	admins := router.Group("/admins")
	{
		admins.GET("/", h.GetAllAdmins)      // #done #withoutPagination
		admins.POST("/", h.CreateAdmin)      // #done
		admins.GET("/:id", h.GetOneAdmin)    // #done
		admins.PUT("/:id", h.UpdateAdmin)    // #done
		admins.DELETE("/:id", h.DeleteAdmin) // #done
	}

	teachers := router.Group("/teachers")
	{
		teachers.GET("/", h.GetAllTeachers)      // #done #withoutPagination
		teachers.POST("/", h.CreateTeacher)      // #done
		teachers.GET("/:id", h.GetOneTeacher)    // #done
		teachers.PUT("/:id", h.UpdateTeacher)    // #done
		teachers.DELETE("/:id", h.DeleteTeacher) // #done
	}

	courses := router.Group("/courses")
	{
		courses.GET("/", h.GetAllCourses)      // #done #withoutPagination
		courses.POST("/", h.CreateCourse)      // #done
		courses.GET("/:id", h.GetOneCourse)    // #done
		courses.PUT("/:id", h.UpdateCourse)    // #done
		courses.DELETE("/:id", h.DeleteCourse) // #done
	}

	groups := router.Group("/groups")
	{
		groups.GET("/", h.GetAllGroups)      // #done #withoutPagination
		groups.POST("/", h.CreateGroup)      // #done
		groups.GET("/:id", h.GetOneGroup)    // #done
		groups.PUT("/:id", h.UpdateGroup)    // #done
		groups.DELETE("/:id", h.DeleteGroup) // #done
	}

	students := router.Group("/students")
	{
		students.GET("/", h.GetAllStudents)            // #done #withoutPagination
		students.POST("/", h.CreateStudent)            // #done #validateAge
		students.GET("/:id", h.GetOneStudent)          // #done
		students.PUT("/:id", h.UpdateStudent)          // #done #validateAge
		students.DELETE("/:id", h.DeleteStudent)       // #done
		students.POST("/:id/avatar", h.UploadAvatar)   // #done
		students.DELETE("/:id/avatar", h.RemoveAvatar) // #done
	}

	lessons := router.Group("/lessons")
	{
		lessons.GET("/", h.GetAllLessons)      // #done
		lessons.POST("/", h.CreateLesson)      // #done #timezone
		lessons.GET("/:id", h.GetOneLesson)    // #done
		lessons.PUT("/:id", h.UpdateLesson)    // #done #timezone
		lessons.DELETE("/:id", h.DeleteLesson) // #done
	}

	router.Run(":4000")
}
