package card

import (
	"fmt"

	"card.layout/internal/domain"
)

type uno struct {
	color  int
	number int
}

const (
	Number0 = iota
	Number1 = iota
	Number2 = iota
	Number3 = iota
	Number4 = iota
	Number5 = iota
	Number6 = iota
	Number7 = iota
	Number8 = iota
	Number9 = iota
)

const (
	ColorBlue   = iota
	ColorRed    = iota
	ColorYellow = iota
	ColorGreen  = iota
)

var (
	Numbers = []int{
		Number0, Number1, Number2, Number3,
		Number4, Number5, Number6, Number7,
		Number8, Number9,
	}
	Colors = []int{
		ColorBlue, ColorRed, ColorYellow, ColorGreen,
	}
)

func NewUno() domain.Card {
	return &uno{}
}

// Compare returns 0 (equal), 1 (different)
func (u *uno) Compare(card domain.Card) int {
	c := card.(*uno)

	if c.number == u.number || c.color == u.color {
		return 0
	}

	return 1
}

func (u *uno) GenerateDeck() []domain.Card {
	cards := make([]domain.Card, 0, 52)

	for _, n := range Numbers {
		for _, c := range Colors {
			cards = append(cards, &uno{number: n, color: c})
		}
	}

	return cards
}

// for convenience of debugging
func (u *uno) String() string {
	color := ""

	switch u.color {
	case ColorBlue:
		color = "Blue"
	case ColorRed:
		color = "Red"
	case ColorYellow:
		color = "Yellow"
	case ColorGreen:
		color = "Green"
	}

	return fmt.Sprintf("%s%d", color, u.number)
}
