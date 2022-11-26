package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// @BasePath /api/v1

// Helloworld
// @Summary helloworld example
// @Schemes
// @Description do helloworld
// @Tags helloworld
// @Success 200 {string} helloworld
// @Router /helloworld [get]
func Helloworld(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "helloworld",
	})
}
