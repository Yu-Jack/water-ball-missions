package game_type

import (
	"fmt"

	"card.layout/internal/domain"
	"card.layout/internal/domain/card"
	"card.layout/internal/domain/player"
)

type showdown struct{}

func NewShowdown() domain.GameType {
	return &showdown{}
}

// LimitCard returns the minimum number of cards for each player
func (u *showdown) LimitCard() int {
	return 13
}

func (u *showdown) NewDeck() *domain.Deck {
	return domain.NewDeck(
		card.NewShowdown().GenerateDeck(),
	)
}

func (u *showdown) Prepare(game *domain.Game) {
	// do nothing
}

func (u *showdown) TurnCheck(game *domain.Game) {
	biggerCard := game.Table[len(game.Table)-1]
	biggerIndex := len(game.Table) - 1

	for i := len(game.Table) - 2; i >= len(game.Table)-1-3; i-- {
		if biggerCard.Compare(game.Table[i]) == 0 /* smaller */ {
			biggerCard = game.Table[i]
			biggerIndex = i
		}
	}

	game.Players[biggerIndex%len(game.Players)].(*player.Showdown).GainPoint()
}

func (u *showdown) EndGame(game *domain.Game) bool {
	for _, p := range game.Players {
		if p.HandSize() != 0 {
			return false
		}
	}

	// 手都要沒牌才會結束
	return true
}

func (u *showdown) CheckWinner(game *domain.Game) {
	bigger, playerName := 0, ""
	_ = bigger

	for _, p := range game.Players {
		p := p.(*player.Showdown)

		fmt.Printf("%s: %d\n", p.GetName(), p.GetPoint())

		if p.GetPoint() > bigger {
			playerName = p.GetName()
		}
	}

	game.SetWinner(playerName)
}
