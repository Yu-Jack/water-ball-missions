package domain

import (
	"fmt"
	"strings"
)

type Big2 struct {
	Matcher   CardPatternMatcher
	Comparer  CardPatternCompareHandler
	Players   []Player
	TopPlay   CardPattern
	TopPlayer Player
	Round     int
	PassCount int
}

func NewBig2(
	matcher CardPatternMatcher,
	comparer CardPatternCompareHandler,
) *Big2 {
	return &Big2{
		Round:     1,
		PassCount: 0,
		Matcher:   matcher,
		Comparer:  comparer,
	}
}

func (b *Big2) AddPlayer(player Player) {
	b.Players = append(b.Players, player)
}

func (b *Big2) Start() {
	for _, player := range b.Players {
		player.Name()
	}

	// TODO: 需要封裝
	input := "S[8] S[9] S[3] D[J] S[7] H[3] C[2] C[9] H[2] D[7] S[K] C[6] C[3] D[4] H[7] C[A] D[A] D[K] H[4] D[8] C[4] H[10] H[A] S[10] H[Q] H[5] S[4] D[5] H[9] H[8] C[10] S[6] S[A] D[3] S[5] D[9] D[Q] H[K] C[Q] H[J] D[10] S[2] H[6] C[K] S[J] C[7] S[Q] D[6] D[2] C[J] C[8] C[5]"
	inputs := strings.Split(input, " ")
	cards := make([]*Card, 0)
	for _, input := range inputs {
		input = strings.Replace(input, "[", " ", -1)
		input = strings.Replace(input, "]", "", -1)

		var suit string
		var rank string
		_, _ = fmt.Sscanf(input, "%s %s", &suit, &rank)
		cards = append(cards, NewCard(rank, suit))

	}

	deck := NewDeck(cards)

	for deck.Size() > 0 {
		for _, player := range b.Players {
			player.AddCard(deck.Deal())
		}
	}

	startPlayer := -1
	for i, _ := range b.Players {
		// TODO: 需要讓梅花三的先出
		startPlayer = i
		break
	}

	winner := ""

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
