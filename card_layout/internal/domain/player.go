package domain

type Player interface {
	NameHimself()
	PickCard(order []int) Card
	DrawCard(deck *Deck)
	HandSize() int
	GetName() string
	GetHand() *Hand
	TakeTurn(player Player)
}
