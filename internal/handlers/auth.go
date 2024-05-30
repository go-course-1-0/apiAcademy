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

func (h *Handlers) Login(c *gin.Context) {
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
		ID: admin.ID,
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

func (h *Handlers) Logout(c *gin.Context) {

}
