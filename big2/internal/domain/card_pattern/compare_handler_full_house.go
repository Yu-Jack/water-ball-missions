package card_pattern

import "big2/internal/domain/card"

type comparerFullHouse struct{}

func (fh *comparerFullHouse) Compare(cp1 Pattern, cp2 Pattern) card.CompareResult {
	card1, card2 := cp1.GetBigOne(), cp2.GetBigOne()
	return card1.Compare(card2)
}

func (fh *comparerFullHouse) Match(cp1 Pattern, cp2 Pattern) bool {
	return cp1.GetName() == CardPatternFullHouse && cp1.GetName() == cp2.GetName()
}

func NewComparerFullHouse() CompareHandler {
	return &comparerFullHouse{}
}
