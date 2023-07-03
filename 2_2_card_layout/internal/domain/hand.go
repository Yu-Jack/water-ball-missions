package domain

import (
	"strings"
)

type Hand struct {
	cards []Card
}

func NewHand() *Hand {
	return &Hand{
		cards: []Card{},
	}
}

func (h *Hand) AddCard(c Card) {
	h.cards = append(h.cards, c)
}

func (h *Hand) PickCard(order int) Card {
	c := h.cards[order]
	h.cards = append(h.cards[:order], h.cards[order+1:]...)
	return c
}

func (h *Hand) FindCardCompareResult(c Card) []int {
	result := make([]int, 0, len(h.cards))

	for _, card := range h.cards {
		result = append(result, card.Compare(c))
	}

	return result
}

func (h *Hand) Size() int {
	return len(h.cards)
}

func (h *Hand) List() string {
	var s []string

	for _, c := range h.cards {
		s = append(s, c.String())
	}

	return strings.Join(s, ", ")
}

func (h *Hand) GetCard(order int) Card {
	return h.cards[order]
}
