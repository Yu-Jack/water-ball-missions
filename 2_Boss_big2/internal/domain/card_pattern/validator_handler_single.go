package card_pattern

import "big2/internal/domain/card"

type validatorSingle struct{}

func (s *validatorSingle) Validate(cards []*card.Card) (bool, Pattern) {
	cp := NewSingle(cards)
	return cp != nil, cp
}

func NewValidatorSingle() Handler {
	return &validatorSingle{}
}
