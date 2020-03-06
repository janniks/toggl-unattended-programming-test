package api

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"janniks.com/toggl/initial/model"
	"net/http"
)

func CreateDeck() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := c.MustGet("db").(*gorm.DB)

		deck := model.Deck{DeckId: uuid.New(), Cards: cardSequence(52)}
		db.NewRecord(deck)
		db.Create(&deck)

		c.JSON(http.StatusOK, gin.H{"deck": deck})
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

func cardSequence(n int) []int {
	a := make([]int, n)
	for i := range a {
		a[i] = i
	}
	return a
}
