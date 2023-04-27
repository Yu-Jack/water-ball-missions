package main

import "big2/internal/domain"

func main() {

	// 設定所有 Card Pattern 「符合」規則
	matchers := domain.NewCardPatternMatcher(domain.NewFullHouseMatcher(),
		domain.NewCardPatternMatcher(domain.NewSingleMatcher(),
			domain.NewCardPatternMatcher(domain.NewPairMatcher(), nil),
		),
	)

	// 設定所有 Card Pattern 「比對」規則
	comparer := domain.NewCardPatternCompareHandler(domain.NewSingleComparer(),
		domain.NewCardPatternCompareHandler(domain.NewFullHouseComparer(),
			domain.NewCardPatternCompareHandler(domain.NewPairComparer(), nil),
		),
	)

	big := domain.NewBig2(matchers, comparer)
	big.AddPlayer(domain.NewPlayer(big, domain.NewHumanStrategy()))
	big.AddPlayer(domain.NewPlayer(big, domain.NewHumanStrategy()))
	big.AddPlayer(domain.NewPlayer(big, domain.NewHumanStrategy()))
	big.AddPlayer(domain.NewPlayer(big, domain.NewHumanStrategy()))

	big.Start()
}
