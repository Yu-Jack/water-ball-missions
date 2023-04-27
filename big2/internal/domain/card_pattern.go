package domain

import "fmt"

type CardPattern interface {
	Validate(cards []*Card) bool
	GetBigOne() *Card
	SetCards(cards []*Card)
	GetName() string
	String() string
}

type cardPattern struct {
	name  string
	cards []*Card
	size  int
}

func (cp *cardPattern) GetName() string {
	return cp.name
}

func (cp *cardPattern) SetCards(cards []*Card) {
	cp.cards = cards
}

func (cp *cardPattern) String() string {
	cardString := ""

	for _, card := range cp.cards {
		cardString += fmt.Sprintf("%s ", card.String())
	}

	return fmt.Sprintf("%s %s", cp.GetName(), cardString)
}
