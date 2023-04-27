package card_pattern

import "big2/internal/domain/card"

type MatcherPair struct{}

func (p *MatcherPair) Match(cards []*card.Card) Pattern {
	cp := NewPair(cards)

	if !cp.Validate() {
		return nil
	}

	return cp
}

func NewMatcherPair() Matcher {
	return &MatcherPair{}
}
