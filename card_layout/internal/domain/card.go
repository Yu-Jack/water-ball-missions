package domain

type Card interface {
	Compare(c Card) int
	GenerateDeck() []Card
	String() string
}
