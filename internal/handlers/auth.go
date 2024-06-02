package handlers

import (
	"apiAcademy/internal/auth"
	"apiAcademy/internal/database/models"
	"apiAcademy/internal/helpers"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
	"log"
	"net/http"
	"time"
)

type loginRequest struct {
	Email    string `json:"email" binding:"required,email,max=64"`
	Password string `json:"password" binding:"required,min=8,max=64"`
}

func (h *Handlers) AdminLogin(c *gin.Context) {
	validationErrors := make(map[string]string)

	var requestBody loginRequest
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		helpers.FillValidationErrorTag(err, validationErrors)
	}

	if len(validationErrors) != 0 {
		c.JSON(http.StatusUnprocessableEntity, validationErrors)
		return
	}

	var admin models.Admin
	if err := h.DB.Where("email = ?", requestBody.Email).First(&admin).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helpers.Unauthorized(c)
			return
		}
		h.Logger.Error("cannot get admin", "err", err.Error())
		helpers.InternalServerError(c)
		return
	}

	if admin.Password != requestBody.Password {
		helpers.Unauthorized(c)
		return
	}

	expirationTime := time.Now().Add(1 * time.Minute)
	claims := &auth.Claims{
		ID:    admin.ID,
		Model: "admin",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("top-secret-key"))
	if err != nil {
		log.Println("cannot sign a token string:", err)
		helpers.InternalServerError(c)
		return
	}

	// save token

	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}

func (h *Handlers) AdminLogout(c *gin.Context) {

}

func (h *Handlers) TeacherLogin(c *gin.Context) {
	validationErrors := make(map[string]string)

	var requestBody loginRequest
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		helpers.FillValidationErrorTag(err, validationErrors)
	}

	if len(validationErrors) != 0 {
		c.JSON(http.StatusUnprocessableEntity, validationErrors)
		return
	}

	var teacher models.Teacher
	if err := h.DB.Where("email = ?", requestBody.Email).First(&teacher).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helpers.Unauthorized(c)
			return
		}
		h.Logger.Error("cannot get teacher", "err", err.Error())
		helpers.InternalServerError(c)
		return
	}

	if teacher.Password != requestBody.Password {
		helpers.Unauthorized(c)
		return
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &auth.Claims{
		ID:    teacher.ID,
		Model: "teacher",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("top-secret-key"))
	if err != nil {
		log.Println("cannot sign a token string:", err)
		helpers.InternalServerError(c)
		return
	}

	// save token

	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}

func (h *Handlers) TeacherLogout(c *gin.Context) {

}

func (h *Handlers) StudentLogin(c *gin.Context) {
	validationErrors := make(map[string]string)

	var requestBody loginRequest
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		helpers.FillValidationErrorTag(err, validationErrors)
	}

	if len(validationErrors) != 0 {
		c.JSON(http.StatusUnprocessableEntity, validationErrors)
		return
	}

	var student models.Student
	if err := h.DB.Where("email = ?", requestBody.Email).First(&student).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helpers.Unauthorized(c)
			return
		}
		h.Logger.Error("cannot get student", "err", err.Error())
		helpers.InternalServerError(c)
		return
	}

	if student.Password != requestBody.Password {
		helpers.Unauthorized(c)
		return
	}

	expirationTime := time.Now().Add(1 * time.Minute)
	claims := &auth.Claims{
		ID:    student.ID,
		Model: "student",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("top-secret-key"))
	if err != nil {
		log.Println("cannot sign a token string:", err)
		helpers.InternalServerError(c)
		return
	}

	// save token

	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}

func (h *Handlers) StudentLogout(c *gin.Context) {

}
