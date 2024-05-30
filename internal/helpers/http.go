package helpers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func InternalServerError(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
		"message": "Внутренняя ошибка сервера",
	})
}

func NotFound(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
		"message": "Ресурс не найден",
	})
}

func BadRequest(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		"message": "Неправильный запрос",
	})
}

func Unauthorized(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		"message": "У вас нет доступа",
	})
}
