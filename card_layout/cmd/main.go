package main

import (
	"fmt"
	"math/rand"
	"time"

	"card.layout/internal/domain"
	gameType "card.layout/internal/domain/game_type"
	player "card.layout/internal/domain/player"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// just for demo quickly
	var input string
	fmt.Print("Which game do you play?: ")
	_, _ = fmt.Scanln(&input)

	if input == "showdown" {
		game := domain.NewGame(gameType.NewShowdown())

		game.AddPlayer(player.NewPlayerShowdown(game, player.NewHuman()))
		game.AddPlayer(player.NewPlayerShowdown(game, player.NewAI()))
		game.AddPlayer(player.NewPlayerShowdown(game, player.NewAI()))
		game.AddPlayer(player.NewPlayerShowdown(game, player.NewAI()))

		game.Start()
	}

	if input == "uno" {
		game := domain.NewGame(gameType.NewUno())

		game.AddPlayer(player.NewPlayerUno(game, player.NewHuman()))
		game.AddPlayer(player.NewPlayerUno(game, player.NewAI()))
		game.AddPlayer(player.NewPlayerUno(game, player.NewAI()))
		game.AddPlayer(player.NewPlayerUno(game, player.NewAI()))

		game.Start()
	}
}
