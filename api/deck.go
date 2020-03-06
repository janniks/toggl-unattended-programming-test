package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateDeck() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"statue": "success"})
	}
}