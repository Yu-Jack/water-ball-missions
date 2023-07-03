package domain

import "fmt"

const MaxTurn = 13

type Showdown struct {
	turn    int
	deck    *Deck
	players []Player
	table   map[PlayerName]Card // key is a HumanPlayer name
}

func NewShowdown(deck *Deck) Showdown {
	return Showdown{
		turn:    0,
		deck:    deck,
		players: []Player{},
		table:   make(map[PlayerName]Card, 4),
	}
}

func (s *Showdown) AddPlayer(p Player) {
	s.players = append(s.players, p)
}

func (s *Showdown) Start() {
	for _, p := range s.players {
		p.Naming()
	}

	s.deck.Shuffle()

	for s.deck.Size() > 0 {
		for _, p := range s.players {
			p.DrawCard(s.deck)
		}
	}

	// 這裡我更改遊戲規則，避免到時候手牌全部都是空的
	for s.turn < MaxTurn {
		// 第一個動作：決定要不要換牌
		for _, p := range s.players {
			p.TakeTurn1(p)
		}

		// 第二個動作：出牌
		for _, p := range s.players {
			p.TakeTurn2(p)
		}

		// find which player is bigger
		var biggerOne PlayerName
		for name, card := range s.table {
			fmt.Printf("Player: %s shows %s  \n", name, card.String())

			if biggerOne == "" {
				biggerOne = name
				continue
			}

			if card.Bigger(s.table[biggerOne]) {
				biggerOne = name
			}
		}

		// give him/her a nice point
		for _, p := range s.players {
			if p.GetName() == biggerOne {
				fmt.Printf("%s got a point!\n", biggerOne)
				p.GainPoint()
			}
		}

		s.turn++
	}

	// find who is the winner
	point, winner := 0, PlayerName("")
	for _, p := range s.players {
		fmt.Printf("Player gets %s point %d!\n", p.GetName(), p.GetPoint())
		if p.GetPoint() > point {
			point = p.GetPoint()
			winner = p.GetName()
		}
	}

	fmt.Printf("Winner is %s! Point is %d\n", winner, point)
}

func (s *Showdown) PutCardOnTable(c Card, p PlayerName) {
	s.table[p] = c
}
