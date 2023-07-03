package domain

import (
	"fmt"
	"math/rand"
	"strconv"
)

type PlayerName string

type Player interface {
	Naming()
	GainPoint()
	PickCard()
	ExchangeHand(exchangee Player)
	MakeExchangeHandDecision()
	DrawCard(deck *Deck)
	GetName() PlayerName
	GetHands() *Hand
	SetHands(hand *Hand)
	GetPoint() int
	TakeTurn1(player Player)
	TakeTurn2(player Player)
}

type exchangeHand struct {
	exchangee Player
	countdown int
}

type player struct {
	name         PlayerName
	point        int
	hand         *Hand
	exchangeAble bool
	showdown     *Showdown
	exchangeHand *exchangeHand
}

func (p *player) GetName() PlayerName {
	return p.name
}

func (p *player) GetPoint() int {
	return p.point
}

func (p *player) GainPoint() {
	p.point++
}

func (p *player) GetHands() *Hand {
	return p.hand
}

func (p *player) SetHands(hand *Hand) {
	p.hand = hand
}

func (p *player) ExchangeHand(exchangee Player) {
	p1Hand, p2Hand := p.GetHands(), exchangee.GetHands()

	p.SetHands(p2Hand)
	exchangee.SetHands(p1Hand)
}

func (p *player) DrawCard(deck *Deck) {
	p.hand.AddCard(deck.DrawCard())
}

// TakeTurn1 is a template method, so you should bring the instance player to this method.
func (p *player) TakeTurn1(player Player) {
	player.MakeExchangeHandDecision()
}

// TakeTurn2 is a template method, so you should bring the instance player to this method.
func (p *player) TakeTurn2(player Player) {
	player.PickCard()
}

type HumanPlayer struct {
	player
}

type AIPlayer struct {
	player
}

func NewHumanPlayer(showdown *Showdown) Player {
	return &HumanPlayer{
		player: player{
			name:         "",
			point:        0,
			hand:         &Hand{},
			exchangeAble: true,
			showdown:     showdown,
		},
	}
}

func (p *HumanPlayer) Naming() {
	var input string
	fmt.Print("What is your name?: ")
	fmt.Scanln(&input)
	fmt.Print("Hello, " + input + "!\n")
	p.name = PlayerName(input)
}

func (p *HumanPlayer) PickCard() {
	fmt.Println("your hand:")
	p.hand.List()

	var input string
	fmt.Print("Which card do you want to pick? (card number): ")
	fmt.Scanln(&input)
	order, _ := strconv.Atoi(input)

	p.showdown.PutCardOnTable(p.hand.PickCard(int(order)), p.name)
}

func (p *HumanPlayer) MakeExchangeHandDecision() {
	if !p.exchangeAble {
		p.exchangeHand.countdown--
		if p.exchangeHand.countdown == 0 {
			p.ExchangeHand(p.exchangeHand.exchangee)
		}
		return
	}

	var input string

	for input != "y" && input != "n" {
		if input != "" {
			fmt.Println("Invalid input")
		}

		fmt.Print("Do you want to exchange your hand? (y/n): ")
		fmt.Scanln(&input)
	}

	if input == "n" {
		return
	}

	fmt.Print("Who do you want to exchange your hand with? (HumanPlayer name): ")
	fmt.Scanln(&input)

	for _, player := range p.showdown.players {
		fmt.Println(player.GetName(), PlayerName(input))
		if player.GetName() == PlayerName(input) {
			p.exchangeAble = false
			p.ExchangeHand(player)
			p.exchangeHand = &exchangeHand{
				exchangee: player,
				countdown: 3,
			}
		}
	}

	if p.exchangeAble {
		fmt.Println("Invalid name")
		return
	}
}

func NewAIPlayer(showdown *Showdown) Player {
	return &AIPlayer{
		player: player{
			name:         "",
			point:        0,
			hand:         &Hand{},
			exchangeAble: true,
			showdown:     showdown,
		},
	}
}

func (p *AIPlayer) Naming() {
	name := fmt.Sprintf("%d", rand.Intn(100000))
	p.name = PlayerName(name)
	fmt.Print("Hello, " + name + "!\n")
}

func (p *AIPlayer) PickCard() {
	// always pick first card for AI Player
	p.showdown.PutCardOnTable(p.hand.PickCard(0), p.name)
}

func (p *AIPlayer) MakeExchangeHandDecision() {
	if !p.exchangeAble {
		p.exchangeHand.countdown--
		if p.exchangeHand.countdown == 0 {
			p.ExchangeHand(p.exchangeHand.exchangee)
		}
		return
	}

	exchangeAble := []bool{true, false}[rand.Intn(2)]

	if !exchangeAble {
		return
	}

	p.exchangeAble = false
	order := rand.Intn(len(p.showdown.players))
	targetPlayer := p.showdown.players[order]
	p.ExchangeHand(targetPlayer)
	p.exchangeHand = &exchangeHand{
		exchangee: targetPlayer,
		countdown: 3,
	}

	fmt.Printf("AIPlayer: %s exchanged hand with %s \n", string(p.GetName()), string(targetPlayer.GetName()))
}
