package card_pattern

import (
	"sort"

	"big2/internal/domain/card"
)

type FullHouse struct{ pattern }

func (fh *FullHouse) Validate(cards []*card.Card) bool {
	if len(cards) != fh.size {
		return false
	}

	sort.Slice(cards, func(i, j int) bool {
		return cards[i].Compare(cards[j]) == card.CompareResultSmaller
	})

	// 1,1,1,3,3
	if cards[0].CompareRank(cards[1]) == card.CompareResultEqual &&
		cards[1].CompareRank(cards[2]) == card.CompareResultEqual &&
		cards[3].CompareRank(cards[4]) == card.CompareResultEqual {
		return true
	}

	// 1,1,3,3,3
	if cards[0].CompareRank(cards[1]) == card.CompareResultEqual &&
		cards[2].CompareRank(cards[3]) == card.CompareResultEqual &&
		cards[3].CompareRank(cards[4]) == card.CompareResultEqual {
		return true
	}

	return false
}

func (fh *FullHouse) GetBigOne() *card.Card {
	biggerOne := fh.cards[0]

	for i := 1; i < len(fh.cards); i++ {
		if fh.cards[i].Compare(biggerOne) == 1 {
			biggerOne = fh.cards[i]
		}
	}

	return biggerOne
}

func NewFullHouse() Pattern {
	return &FullHouse{
		pattern{
			size: 5,
			name: "FullHouse",
		},
	}
}
