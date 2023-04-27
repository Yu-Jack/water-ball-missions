package card_pattern

import "big2/internal/domain/card"

type Single struct{ pattern }

func (s *Single) Validate() bool {
	return len(s.cards) == s.size
}

func (s *Single) GetBigOne() *card.Card {
	return s.cards[0]
}

func NewSingle(cards []*card.Card) Pattern {
	return &Single{
		pattern{
			cards: cards,
			size:  1,
			name:  CardPatternSingle,
		},
	}
}
