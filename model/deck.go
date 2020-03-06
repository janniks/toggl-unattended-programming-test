package model

import (
	"github.com/google/uuid"
	"math/rand"
	"strconv"
)

const CardN = 52
const SuitN = 4

type Deck struct {
	DeckId     uuid.UUID
	IsShuffled bool
	Cards      []uint
}

type DeckJson struct {
	DeckId     uuid.UUID  `json:"deck_id"`
	IsShuffled bool       `json:"shuffled"`
	Cards      []CardJson `json:"cards"`
	Remaining  uint       `json:"remaining"`
}

type CardJson struct {
	Value string `json:"value"`
	Suit  string `json:"suit"`
	Code  string `json:"code"`
}

func (d *Deck) ToJson() DeckJson {
	return DeckJson{d.DeckId, d.IsShuffled, d.cardJsons(), uint(len(d.Cards))}
}

func (d *Deck) Shuffle() {
	rand.Shuffle(len(d.Cards), func(i, j int) {
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	})
	d.IsShuffled = true
}

func (d *Deck) cardJsons() (cardJsons []CardJson) {
	for _, code := range d.Cards {
		cardJsons = append(cardJsons, idToCardJson(code))
	}
	return
}

func idToCardJson(id uint) (cardJson CardJson) {
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
		cardJson.Value = strconv.Itoa(int(value))
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

func codeToId(code string) uint {
	return 0
}
