package domain

type CardPatternCompareHandler interface {
	Compare(cp1 CardPattern, cp2 CardPattern) CompareResult
	Match(cp1 CardPattern, cp2 CardPattern) bool
}

type cardPatternCompareHandler struct {
	next     CardPatternCompareHandler
	comparer CardPatternCompareHandler
}

func (c *cardPatternCompareHandler) Compare(cp1 CardPattern, cp2 CardPattern) CompareResult {

	if c.comparer.Match(cp1, cp2) {
		return c.comparer.Compare(cp1, cp2)
	}

	if c.next != nil {
		return c.next.Compare(cp1, cp2)
	}

	return -1
}

func (c *cardPatternCompareHandler) Match(cp1 CardPattern, cp2 CardPattern) bool {
	// parent method, it will not be used in this handler
	return false
}

func NewCardPatternCompareHandler(comparer CardPatternCompareHandler, next CardPatternCompareHandler) CardPatternCompareHandler {
	return &cardPatternCompareHandler{
		next:     next,
		comparer: comparer,
	}
}
