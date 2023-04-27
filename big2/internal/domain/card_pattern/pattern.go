package card_pattern

import (
	"fmt"

	"big2/internal/domain/card"
)

type Pattern interface {
	Validate(cards []*card.Card) bool
	GetBigOne() *card.Card
	SetCards(cards []*card.Card)
	GetName() string
	String() string
}

type pattern struct {
	name  string
	cards []*card.Card
	size  int
}

func (cp *pattern) GetName() string {
	return cp.name
}

func (cp *pattern) SetCards(cards []*card.Card) {
	cp.cards = cards
}

func (cp *pattern) String() string {
	cardString := ""

	for _, c := range cp.cards {
		cardString += fmt.Sprintf("%s ", c.String())
	}

	return fmt.Sprintf("%s %s", cp.GetName(), cardString)
}
