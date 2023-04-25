package player

import (
	"fmt"
	"math/rand"

	"card.layout/internal/domain"
)

type ai struct{}

func NewAI() Strategy {
	return &ai{}
}

func (a *ai) NameHimself() string {
	name := fmt.Sprintf("%d", rand.Intn(100000))
	fmt.Print("Hello, " + name + "!\n")
	return name
}

func (a *ai) PickCard(orders []int, hand *domain.Hand) domain.Card {
	return hand.PickCard(orders[0])
}
