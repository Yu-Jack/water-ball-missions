package card_pattern

import (
	"fmt"
	"sort"

	"big2/internal/domain/card"
)

type FullHouse struct{ pattern }

func (fh *FullHouse) Validate() bool {
	cards := fh.cards

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

func (fh *FullHouse) String() string {
	return fmt.Sprintf("%s %s", "葫蘆", fh.pattern.String())
}

func NewFullHouse(cards []*card.Card) Pattern {
	fh := &FullHouse{
		pattern{
			cards: cards,
			size:  5,
			name:  CardPatternFullHouse,
		},
	}

	if !fh.Validate() {
		return nil
	}

	return fh
}
