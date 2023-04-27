package card_pattern

import "big2/internal/domain/card"

type comparerPair struct{}

func (p *comparerPair) Compare(cp1 Pattern, cp2 Pattern) card.CompareResult {
	card1, card2 := cp1.GetBigOne(), cp2.GetBigOne()
	return card1.Compare(card2)
}

func (p *comparerPair) Match(cp1 Pattern, cp2 Pattern) bool {
	if cp1.GetName() == "Pair" && cp1.GetName() == cp2.GetName() {
		return true
	}
	return false
}

func NewComparerPair() CompareHandler {
	return &comparerPair{}
}
