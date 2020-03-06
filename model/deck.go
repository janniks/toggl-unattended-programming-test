package model

import (
	"errors"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
	"math/rand"
	"strconv"
	"strings"
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

func CodeToId(code string) (id int64, err error) {
	defer func() {
		if r := recover(); r != nil {
			id = -1
			err = errors.New("invalid card parameter")
		}
	}()

	valueCharacter := code[0]
	suitCharacter := code[1]

	if !strings.Contains("0123456789AJQK", string(valueCharacter)) || !strings.Contains("CDHS", string(suitCharacter)) {
		return -1, errors.New("invalid card code provided")
	}

	var value int64
	switch valueCharacter {
	case 'A':
		value = 0
	case 'J':
		value = 10
	case 'Q':
		value = 11
	case 'K':
		value = 12
	default:
		valueInt := int64(valueCharacter - '0')
		switch valueInt {
		case 1:
			value = 9
		default:
			value = valueInt - 1
		}
	}

	var suit int64
	switch suitCharacter {
	case 'C':
		suit = 0
	case 'D':
		suit = 1
	case 'H':
		suit = 2
	case 'S':
		suit = 3
	}

	return suit*(CardN/SuitN) + value, nil
}
