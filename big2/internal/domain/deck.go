package domain

type Deck struct {
	cards []*Card
}

func NewDeck(cards []*Card) *Deck {
	return &Deck{
		cards: cards,
	}
}

// Deal pops last one
func (d *Deck) Deal() *Card {
	c := d.cards[len(d.cards)-1]
	d.cards = d.cards[:len(d.cards)-1]
	return c
}

func (d *Deck) Size() int {
	return len(d.cards)
}
