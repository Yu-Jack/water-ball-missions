package player

import "card.layout/internal/domain"

type Showdown struct {
	player
	point    int
	strategy Strategy
}

type HumanPlayerShowdown struct{ Showdown }
type AIPlayerShowdown struct{ Showdown }

func (ps *Showdown) GainPoint() {
	ps.point++
}

func (ps *Showdown) GetPoint() int {
	return ps.point
}

func (ps *Showdown) TakeTurn(player domain.Player) {
	orders := make([]int, player.HandSize())

	for i := 0; i < player.HandSize(); i++ {
		orders[i] = i
	}

	ps.game.AddCardToTable(player.PickCard(orders))
}

func (ps *Showdown) NameHimself() {
	ps.name = ps.strategy.NameHimself()
}

func (ps *Showdown) PickCard(orders []int) domain.Card {
	return ps.strategy.PickCard(orders, ps.hand)
}

func NewPlayerShowdown(game *domain.Game, playerStrategy Strategy) domain.Player {
	return &Showdown{
		player: player{
			game: game,
			hand: domain.NewHand(),
		},
		point:    0,
		strategy: playerStrategy,
	}
}
