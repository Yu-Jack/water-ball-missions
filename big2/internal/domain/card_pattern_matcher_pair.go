package domain

type PairMatcher struct{}

func (p *PairMatcher) Match(cards []*Card) CardPattern {
	cp := NewPairPattern()

	if !cp.Validate(cards) {
		return nil
	}

	return cp
}

func NewPairMatcher() CardPatternMatcher {
	return &PairMatcher{}
}
