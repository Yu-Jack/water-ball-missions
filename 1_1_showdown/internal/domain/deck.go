package domain

import "fmt"

type Deck struct {
	cards []Card
}

var (
	Ranks = []int{
		RANK2, RANK3, RANK4, RANK5,
		RANK6, RANK7, RANK8, RANK9,
		RANK10, RANKJ, RANKQ, RANKK, RANKA,
	}
	Suits = []int{
		SuitClub, SuitDiamond, SuitHeart, SuitSpade,
	}
)

func NewDeck() Deck {
	cards := make([]Card, 0, 52)

	for _, r := range Ranks {
		for _, s := range Suits {
			cards = append(cards, Card{rank: r, suit: s})
		}
	}

	return Deck{
		cards: cards,
	}
}

func (d *Deck) DrawCard() Card {
	c := d.cards[0] // pop first one
	d.cards = d.cards[1:]
	return c
}

func (d *Deck) Shuffle() {
	// do nothing for first version
	fmt.Println("Shuffle")
}

func (d *Deck) Size() int {
	return len(d.cards)
}
