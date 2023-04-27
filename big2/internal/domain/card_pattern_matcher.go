package domain

type CardPatternMatcher interface {
	Match(cards []*Card) CardPattern
}

type cardPatternMatcher struct {
	next    CardPatternMatcher
	matcher CardPatternMatcher
}

func NewCardPatternMatcher(matcher CardPatternMatcher, next CardPatternMatcher) CardPatternMatcher {
	return &cardPatternMatcher{
		next:    next,
		matcher: matcher,
	}
}

func (c *cardPatternMatcher) Match(cards []*Card) CardPattern {
	cp := c.matcher.Match(cards)

	if cp != nil {
		cp.SetCards(cards)
		return cp
	}

	if cp == nil && c.next != nil {
		return c.next.Match(cards)
	}

	return nil
}
