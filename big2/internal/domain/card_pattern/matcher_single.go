package card_pattern

import "big2/internal/domain/card"

type MatcherSingle struct{}

func (s *MatcherSingle) Match(cards []*card.Card) Pattern {
	cp := NewSingle(cards)

	if !cp.Validate() {
		return nil
	}

	return cp
}

func NewMatcherSingle() Matcher {
	return &MatcherSingle{}
}
