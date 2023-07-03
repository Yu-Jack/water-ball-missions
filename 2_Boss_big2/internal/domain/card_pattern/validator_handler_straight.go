package card_pattern

import "big2/internal/domain/card"

type comparerStraight struct{}

func (s *comparerStraight) Validate(cards []*card.Card) (bool, Pattern) {
	cp := NewStraight(cards)
	return cp != nil, cp
}

func NewValidatorStraight() Handler {
	return &comparerStraight{}
}
