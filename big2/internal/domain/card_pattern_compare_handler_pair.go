package domain

type pairComparer struct{}

func (p *pairComparer) Compare(cp1 CardPattern, cp2 CardPattern) CompareResult {
	card1, card2 := cp1.GetBigOne(), cp2.GetBigOne()
	return card1.Compare(card2)
}

func (p *pairComparer) Match(cp1 CardPattern, cp2 CardPattern) bool {
	if cp1.GetName() == "Pair" && cp1.GetName() == cp2.GetName() {
		return true
	}
	return false
}

func NewPairComparer() CardPatternCompareHandler {
	return &pairComparer{}
}
