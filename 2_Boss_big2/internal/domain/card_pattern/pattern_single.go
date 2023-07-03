package card_pattern

import (
	"fmt"

	"big2/internal/domain/card"
)

type Single struct{ pattern }

func (s *Single) Validate() bool {
	return len(s.cards) == s.size
}

func (s *Single) GetBigOne() *card.Card {
	return s.cards[0]
}

func (s *Single) String() string {
	return fmt.Sprintf("%s %s", "單張", s.pattern.String())
}

func NewSingle(cards []*card.Card) Pattern {
	s := &Single{
		pattern{
			cards: cards,
			size:  1,
			name:  CardPatternSingle,
		},
	}

	if !s.Validate() {
		return nil
	}

	return s
}
