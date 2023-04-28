package card_pattern

import "big2/internal/domain/card"

type MatcherFullHouse struct{}

func (fh *MatcherFullHouse) Match(cards []*card.Card) Pattern {
	return NewFullHouse(cards)
}

func NewMatcherFullHouse() Matcher {
	return &MatcherFullHouse{}
}