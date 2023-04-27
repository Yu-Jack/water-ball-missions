package card_pattern

import (
	"sort"

	"big2/internal/domain/card"
)

type Straight struct{ pattern }

func (s *Straight) Validate(cards []*card.Card) bool {
	if len(cards) != s.size {
		return false
	}

	sort.Slice(s.cards, func(i, j int) bool {
		return s.cards[i].Compare(s.cards[j]) == card.CompareResultSmaller
	})

	// 每一個牌之間 rank 都差 1
	for i := 1; i < len(s.cards); i++ {
		if s.cards[i].CompareRank(s.cards[i-1]) != 1 {
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

func NewStraight() Pattern {
	return &Straight{
		pattern{
			size: 5,
			name: CardPatternStraight,
		},
	}
}
