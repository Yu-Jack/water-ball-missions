package card_pattern

import "big2/internal/domain/card"

type CompareHandler interface {
	Compare(cp1 Pattern, cp2 Pattern) card.CompareResult
	Match(cp1 Pattern, cp2 Pattern) bool
}

type compareHandler struct {
	next     CompareHandler
	comparer CompareHandler
}

func (c *compareHandler) Compare(cp1 Pattern, cp2 Pattern) card.CompareResult {

	if c.comparer.Match(cp1, cp2) {
		return c.comparer.Compare(cp1, cp2)
	}

	if c.next != nil {
		return c.next.Compare(cp1, cp2)
	}

	return -1
}

func (c *compareHandler) Match(cp1 Pattern, cp2 Pattern) bool {
	// parent method, it will not be used in this handler
	return false
}

func NewComparerHandler(comparer CompareHandler, next CompareHandler) CompareHandler {
	return &compareHandler{
		next:     next,
		comparer: comparer,
	}
}
