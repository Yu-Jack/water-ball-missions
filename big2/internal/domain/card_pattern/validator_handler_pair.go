package card_pattern

import "big2/internal/domain/card"

type validatorPair struct{}

func (p *validatorPair) Validate(cards []*card.Card) (bool, Pattern) {
	cp := NewPair(cards)
	return cp != nil, cp
}

func NewValidatorPair() Handler {
	return &validatorPair{}
}
