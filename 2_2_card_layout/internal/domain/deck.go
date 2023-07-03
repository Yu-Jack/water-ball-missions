package domain

import "fmt"

type Deck struct {
	cards []Card
}

func NewDeck(cards []Card) *Deck {
	return &Deck{
		cards: cards,
	}
}

func (d *Deck) Shuffle() {
	fmt.Println("Shuffle")
}

func (d *Deck) DrawCard() Card {
	c := d.cards[0] // pop first one
	d.cards = d.cards[1:]
	return c
}

func (d *Deck) AddCard(c Card) {
	d.cards = append(d.cards, c)
}

func (d *Deck) Size() int {
	return len(d.cards)
}
