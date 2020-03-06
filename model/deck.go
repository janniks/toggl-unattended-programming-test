package model

import "github.com/google/uuid"

type Deck struct {
	DeckId   uuid.UUID
	Shuffled bool
	Cards    []int
}
