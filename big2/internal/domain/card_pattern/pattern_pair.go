package card_pattern

import "big2/internal/domain/card"

type Pair struct{ pattern }

func (p *Pair) Validate() bool {
	if len(p.cards) != p.size {
		return false
	}

	return p.cards[0].CompareRank(p.cards[1]) == card.CompareResultEqual
}

func (p *Pair) GetBigOne() *card.Card {
	if p.cards[0].Compare(p.cards[1]) == card.CompareResultBigger {
		return p.cards[0]
	}

	return p.cards[1]
}

func NewPair(cards []*card.Card) Pattern {
	return &Pair{
		pattern{
			cards: cards,
			size:  2,
			name:  CardPatternPair,
		},
	}
}
