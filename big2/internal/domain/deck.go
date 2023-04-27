package domain

import "big2/internal/domain/card"

type Deck struct {
	cards []*card.Card
}

func NewDeck(cards []*card.Card) *Deck {
	return &Deck{
		cards: cards,
	}
}

// Deal pops last one
func (d *Deck) Deal() *card.Card {
	c := d.cards[len(d.cards)-1]
	d.cards = d.cards[:len(d.cards)-1]
	return c
}

func (d *Deck) Size() int {
	return len(d.cards)
}
