package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_CodeToId(t *testing.T) {
	id, err := CodeToId("AC")
	assert.Equal(t, int64(0), id)
	assert.Equal(t, nil, err)

	id, _ = CodeToId("2C")
	assert.Equal(t, int64(1), id)

	id, _ = CodeToId("3C")
	assert.Equal(t, int64(2), id)

	id, _ = CodeToId("AD")
	assert.Equal(t, int64(13), id)

	id, _ = CodeToId("1D")
	assert.Equal(t, int64(22), id)

	id, _ = CodeToId("AH")
	assert.Equal(t, int64(26), id)

	id, _ = CodeToId("QH")
	assert.Equal(t, int64(37), id)

	id, _ = CodeToId("2S")
	assert.Equal(t, int64(40), id)

	id, _ = CodeToId("6S")
	assert.Equal(t, int64(44), id)

	id, _ = CodeToId("JS")
	assert.Equal(t, int64(49), id)

	id, _ = CodeToId("QS")
	assert.Equal(t, int64(50), id)

	id, _ = CodeToId("KS")
	assert.Equal(t, int64(51), id)

	_, err = CodeToId("XX")
	assert.NotEqual(t, nil, err)

	_, err = CodeToId("1")
	assert.NotEqual(t, nil, err)

	_, err = CodeToId("ac")
	assert.NotEqual(t, nil, err)
}

func Test_IdToCardJson(t *testing.T) {
	cardJson := IdToCardJson(0)
	assert.Equal(t, CardJson{
		Value: "ACE",
		Suit:  "CLUBS",
		Code:  "AC",
	}, cardJson)

	cardJson = IdToCardJson(22)
	assert.Equal(t, CardJson{
		Value: "10",
		Suit:  "DIAMONDS",
		Code:  "1D",
	}, cardJson)

	cardJson = IdToCardJson(49)
	assert.Equal(t, CardJson{
		Value: "JACK",
		Suit:  "SPADES",
		Code:  "JS",
	}, cardJson)
}
