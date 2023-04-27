package card_pattern

import "big2/internal/domain/card"

type comparerSingle struct{}

func (s *comparerSingle) Compare(cp1 Pattern, cp2 Pattern) card.CompareResult {
	card1, card2 := cp1.GetBigOne(), cp2.GetBigOne()
	return card1.Compare(card2)
}

func (s *comparerSingle) Match(cp1 Pattern, cp2 Pattern) bool {
	return cp1.GetName() == CardPatternSingle && cp1.GetName() == cp2.GetName()
}

func NewComparerSingle() CompareHandler {
	return &comparerSingle{}
}
