package main

import (
	"match.making/internal/domain"
	"match.making/internal/domain/strategy"
)

func main() {
	matchMaking := domain.NewMatchMaking(
		//strategy.NewReverse(strategy.NewDistanceBased()),
		strategy.NewDistanceBased(),
	)

	matchMaking.Match()
}
