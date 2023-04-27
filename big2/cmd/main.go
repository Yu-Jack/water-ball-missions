package main

import (
	"big2/internal/domain"
	cardPattern "big2/internal/domain/card_pattern"
	playStrategy "big2/internal/domain/play_strategy"
)

func main() {

	// 設定所有 Card Pattern 「符合」規則
	matchers :=
		cardPattern.NewCardPatternMatcher(cardPattern.NewMatcherFullHouse(),
			cardPattern.NewCardPatternMatcher(cardPattern.NewMatcherSingle(),
				cardPattern.NewCardPatternMatcher(cardPattern.NewMatcherPair(),
					cardPattern.NewCardPatternMatcher(cardPattern.NewMatcherStraight(), nil),
				),
			),
		)

	// 設定所有 Card Pattern 「比對」規則
	comparer :=
		cardPattern.NewComparerHandler(cardPattern.NewComparerSingle(),
			cardPattern.NewComparerHandler(cardPattern.NewComparerFullHouse(),
				cardPattern.NewComparerHandler(cardPattern.NewComparerPair(),
					cardPattern.NewComparerHandler(cardPattern.NewComparerStraight(), nil),
				),
			),
		)

	big := domain.NewBig2(matchers, comparer)
	big.AddPlayer(domain.NewPlayer(big, playStrategy.NewAIStrategy()))
	big.AddPlayer(domain.NewPlayer(big, playStrategy.NewAIStrategy()))
	big.AddPlayer(domain.NewPlayer(big, playStrategy.NewAIStrategy()))
	big.AddPlayer(domain.NewPlayer(big, playStrategy.NewAIStrategy()))

	big.Start()
}
