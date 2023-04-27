package domain

type SingleMatcher struct{}

func (s *SingleMatcher) Match(cards []*Card) CardPattern {
	cp := NewSinglePattern()

	if !cp.Validate(cards) {
		return nil
	}

	return cp
}

func NewSingleMatcher() CardPatternMatcher {
	return &SingleMatcher{}
}
