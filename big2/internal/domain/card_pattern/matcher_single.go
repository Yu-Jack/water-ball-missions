package card_pattern

import "big2/internal/domain/card"

type MatcherSingle struct{}

func (s *MatcherSingle) Match(cards []*card.Card) Pattern {
	cp := NewSingle()

	if !cp.Validate(cards) {
		return nil
	}

	return cp
}

func NewMatcherSingle() Matcher {
	return &MatcherSingle{}
}
