package handlers

import (
	"apiAcademy/internal/database/models"
	"apiAcademy/internal/helpers"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/paraparadox/datetime"
	"gorm.io/gorm"
	"net/http"
	"time"
)

func (h *Handlers) GetAllGroups(c *gin.Context) {
	var groups []models.Group
	if err := h.DB.Preload("Course").
		Preload("Teacher").
		Find(&groups).Error; err != nil {
		h.Logger.Error("cannot get groups", "err", err.Error())
		helpers.InternalServerError(c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"groups": groups,
	})
}

// UUID or GUID
type groupRequest struct {
	CourseID  int    `json:"courseID" binding:"required,gte=1"`
	TeacherID int    `json:"teacherID" binding:"required,gte=1"`
	Title     string `json:"title" binding:"required,max=64"`
	Start     string `json:"start" binding:"omitempty"`
	Finish    string `json:"finish" binding:"omitempty"`
}

func (h *Handlers) CreateGroup(c *gin.Context) {
	validationErrors := make(map[string]string)

	var requestBody groupRequest
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		helpers.FillValidationErrorTag(err, validationErrors)
	}

	start, err := time.Parse(datetime.DateLayout, requestBody.Start)
	if err != nil {
		if _, exists := validationErrors["start"]; !exists {
			validationErrors["start"] = helpers.ValidationMessageForTag("date-format", "")
		}
	}

	finish, err := time.Parse(datetime.DateLayout, requestBody.Finish)
	if err != nil {
		if _, exists := validationErrors["finish"]; !exists {
			validationErrors["finish"] = helpers.ValidationMessageForTag("date-format", "")
		}
	}

	var course models.Course
	if err := h.DB.Where("id = ?", requestBody.CourseID).First(&course).Error; err != nil {
		if _, exists := validationErrors["courseID"]; !exists {
			validationErrors["courseID"] = helpers.ValidationMessageForTag("exists", "")
		}
	}

	var teacher models.Teacher
	if err := h.DB.Where("id = ?", requestBody.TeacherID).First(&teacher).Error; err != nil {
		if _, exists := validationErrors["teacherID"]; !exists {
			validationErrors["teacherID"] = helpers.ValidationMessageForTag("exists", "")
		}
	}

	if len(validationErrors) != 0 {
		c.JSON(http.StatusUnprocessableEntity, validationErrors)
		return
	}

	group := models.Group{
		CourseID:  course.ID,
		TeacherID: teacher.ID,
		Title:     requestBody.Title,
		Start:     datetime.Date(start),
		Finish:    datetime.Date(finish),
	}

	if err := h.DB.Create(&group).Error; err != nil {
		h.Logger.Error("cannot create group", "err", err.Error())
		helpers.InternalServerError(c)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"group": group,
	})
}

func (h *Handlers) GetOneGroup(c *gin.Context) {
	var group models.Group
	if err := h.DB.Where("id = ?", c.Param("id")).
		Preload("Course").
		Preload("Teacher").
		Preload("Students").
		Preload("Lessons").
		First(&group).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helpers.NotFound(c)
			return
		}
		h.Logger.Error("cannot get group", "err", err.Error())
		helpers.InternalServerError(c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"group": group,
	})
}

func (h *Handlers) UpdateGroup(c *gin.Context) {
	var group models.Group
	if err := h.DB.Where("id = ?", c.Param("id")).
		First(&group).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helpers.NotFound(c)
			return
		}
		h.Logger.Error("cannot get group", "err", err.Error())
		helpers.InternalServerError(c)
		return
	}

	validationErrors := make(map[string]string)

	var requestBody groupRequest
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		helpers.FillValidationErrorTag(err, validationErrors)
	}

	start, err := time.Parse(datetime.DateLayout, requestBody.Start)
	if err != nil {
		if _, exists := validationErrors["start"]; !exists {
			validationErrors["start"] = helpers.ValidationMessageForTag("date-format", "")
		}
	}

	finish, err := time.Parse(datetime.DateLayout, requestBody.Finish)
	if err != nil {
		if _, exists := validationErrors["finish"]; !exists {
			validationErrors["finish"] = helpers.ValidationMessageForTag("date-format", "")
		}
	}

	var course models.Course
	if err := h.DB.Where("id = ?", requestBody.CourseID).First(&course).Error; err != nil {
		if _, exists := validationErrors["courseID"]; !exists {
			validationErrors["courseID"] = helpers.ValidationMessageForTag("exists", "")
		}
	}

	var teacher models.Teacher
	if err := h.DB.Where("id = ?", requestBody.TeacherID).First(&teacher).Error; err != nil {
		if _, exists := validationErrors["teacherID"]; !exists {
			validationErrors["teacherID"] = helpers.ValidationMessageForTag("exists", "")
		}
	}

	if len(validationErrors) != 0 {
		c.JSON(http.StatusUnprocessableEntity, validationErrors)
		return
	}

	group.CourseID = course.ID
	group.TeacherID = teacher.ID
	group.Title = requestBody.Title
	group.Start = datetime.Date(start)
	group.Finish = datetime.Date(finish)

	if err := h.DB.Save(&group).Error; err != nil {
		h.Logger.Error("cannot update group", "err", err.Error())
		helpers.InternalServerError(c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"group": group,
	})
}

func (h *Handlers) DeleteGroup(c *gin.Context) {
	var group models.Group
	if err := h.DB.Where("id = ?", c.Param("id")).
		First(&group).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helpers.NotFound(c)
			return
		}
		h.Logger.Error("cannot get group", "err", err.Error())
		helpers.InternalServerError(c)
		return
	}

	if err := h.DB.Delete(&group).Error; err != nil {
		h.Logger.Error("cannot delete group", "err", err.Error())
		helpers.InternalServerError(c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Успешно удалено",
	})
}
