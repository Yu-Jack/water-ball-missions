package card_pattern

import "big2/internal/domain/card"

type Handler interface {
	Validate(cards []*card.Card) (bool, Pattern)
}

type handler struct {
	next      Handler
	validator Handler
}

func (h *handler) Validate(cards []*card.Card) (bool, Pattern) {

	if valid, cp := h.validator.Validate(cards); valid {
		return valid, cp
	} else {
		if h.next != nil {
			return h.next.Validate(cards)
		}
	}

	return false, nil
}

func NewValidatorHandler(validator Handler, next Handler) Handler {
	return &handler{
		next:      next,
		validator: validator,
	}
}
