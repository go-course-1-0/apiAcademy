package handlers

import (
	"apiAcademy/internal/database/models"
	"apiAcademy/internal/helpers"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/paraparadox/datetime"
	"gorm.io/gorm"
	"log"
	"net/http"
	"time"
)

func (h *Handlers) GetAllLessons(c *gin.Context) {
	var lessons []models.Lesson
	if err := h.DB.Preload("Group.Teacher").
		Find(&lessons).Error; err != nil {
		log.Println("cannot get lessons:", err.Error())
		helpers.InternalServerError(c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"lessons": lessons,
	})
}

type lessonRequest struct {
	GroupID   int          `json:"groupID" binding:"required,gte=1"`
	DayOfWeek time.Weekday `json:"dayOfWeek" binding:"required,oneof=0 1 2 3 4 5 6"`
	Time      string       `json:"time" binding:"omitempty,max=64"`
}

func (h *Handlers) CreateLesson(c *gin.Context) {
	validationErrors := make(map[string]string)

	var requestBody lessonRequest
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		helpers.FillValidationErrorTag(err, validationErrors)
	}

	lessonTime, err := time.Parse(datetime.TimeLayout, requestBody.Time)
	if err != nil {
		if _, exists := validationErrors["time"]; !exists {
			validationErrors["time"] = helpers.ValidationMessageForTag("time-format", "")
		}
	}

	var group models.Group
	if err := h.DB.Where("id = ?", requestBody.GroupID).First(&group).Error; err != nil {
		if _, exists := validationErrors["groupID"]; !exists {
			validationErrors["groupID"] = helpers.ValidationMessageForTag("exists", "")
		}
	}

	if len(validationErrors) != 0 {
		c.JSON(http.StatusUnprocessableEntity, validationErrors)
		return
	}

	lesson := models.Lesson{
		GroupID:   group.ID,
		DayOfWeek: requestBody.DayOfWeek,
		Time:      datetime.Time(lessonTime),
	}

	if err := h.DB.Create(&lesson).Error; err != nil {
		log.Println("cannot create lesson:", err.Error())
		helpers.InternalServerError(c)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"lesson": lesson,
	})
}

func (h *Handlers) GetOneLesson(c *gin.Context) {
	var lesson models.Lesson
	if err := h.DB.Where("id = ?", c.Param("id")).
		Preload("Group.Teacher").
		First(&lesson).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helpers.NotFound(c)
			return
		}
		helpers.InternalServerError(c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"lesson": lesson,
	})
}

func (h *Handlers) UpdateLesson(c *gin.Context) {
	var lesson models.Lesson
	if err := h.DB.Where("id = ?", c.Param("id")).
		First(&lesson).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helpers.NotFound(c)
			return
		}
		helpers.InternalServerError(c)
		return
	}

	validationErrors := make(map[string]string)

	var requestBody lessonRequest
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		helpers.FillValidationErrorTag(err, validationErrors)
	}

	lessonTime, err := time.Parse(datetime.TimeLayout, requestBody.Time)
	if err != nil {
		if _, exists := validationErrors["time"]; !exists {
			validationErrors["time"] = helpers.ValidationMessageForTag("time-format", "")
		}
	}

	var group models.Group
	if err := h.DB.Where("id = ?", requestBody.GroupID).First(&group).Error; err != nil {
		if _, exists := validationErrors["groupID"]; !exists {
			validationErrors["groupID"] = helpers.ValidationMessageForTag("exists", "")
		}
	}

	if len(validationErrors) != 0 {
		c.JSON(http.StatusUnprocessableEntity, validationErrors)
		return
	}

	lesson.GroupID = group.ID
	lesson.DayOfWeek = requestBody.DayOfWeek
	lesson.Time = datetime.Time(lessonTime)

	if err := h.DB.Save(&lesson).Error; err != nil {
		log.Println("cannot update lesson:", err.Error())
		helpers.InternalServerError(c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"lesson": lesson,
	})
}

func (h *Handlers) DeleteLesson(c *gin.Context) {
	var lesson models.Lesson
	if err := h.DB.Where("id = ?", c.Param("id")).
		First(&lesson).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helpers.NotFound(c)
			return
		}
		helpers.InternalServerError(c)
		return
	}

	if err := h.DB.Delete(&lesson).Error; err != nil {
		helpers.InternalServerError(c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Успешно удалено",
	})
}
