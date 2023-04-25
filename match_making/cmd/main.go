package main

import (
	"match.making/internal/domain"
	"match.making/internal/domain/strategy"
)

func main() {
	matchMaking := domain.NewMatchMaking(
		[]domain.Strategy{
			strategy.NewDistanceBased(),
			strategy.NewReverse(),
		},
	)

	matchMaking.Match()
}
