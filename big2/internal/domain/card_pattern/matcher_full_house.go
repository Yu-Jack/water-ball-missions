package card_pattern

import "big2/internal/domain/card"

type MatcherFullHouse struct{}

func (fh *MatcherFullHouse) Match(cards []*card.Card) Pattern {
	cp := NewFullHouse()

	if !cp.Validate(cards) {
		return nil
	}

	return cp
}

func NewMatcherFullHouse() Matcher {
	return &MatcherFullHouse{}
}
