package domain

import "fmt"

type GameType interface {
	LimitCard() int
	NewDeck() *Deck
	Prepare(game *Game)
	CheckWinner(game *Game)
	EndGame(game *Game) bool
	TurnCheck(game *Game)
}

type Game struct {
	Players []Player
	Deck    *Deck
	Table   []Card

	winner   string
	gameType GameType
}

func NewGame(gameType GameType) *Game {
	return &Game{
		Players:  []Player{},
		Deck:     nil,
		gameType: gameType,
	}
}

func (g *Game) AddCardToTable(c Card) {
	g.Table = append(g.Table, c)
}

func (g *Game) AddPlayer(p Player) {
	g.Players = append(g.Players, p)
}

func (g *Game) Naming() {
	for _, player := range g.Players {
		player.NameHimself()
	}
}

func (g *Game) Shuffle() {
	g.Deck = g.gameType.NewDeck()
	g.Deck.Shuffle()
}

func (g *Game) SetWinner(player string) {
	g.winner = player
}

func (g *Game) Start() {
	// 前置準備
	for _, player := range g.Players {
		player.NameHimself()
	}

	g.Deck = g.gameType.NewDeck()
	g.Deck.Shuffle()

	totalCard := g.gameType.LimitCard() * len(g.Players)

	for totalCard > 0 {
		for _, p := range g.Players {
			p.DrawCard(g.Deck)
			totalCard--
		}
	}

	// 前置準備好後，第一個是遊戲正式開始前的遊戲設置
	g.gameType.Prepare(g)

	// 正式開始
	endGame := false

	for !endGame {
		for _, p := range g.Players {
			p.TakeTurn(p)

			if endGame = g.gameType.EndGame(g); endGame {
				break
			}
		}

		g.gameType.TurnCheck(g)
	}

	g.gameType.CheckWinner(g)

	fmt.Printf("Winner is %s\n", g.winner)
}
