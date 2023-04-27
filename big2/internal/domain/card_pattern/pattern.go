//go:generate go-enum -f=$GOFILE --nocase

package card_pattern

import (
	"fmt"

	"big2/internal/domain/card"
)

// CardPattern
/*
ENUM(
Straight
FullHouse
Single
Pair
)
*/
type CardPattern string

type Pattern interface {
	Validate() bool
	GetBigOne() *card.Card
	GetCards() []*card.Card
	GetName() CardPattern
	String() string
}

type pattern struct {
	name  CardPattern
	cards []*card.Card
	size  int
}

func (cp *pattern) GetName() CardPattern {
	return cp.name
}

func (cp *pattern) SetCards(cards []*card.Card) {
	cp.cards = cards
}

func (cp *pattern) GetCards() []*card.Card {
	return cp.cards
}

func (cp *pattern) String() string {
	cardString := ""

	for _, c := range cp.cards {
		cardString += fmt.Sprintf("%s ", c.String())
	}

	name := ""
	switch cp.GetName() {
	case CardPatternStraight:
		name = "順子"
	case CardPatternFullHouse:
		name = "葫蘆"
	case CardPatternSingle:
		name = "單張"
	case CardPatternPair:
		name = "對子"
	}

	return fmt.Sprintf("%s %s", name, cardString)
}
