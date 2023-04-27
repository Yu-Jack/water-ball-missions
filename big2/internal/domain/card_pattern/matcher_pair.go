package card_pattern

import "big2/internal/domain/card"

type MatcherPair struct{}

func (p *MatcherPair) Match(cards []*card.Card) Pattern {
	cp := NewPair()

	if !cp.Validate(cards) {
		return nil
	}

	return cp
}

func NewMatcherPair() Matcher {
	return &MatcherPair{}
}
