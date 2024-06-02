package handlers

import (
	"apiAcademy/internal/database/models"
	"apiAcademy/internal/helpers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handlers) GetTeacherGroups(c *gin.Context) {
	teacher, exists := c.Get("authenticatedTeacher")
	if !exists {
		helpers.Unauthorized(c)
		return
	}

	t := teacher.(models.Teacher)
	var groups []models.Group

	if err := h.DB.Where("teacher_id = ?", t.ID).
		Preload("Students").
		Preload("Teacher").
		Find(&groups).Error; err != nil {
		helpers.InternalServerError(c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"groups": groups,
	})
}
