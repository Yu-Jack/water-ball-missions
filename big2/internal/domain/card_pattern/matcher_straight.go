package card_pattern

import "big2/internal/domain/card"

type MatcherStraight struct{}

func (fh *MatcherStraight) Match(cards []*card.Card) Pattern {
	return NewStraight(cards)
}

func NewMatcherStraight() Matcher {
	return &MatcherStraight{}
}
