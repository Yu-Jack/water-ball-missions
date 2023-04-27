package card_pattern

import "big2/internal/domain/card"

type comparerStraight struct{}

func (c *comparerStraight) Compare(cp1 Pattern, cp2 Pattern) card.CompareResult {
	card1, card2 := cp1.GetBigOne(), cp2.GetBigOne()
	return card1.Compare(card2)
}

func (c *comparerStraight) Match(cp1 Pattern, cp2 Pattern) bool {
	if cp1.GetName() == "Straight" && cp1.GetName() == cp2.GetName() {
		return true
	}
	return false
}

func NewComparerStraight() CompareHandler {
	return &comparerStraight{}
}
