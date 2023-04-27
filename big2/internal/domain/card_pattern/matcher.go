package card_pattern

import "big2/internal/domain/card"

type Matcher interface {
	Match(cards []*card.Card) Pattern
}

type cardPatternMatcher struct {
	next    Matcher
	matcher Matcher
}

func NewCardPatternMatcher(matcher Matcher, next Matcher) Matcher {
	return &cardPatternMatcher{
		next:    next,
		matcher: matcher,
	}
}

func (c *cardPatternMatcher) Match(cards []*card.Card) Pattern {
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
