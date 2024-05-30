package middlewares

import (
	"apiAcademy/internal/auth"
	"apiAcademy/internal/database/models"
	"apiAcademy/internal/helpers"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
	"log/slog"
	"net/http"
	"slices"
	"strings"
)

func APIKeyAuth(allowList []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// API Key
		service := c.GetHeader("Service-Token")
		if !slices.Contains(allowList, service) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "У вас нет доступа",
			})
			return
		}
		c.Next()
	}
}

func JWTAuth(db *gorm.DB, logger *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		authToken := c.GetHeader("Authorization")
		authSlice := strings.Split(authToken, " ")
		if len(authSlice) != 2 {
			helpers.Unauthorized(c)
			return
		}
		jwtToken := authSlice[1]
		claims := &auth.Claims{}

		token, err := jwt.ParseWithClaims(jwtToken, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("top-secret-key"), nil
		})
		if err != nil || !token.Valid {
			logger.Error("JWT token provided in cookie is not valid or there is an error while parsing it", "err", err.Error())
			helpers.Unauthorized(c)
			return
		}

		var admin models.Admin
		if err := db.Where("id = ?", claims.ID).
			First(&admin).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				helpers.Unauthorized(c)
				return
			}
			logger.Error("cannot get admin", "err", err.Error())
			helpers.InternalServerError(c)
			return
		}

		c.Next()
	}
}
