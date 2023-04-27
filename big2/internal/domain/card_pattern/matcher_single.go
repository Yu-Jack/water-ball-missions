package card_pattern

import "big2/internal/domain/card"

type MatcherSingle struct{}

func (s *MatcherSingle) Match(cards []*card.Card) Pattern {
	return NewSingle(cards)
}

func NewMatcherSingle() Matcher {
	return &MatcherSingle{}
}
