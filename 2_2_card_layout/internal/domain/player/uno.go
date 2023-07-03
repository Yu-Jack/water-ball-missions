package player

import (
	"fmt"

	"card.layout/internal/domain"
)

type Uno struct {
	player
	strategy Strategy
}

func (pu *Uno) TakeTurn(player domain.Player) {
	fmt.Printf("%s's turn, hand size is %d\n", player.GetName(), player.HandSize())
	lastCard := pu.game.Table[len(pu.game.Table)-1]
	fmt.Printf("Table Current Card: %s\n", lastCard.String())

	result := player.GetHand().FindCardCompareResult(lastCard)

	var orders []int

	for i, r := range result {
		if r == 0 /* equal */ {
			// 存放所有可選的卡片順序
			// e.g. 1,2,5,9
			orders = append(orders, i)
		}
	}

	// 當沒有可以出的牌的時候
	if len(orders) == 0 {

		// deck 沒牌了
		if pu.game.Deck.Size() == 0 {
			for _, c := range pu.game.Table {
				pu.game.Deck.AddCard(c)
			}
			pu.game.Table = []domain.Card{} // reset
		}

		player.DrawCard(pu.game.Deck)
		fmt.Printf("%s draw a card because of rules\n", player.GetName())
		return
	}

	pu.game.AddCardToTable(player.PickCard(orders))
}

func (pu *Uno) NameHimself() {
	pu.name = pu.strategy.NameHimself()
}

func (pu *Uno) PickCard(orders []int) domain.Card {
	return pu.strategy.PickCard(orders, pu.hand)
}

func NewPlayerUno(game *domain.Game, playerStrategy Strategy) domain.Player {
	return &Uno{
		player: player{
			game: game,
			hand: domain.NewHand(),
		},
		strategy: playerStrategy,
	}
}
