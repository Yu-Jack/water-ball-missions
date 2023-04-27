package card_pattern

import "big2/internal/domain/card"

type MatcherPair struct{}

func (p *MatcherPair) Match(cards []*card.Card) Pattern {
	return NewPair(cards)
}

func NewMatcherPair() Matcher {
	return &MatcherPair{}
}
