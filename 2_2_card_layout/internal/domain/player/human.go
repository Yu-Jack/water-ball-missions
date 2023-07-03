package player

import (
	"fmt"
	"strconv"

	"card.layout/internal/domain"
)

type human struct{}

func NewHuman() Strategy {
	return &human{}
}

func (h *human) NameHimself() string {
	var name string
	fmt.Print("What is your name?: ")
	_, _ = fmt.Scanln(&name)
	fmt.Print("Hello, " + name + "!\n")
	return name
}

func (h *human) PickCard(orders []int, hand *domain.Hand) domain.Card {
	fmt.Println("Your hand:", hand.List())

	var input string
	fmt.Print("Which card do you want to pick? (card number): \n")
	for _, o := range orders {
		fmt.Printf("%d: %s\n", o, hand.GetCard(o).String())
	}
	_, _ = fmt.Scanln(&input)
	order, _ := strconv.Atoi(input)

	return hand.PickCard(order)
}
