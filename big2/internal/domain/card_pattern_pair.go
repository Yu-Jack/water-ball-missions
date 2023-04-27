package domain

type PairPattern struct{ cardPattern }

func (p *PairPattern) Validate(cards []*Card) bool {
	if len(cards) != p.size {
		return false
	}

	return cards[0].CompareRank(cards[1]) == CompareResultEqual
}

func (p *PairPattern) GetBigOne() *Card {
	return p.cards[0]
}

func NewPairPattern() CardPattern {
	return &PairPattern{
		cardPattern{
			size: 2,
			name: "Pair",
		},
	}
}
