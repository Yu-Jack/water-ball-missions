package domain

type singleComparer struct{}

func (s *singleComparer) Compare(cp1 CardPattern, cp2 CardPattern) CompareResult {
	card1, card2 := cp1.GetBigOne(), cp2.GetBigOne()
	return card1.Compare(card2)
}

func (s *singleComparer) Match(cp1 CardPattern, cp2 CardPattern) bool {
	if cp1.GetName() == "Single" && cp1.GetName() == cp2.GetName() {
		return true
	}
	return false
}

func NewSingleComparer() CardPatternCompareHandler {
	return &singleComparer{}
}
