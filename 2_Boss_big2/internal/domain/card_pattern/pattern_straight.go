package card_pattern

import (
	"fmt"
	"sort"

	"big2/internal/domain/card"
)

type Straight struct{ pattern }

func (s *Straight) Validate() bool {
	if len(s.cards) != s.size {
		return false
	}

	sort.Slice(s.cards, func(i, j int) bool {
		return s.cards[i].Compare(s.cards[j]) == card.CompareResultSmaller
	})

	// 每一個牌之間 rank 都差 1
	for i := 1; i < len(s.cards); i++ {
		if s.cards[i].DiffRank(s.cards[i-1]) != 1 {
			return false
		}
	}

	return true
}

func (s *Straight) GetBigOne() *card.Card {
	sort.Slice(s.cards, func(i, j int) bool {
		return s.cards[i].Compare(s.cards[j]) == card.CompareResultBigger
	})
	return s.cards[0]
}

func (s *Straight) String() string {
	return fmt.Sprintf("%s %s", "順子", s.pattern.String())
}

func NewStraight(cards []*card.Card) Pattern {
	s := &Straight{
		pattern{
			cards: cards,
			size:  5,
			name:  CardPatternStraight,
		},
	}

	if !s.Validate() {
		return nil
	}

	return s
}
