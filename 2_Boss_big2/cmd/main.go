package main

import (
	"big2/internal/domain"
	cardPattern "big2/internal/domain/card_pattern"
	playStrategy "big2/internal/domain/play_strategy"
)

func main() {

	// 設定所有 Card Pattern 「比對」規則
	comparer :=
		cardPattern.NewValidatorHandler(cardPattern.NewValidatorSingle(),
			cardPattern.NewValidatorHandler(cardPattern.NewValidatorFullHouse(),
				cardPattern.NewValidatorHandler(cardPattern.NewValidatorPair(),
					cardPattern.NewValidatorHandler(cardPattern.NewValidatorStraight(), nil),
				),
			),
		)

	big := domain.NewBig2(comparer)
	big.AddPlayer(domain.NewPlayer(big, playStrategy.NewHumanStrategy()))
	big.AddPlayer(domain.NewPlayer(big, playStrategy.NewHumanStrategy()))
	big.AddPlayer(domain.NewPlayer(big, playStrategy.NewHumanStrategy()))
	big.AddPlayer(domain.NewPlayer(big, playStrategy.NewHumanStrategy()))

	big.Start()
}
