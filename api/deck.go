package api

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"janniks.com/toggl/initial/model"
	"net/http"
	"strconv"
	"strings"
)

func CreateDeck() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Shuffle parameter
		shuffledQuery := c.DefaultQuery("shuffled", "false")
		shuffle := shuffledQuery == "true"

		// Cards parameter (i.e. whitelisted cards)
		cardsQuery := c.DefaultQuery("cards", "")
		var ids []int64
		if cardsQuery == "" {
			ids = cardSequence(model.CardN)
		} else {
			ids = parseCardCodes(cardsQuery)
		}

		db := c.MustGet("db").(*gorm.DB)

		deck := model.Deck{DeckId: uuid.New(), Cards: ids}
		if shuffle {
			deck.Shuffle()
		}

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
		count, err := strconv.Atoi(c.DefaultQuery("count", "1"))
		if err != nil {
			panic(err)
		}

		db := c.MustGet("db").(*gorm.DB)

		var deck model.Deck
		db.First(&deck, "deck_id = ?", c.Params.ByName("deck_id"))
		if len(deck.Cards) < count {
			count = len(deck.Cards)
		}

		cards := deck.Cards[:count]
		deck.Cards = deck.Cards[len(deck.Cards)-count:]
		db.Save(&deck)

		c.JSON(http.StatusOK, gin.H{"cards": model.IdsToCardJsons(cards)})
	}
}

func cardSequence(n int) (a []int64) {
	for i := range make([]int64, n) {
		a = append(a, int64(i))
	}
	return
}

func parseCardCodes(codeQuery string) (ids []int64) {
	codes := strings.Split(codeQuery, ":")
	for _, code := range codes {
		ids = append(ids, model.CodeToId(code))
	}
	return
}
