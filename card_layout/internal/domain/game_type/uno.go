package game_type

import (
	"fmt"

	"card.layout/internal/domain"
	"card.layout/internal/domain/card"
)

type uno struct{}

func NewUno() domain.GameType {
	return &uno{}
}

// LimitCard returns the minimum number of cards for each player
func (u *uno) LimitCard() int {
	return 5
}

func (u *uno) NewDeck() *domain.Deck {
	return domain.NewDeck(
		card.NewUno().GenerateDeck(),
	)
}

func (u *uno) Prepare(game *domain.Game) {
	game.AddCardToTable(game.Deck.DrawCard())
}

func (u *uno) TurnCheck(game *domain.Game) {
	// do nothing for uno
}

func (u *uno) EndGame(game *domain.Game) bool {
	for _, p := range game.Players {
		// 一有人手沒牌就結束
		if p.HandSize() == 0 {
			return true
		}
	}

	return false
}

func (u *uno) CheckWinner(game *domain.Game) {
	for _, p := range game.Players {
		if p.HandSize() == 0 {

			fmt.Println("===Table===")
			for _, c := range game.Table {
				fmt.Printf("%s\n", c.String())
			}
			fmt.Println("===Table===")

			game.SetWinner(p.GetName())
			return
		}
	}
}
