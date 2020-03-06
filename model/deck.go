package model

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
	"math/rand"
	"strconv"
)

const CardN = 52
const SuitN = 4

type Deck struct {
	gorm.Model
	DeckId     uuid.UUID
	IsShuffled bool
	Cards      pq.Int64Array `gorm:"type:integer[]"`
}

type ClosedDeckJson struct {
	DeckId     uuid.UUID `json:"deck_id"`
	IsShuffled bool      `json:"shuffled"`
	Remaining  int64     `json:"remaining"`
}

type OpenDeckJson struct {
	DeckId     uuid.UUID  `json:"deck_id"`
	IsShuffled bool       `json:"shuffled"`
	Remaining  int64      `json:"remaining"`
	Cards      []CardJson `json:"cards"`
}

type CardJson struct {
	Value string `json:"value"`
	Suit  string `json:"suit"`
	Code  string `json:"code"`
}

func (d *Deck) ToClosedDeckJson() ClosedDeckJson {
	return ClosedDeckJson{d.DeckId, d.IsShuffled, int64(len(d.Cards))}
}

func (d *Deck) ToOpenDeckJson() OpenDeckJson {
	return OpenDeckJson{d.DeckId, d.IsShuffled, int64(len(d.Cards)), IdsToCardJsons(d.Cards)}
}

func (d *Deck) Shuffle() {
	rand.Shuffle(len(d.Cards), func(i, j int) {
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	})
	d.IsShuffled = true
}

func IdsToCardJsons(ids []int64) (cardJsons []CardJson) {
	for _, id := range ids {
		cardJsons = append(cardJsons, IdToCardJson(id))
	}
	return
}

func IdToCardJson(id int64) (cardJson CardJson) {
	// Match value
	value := id % (CardN / SuitN)
	switch value {
	case 0:
		cardJson.Value = "ACE"
	case 10:
		cardJson.Value = "JACK"
	case 11:
		cardJson.Value = "QUEEN"
	case 12:
		cardJson.Value = "KING"
	default:
		cardJson.Value = strconv.Itoa(int(value) + 1)
	}
	cardJson.Code = cardJson.Value[:1]

	// Match suit
	switch id / (CardN / SuitN) {
	case 0:
		cardJson.Suit = "CLUBS"
	case 1:
		cardJson.Suit = "DIAMONDS"
	case 2:
		cardJson.Suit = "HEARTS"
	case 3:
		cardJson.Suit = "SPADES"
	}
	cardJson.Code += cardJson.Suit[:1]

	return
}

func CodeToId(code string) int64 {
	return 0
}
