package domain

type SinglePattern struct{ cardPattern }

func (s *SinglePattern) Validate(cards []*Card) bool {
	return len(cards) == s.size
}

func (s *SinglePattern) GetBigOne() *Card {
	return s.cards[0]
}

func NewSinglePattern() CardPattern {
	return &SinglePattern{
		cardPattern{
			size: 1,
			name: "Single",
		},
	}
}
