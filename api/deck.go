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
		shuffleQuery := c.DefaultQuery("shuffle", "false")
		shuffle := shuffleQuery == "true"

		// Cards parameter (i.e. whitelisted cards)
		cardsQuery := c.DefaultQuery("cards", "")
		var ids []int64
		if cardsQuery == "" {
			ids = cardSequence(model.CardN)
		} else {
			var err error
			ids, err = parseCardCodes(cardsQuery)
			if err != nil {
				c.JSON(http.StatusBadRequest, err.Error())
				return
			}
		}

		db := c.MustGet("db").(*gorm.DB)

		deck := model.Deck{DeckId: uuid.New(), Cards: ids}
		if shuffle {
			deck.Shuffle()
		}

		db.NewRecord(deck)
		db.Create(&deck)

		c.JSON(http.StatusCreated, deck.ToClosedDeckJson())
	}
}

func OpenDeck() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := c.MustGet("db").(*gorm.DB)

		var deck model.Deck
		if db.First(&deck, "deck_id = ?", c.Params.ByName("deck_id")).RecordNotFound() {
			c.JSON(http.StatusNotFound, "deck not found")
			return
		}

		c.JSON(http.StatusOK, deck.ToOpenDeckJson())
	}
}

func Draw() gin.HandlerFunc {
	return func(c *gin.Context) {
		count, err := strconv.Atoi(c.DefaultQuery("count", "1"))
		if err != nil || count < 1 {
			c.JSON(http.StatusBadRequest, "invalid count parameter")
			return
		}

		db := c.MustGet("db").(*gorm.DB)

		var deck model.Deck
		if db.First(&deck, "deck_id = ?", c.Params.ByName("deck_id")).RecordNotFound() {
			c.JSON(http.StatusNotFound, "deck not found")
			return
		}

		remaining := int(deck.Remaining())
		if remaining == 0 {
			c.JSON(http.StatusOK, "deck is empty")
			return
		}

		if remaining < count {
			count = remaining
		}

		cards := deck.Cards[:count]
		deck.Cards = deck.Cards[count:]
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

func parseCardCodes(codeQuery string) (ids []int64, err error) {
	codes := strings.Split(codeQuery, ",")
	for _, code := range codes {
		id, err := model.CodeToId(code)
		if err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}
	return
}
