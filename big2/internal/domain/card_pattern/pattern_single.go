package card_pattern

import "big2/internal/domain/card"

type Single struct{ pattern }

func (s *Single) Validate(cards []*card.Card) bool {
	return len(cards) == s.size
}

func (s *Single) GetBigOne() *card.Card {
	return s.cards[0]
}

func NewSingle() Pattern {
	return &Single{
		pattern{
			size: 1,
			name: CardPatternSingle,
		},
	}
}
