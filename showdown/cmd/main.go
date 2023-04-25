package main

import (
	"math/rand"
	"time"

	"jackyu.showdown/internal/domain"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	deck := domain.NewDeck()

	showdown := domain.NewShowdown(&deck)
	showdown.AddPlayer(domain.NewHumanPlayer(&showdown))
	showdown.AddPlayer(domain.NewAIPlayer(&showdown))
	showdown.AddPlayer(domain.NewAIPlayer(&showdown))
	showdown.AddPlayer(domain.NewAIPlayer(&showdown))

	showdown.Start()
}
