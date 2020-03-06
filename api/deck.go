package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateDeck() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"statue": "success"})
	}
}

func OpenDeck() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"statue": "success"})
	}
}

func Draw() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"statue": "success"})
	}
}
