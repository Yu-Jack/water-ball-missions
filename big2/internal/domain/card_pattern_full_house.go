package domain

import "sort"

type FullHousePattern struct{ cardPattern }

func (fh *FullHousePattern) Validate(cards []*Card) bool {
	if len(cards) != fh.size {
		return false
	}

	sort.Slice(cards, func(i, j int) bool {
		return cards[i].Compare(cards[j]) == CompareResultSmaller
	})

	// 1,1,1,3,3
	if cards[0].CompareRank(cards[1]) == CompareResultEqual &&
		cards[1].CompareRank(cards[2]) == CompareResultEqual &&
		cards[3].CompareRank(cards[4]) == CompareResultEqual {
		return true
	}

	// 1,1,3,3,3
	if cards[0].CompareRank(cards[1]) == CompareResultEqual &&
		cards[2].CompareRank(cards[3]) == CompareResultEqual &&
		cards[3].CompareRank(cards[4]) == CompareResultEqual {
		return true
	}

	return false
}

func (fh *FullHousePattern) GetBigOne() *Card {
	biggerOne := fh.cards[0]

	for i := 1; i < len(fh.cards); i++ {
		if fh.cards[i].Compare(biggerOne) == 1 {
			biggerOne = fh.cards[i]
		}
	}

	return biggerOne
}

func NewFullHousePattern() CardPattern {
	return &FullHousePattern{
		cardPattern{
			size: 5,
			name: "FullHouse",
		},
	}
}
