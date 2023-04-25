package player

import (
	"card.layout/internal/domain"
)

type player struct {
	name string
	hand *domain.Hand
	game *domain.Game
}

func (p *player) GetName() string {
	return p.name
}

func (p *player) GetHand() *domain.Hand {
	return p.hand
}

func (p *player) DrawCard(deck *domain.Deck) {
	p.hand.AddCard(deck.DrawCard())
}

func (p *player) HandSize() int {
	return p.hand.Size()
}

type Strategy interface {
	NameHimself() string
	PickCard(orders []int, hand *domain.Hand) domain.Card
}
