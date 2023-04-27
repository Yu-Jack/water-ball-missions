package card_pattern

import "big2/internal/domain/card"

type MatcherStraight struct{}

func (fh *MatcherStraight) Match(cards []*card.Card) Pattern {
	cp := NewStraight()

	if !cp.Validate(cards) {
		return nil
	}

	return cp
}

func NewMatcherStraight() Matcher {
	return &MatcherStraight{}
}
