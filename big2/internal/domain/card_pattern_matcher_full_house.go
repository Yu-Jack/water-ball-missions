package domain

type FullHouseMatcher struct{}

func (fh *FullHouseMatcher) Match(cards []*Card) CardPattern {
	cp := NewFullHousePattern()

	if !cp.Validate(cards) {
		return nil
	}

	return cp
}

func NewFullHouseMatcher() CardPatternMatcher {
	return &FullHouseMatcher{}
}
