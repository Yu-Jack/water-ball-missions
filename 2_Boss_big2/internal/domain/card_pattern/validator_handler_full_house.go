package card_pattern

import "big2/internal/domain/card"

type validatorFullHouse struct{}

func (fh *validatorFullHouse) Validate(cards []*card.Card) (bool, Pattern) {
	cp := NewFullHouse(cards)
	return cp != nil, cp
}

func NewValidatorFullHouse() Handler {
	return &validatorFullHouse{}
}
