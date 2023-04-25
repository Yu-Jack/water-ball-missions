package domain

import "fmt"

type Hand struct {
	cards []Card
}

func (h *Hand) AddCard(c Card) {
	h.cards = append(h.cards, c)
}

func (h *Hand) PickCard(order int) Card {
	c := h.cards[order]
	h.cards = append(h.cards[:order], h.cards[order+1:]...)
	return c
}

func (h *Hand) Size() int {
	return len(h.cards)
}

func (h *Hand) List() {
	for i, c := range h.cards {
		fmt.Printf("%d: %s\n", i, c.String())
	}
}
