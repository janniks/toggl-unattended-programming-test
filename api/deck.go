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

		c.JSON(http.StatusOK, deck.ToClosedDeckJson())
	}
}

func OpenDeck() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := c.MustGet("db").(*gorm.DB)

		var deck model.Deck
		db.First(&deck, "deck_id = ?", c.Params.ByName("deck_id"))

		c.JSON(http.StatusOK, deck.ToOpenDeckJson())
	}
}

func Draw() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := c.MustGet("db").(*gorm.DB)

		var deck model.Deck
		db.First(&deck, "deck_id = ?", c.Params.ByName("deck_id"))
		// todo
		db.Save(&deck)

		c.JSON(http.StatusOK, deck)
	}
}

func cardSequence(n int) (a []int64) {
	for i := range make([]int64, n) {
		a = append(a, int64(i))
	}
	return
}
