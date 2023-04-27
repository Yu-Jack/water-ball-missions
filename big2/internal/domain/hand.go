package domain

import (
	"fmt"
	"sort"
	"strings"

	"big2/internal/domain/card"
)

type Hand struct {
	cards []*card.Card
}

func NewHand() *Hand {
	return &Hand{}
}

func (h *Hand) Size() int {
	return len(h.cards)
}

func (h *Hand) AddCard(card *card.Card) {
	h.cards = append(h.cards, card)
}

func (h *Hand) PickCard(orders []int) []*card.Card {
	cards := make([]*card.Card, 0, len(orders))
	for _, order := range orders {

		// 如果 order 不合理直接回傳空的回去
		if order >= h.Size() {
			return nil
		}

		cards = append(cards, h.cards[order])
	}
	return cards
}

func (h *Hand) Remove(orders []int) {
	right := len(h.cards) - 1

	for _, order := range orders {
		h.cards[order] = h.cards[right]
		right--
	}

	h.cards = h.cards[:right+1]
}

func (h *Hand) Sort() {
	sort.Slice(h.cards, func(i, j int) bool {
		return h.cards[i].Compare(h.cards[j]) == card.CompareResultSmaller
	})
}

func (h *Hand) List() {
	index := ""
	cards := ""

	for i, card := range h.cards {
		indexS := fmt.Sprintf("%d", i)
		spaceNumber := len(card.String()) - len(indexS) + 1
		index += fmt.Sprintf("%d", i)
		index += strings.Repeat(" ", spaceNumber)
		cards += fmt.Sprintf("%s ", card.String())
	}

	fmt.Printf("%s\n%s\n", index, cards)
}
