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
	"path/filepath"
	"time"
)

func (h *Handlers) GetAllStudents(c *gin.Context) {
	var students []models.Student
	if err := h.DB.Preload("Group").
		Find(&students).Error; err != nil {
		log.Println("cannot get students:", err.Error())
		helpers.InternalServerError(c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"students": students,
	})
}

type studentRequest struct {
	GroupID     int    `json:"groupID" binding:"required,gte=1"`
	FullName    string `json:"fullName" binding:"required,max=64"`
	Phone       string `json:"phone" binding:"required,len=9,numeric"`
	DateOfBirth string `json:"dateOfBirth" binding:"required"`
}

func (h *Handlers) CreateStudent(c *gin.Context) {
	validationErrors := make(map[string]string)

	var requestBody studentRequest
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		helpers.FillValidationErrorTag(err, validationErrors)
	}

	dateOfBirth, err := time.Parse(datetime.DateLayout, requestBody.DateOfBirth)
	if err != nil {
		if _, exists := validationErrors["dateOfBirth"]; !exists {
			validationErrors["dateOfBirth"] = helpers.ValidationMessageForTag("date-format", "")
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

	student := models.Student{
		GroupID:     group.ID,
		FullName:    requestBody.FullName,
		Phone:       requestBody.Phone,
		DateOfBirth: datetime.Date(dateOfBirth),
	}

	if err := h.DB.Create(&student).Error; err != nil {
		log.Println("cannot create student:", err.Error())
		helpers.InternalServerError(c)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"student": student,
	})
}

func (h *Handlers) GetOneStudent(c *gin.Context) {
	var student models.Student
	if err := h.DB.Where("id = ?", c.Param("id")).
		Preload("Group.Teacher").
		First(&student).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helpers.NotFound(c)
			return
		}
		helpers.InternalServerError(c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"student": student,
	})
}

func (h *Handlers) UpdateStudent(c *gin.Context) {
	var student models.Student
	if err := h.DB.Where("id = ?", c.Param("id")).
		First(&student).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helpers.NotFound(c)
			return
		}
		helpers.InternalServerError(c)
		return
	}

	validationErrors := make(map[string]string)

	var requestBody studentRequest
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		helpers.FillValidationErrorTag(err, validationErrors)
	}

	dateOfBirth, err := time.Parse(datetime.DateLayout, requestBody.DateOfBirth)
	if err != nil {
		if _, exists := validationErrors["dateOfBirth"]; !exists {
			validationErrors["dateOfBirth"] = helpers.ValidationMessageForTag("date-format", "")
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

	student.GroupID = group.ID
	student.FullName = requestBody.FullName
	student.Phone = requestBody.Phone
	student.DateOfBirth = datetime.Date(dateOfBirth)

	if err := h.DB.Save(&student).Error; err != nil {
		log.Println("cannot update student:", err.Error())
		helpers.InternalServerError(c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"student": student,
	})
}

func (h *Handlers) DeleteStudent(c *gin.Context) {
	var student models.Student
	if err := h.DB.Where("id = ?", c.Param("id")).
		First(&student).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helpers.NotFound(c)
			return
		}
		helpers.InternalServerError(c)
		return
	}

	if err := h.DB.Delete(&student).Error; err != nil {
		helpers.InternalServerError(c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Успешно удалено",
	})
}

func (h *Handlers) UploadAvatar(c *gin.Context) {
	validationErrors := make(map[string]string)

	// find the entity by ID specified in URI params
	var student models.Student
	if err := h.DB.Where("id = ?", c.Param("id")).
		First(&student).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helpers.NotFound(c)
			return
		}
		helpers.InternalServerError(c)
		return
	}

	// Validate and bind the form data
	form, err := c.MultipartForm()
	if err != nil {
		helpers.BadRequest(c)
		return
	}

	// Get the uploaded files
	files := form.File["file"]
	if len(files) == 0 {
		validationErrors["file"] = helpers.ValidationMessageForTag("required", "")
	} else if len(files) > 1 {
		validationErrors["file"] = helpers.ValidationMessageForTag("max-number-of-elements", 1)
	}

	if len(validationErrors) != 0 {
		c.JSON(http.StatusUnprocessableEntity, validationErrors)
		return
	}

	fileName := "storage/student-avatars/" + student.Phone + "Avatar" + filepath.Ext(files[0].Filename)

	if err := c.SaveUploadedFile(files[0], "./"+fileName); err != nil {
		log.Println("cannot save image into filesystem:", err)
		helpers.InternalServerError(c)
		return
	}

	student.Avatar = "http://localhost:4000/" + fileName

	// Save the image record to the database
	if err := h.DB.Save(&student).Error; err != nil {
		log.Println("cannot save student:", err)
		helpers.InternalServerError(c)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"student": student,
	})
}

func (h *Handlers) RemoveAvatar(c *gin.Context) {
	// find the entity by ID specified in URI params
	var student models.Student
	if err := h.DB.Where("id = ?", c.Param("id")).
		First(&student).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helpers.NotFound(c)
			return
		}
		helpers.InternalServerError(c)
		return
	}

	if err := helpers.DeleteImage(student.Avatar); err != nil {
		log.Println("cannot remove student's avatar:", err.Error())
		helpers.InternalServerError(c)
		return
	}

	student.Avatar = ""

	// Save the image record to the database
	if err := h.DB.Save(&student).Error; err != nil {
		log.Println("cannot save student:", err)
		helpers.InternalServerError(c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"student": student,
	})
}
