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
		deck.Shuffle()

		db.NewRecord(deck)
		db.Create(&deck)

		c.JSON(http.StatusOK, deck.ToJson())
	}
}

func OpenDeck() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "success"})
	}
}

func Draw() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "success"})
	}
}

func cardSequence(n int) (a []uint) {
	for i := range make([]uint, n) {
		a = append(a, uint(i))
	}
	return
}
