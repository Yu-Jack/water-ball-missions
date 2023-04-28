package domain

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"big2/internal/domain/card"
	cardPattern "big2/internal/domain/card_pattern"
)

type Big2 struct {
	Matcher   cardPattern.Matcher
	Comparer  cardPattern.CompareHandler
	Players   []Player
	TopPlay   cardPattern.Pattern
	TopPlayer Player
	Round     int
	PassCount int
}

func NewBig2(
	matcher cardPattern.Matcher,
	comparer cardPattern.CompareHandler,
) *Big2 {
	return &Big2{
		Round:     1,
		PassCount: 0,
		TopPlay:   nil,
		TopPlayer: nil,
		Players:   []Player{},
		Matcher:   matcher,
		Comparer:  comparer,
	}
}

func (b *Big2) AddPlayer(player Player) {
	b.Players = append(b.Players, player)
}

func (b *Big2) generateDeck() *Deck {
	var input string

	reader := bufio.NewReader(os.Stdin)
	input, _ = reader.ReadString('\n')
	input = strings.TrimSuffix(input, "\n")

	inputs := strings.Split(input, " ")
	cards := make([]*card.Card, 0)
	for _, input := range inputs {
		input = strings.Replace(input, "[", " ", -1)
		input = strings.Replace(input, "]", "", -1)

		var suit string
		var rank string
		_, _ = fmt.Sscanf(input, "%s %s", &suit, &rank)
		cards = append(cards, card.NewCard(rank, suit))
	}

	return NewDeck(cards)
}

func (b *Big2) Start() {
	deck := &Deck{}

	for deck.Size() != 52 {
		deck = b.generateDeck()
	}

	for _, player := range b.Players {
		player.Name()
	}

	// deal cards
	for deck.Size() > 0 {
		for _, player := range b.Players {
			player.AddCard(deck.Deal())
		}
	}

	// choose who is first one
	startPlayer := -1
	for i, _ := range b.Players {
		if b.Players[i].IsClub3() {
			startPlayer = i
			break
		}
	}

	winner := ""

	fmt.Println("新的回合開始了。")

	for {
		order := startPlayer % 4
		currentPlayer := b.Players[order]
		currentPlayer.Play()

		if currentPlayer.HandSize() == 0 {
			winner = currentPlayer.GetName()
			break
		}

		if b.PassCount == 3 {
			for i, p := range b.Players {
				if p.GetName() == b.TopPlayer.GetName() {
					startPlayer = i
					break
				}
			}
			b.TopPlay = nil // clear top cards
			b.PassCount = 0
			b.Round++
			fmt.Println("新的回合開始了。")
			continue
		}

		startPlayer++
	}

	fmt.Printf("遊戲結束，遊戲的勝利者為 %s\n", winner)
}
