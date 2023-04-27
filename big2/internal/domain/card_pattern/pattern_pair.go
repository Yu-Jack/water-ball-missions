package card_pattern

import "big2/internal/domain/card"

type Pair struct{ pattern }

func (p *Pair) Validate(cards []*card.Card) bool {
	if len(cards) != p.size {
		return false
	}

	return cards[0].CompareRank(cards[1]) == card.CompareResultEqual
}

func (p *Pair) GetBigOne() *card.Card {
	return p.cards[0]
}

func NewPair() Pattern {
	return &Pair{
		pattern{
			size: 2,
			name: "Pair",
		},
	}
}
