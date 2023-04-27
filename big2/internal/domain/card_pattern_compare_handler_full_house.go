package domain

type fullHouseComparer struct{}

func (fh *fullHouseComparer) Compare(cp1 CardPattern, cp2 CardPattern) CompareResult {
	card1, card2 := cp1.GetBigOne(), cp2.GetBigOne()
	return card1.Compare(card2)
}

func (fh *fullHouseComparer) Match(cp1 CardPattern, cp2 CardPattern) bool {
	if cp1.GetName() == "FullHouse" && cp1.GetName() == cp2.GetName() {
		return true
	}
	return false
}

func NewFullHouseComparer() CardPatternCompareHandler {
	return &fullHouseComparer{}
}
