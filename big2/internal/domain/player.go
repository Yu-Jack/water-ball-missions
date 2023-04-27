package domain

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"

	"big2/internal/domain/card"
	cardPattern "big2/internal/domain/card_pattern"
)

type InputStrategy interface {
	Input() (orders []int, pass bool)
	Name() string
}

type Player interface {
	Name()
	GetName() string
	Play()
	AddCard(card *card.Card)
	HandSize() int
	IsClub3() bool
}

type player struct {
	name          string
	big2          *Big2
	hand          *Hand
	inputStrategy InputStrategy
}

func NewPlayer(big2 *Big2, inputStrategy InputStrategy) Player {
	return &player{
		big2:          big2,
		inputStrategy: inputStrategy,
		hand:          NewHand(),
	}
}

func (p *player) IsClub3() bool {
	for _, c := range p.hand.cards {
		if c.IsClub3() {
			return true
		}
	}
	return false
}

func (p *player) AddCard(card *card.Card) {
	p.hand.AddCard(card)
}

func (p *player) GetName() string {
	return p.name
}

func (p *player) Play() {
	p.hand.Sort() // 整理一下手牌，方便閱讀
	fmt.Printf("輪到 %s 了\n", p.GetName())
	p.hand.List()

	orders, pass := p.inputStrategy.Input()
	for pass && p.big2.TopPlay == nil {
		fmt.Println("你不能在新的回合中喊 PASS")
		orders, pass = p.inputStrategy.Input()
	}

	var (
		invalid bool
		cp      cardPattern.Pattern
	)

	for !pass {

		if invalid {
			fmt.Println("此牌型不合法，請再嘗試一次。")
			orders, pass = p.inputStrategy.Input()
		}

		cp = p.big2.Matcher.Match(p.hand.PickCard(orders))

		if cp == nil {
			invalid = true
			continue
		}

		// 每回合首發玩家，直接出牌
		if p.big2.TopPlay == nil {

			// 第一回合玩家，一定要出梅花三
			if p.big2.Round == 1 {
				invalid = true

				for _, c := range cp.GetCards() {
					if c.IsClub3() {
						invalid = false
						break
					}
				}

				if invalid {
					continue
				}
			}

			// 其餘回合直接出牌
			break
		}

		compareResult := p.big2.Comparer.Compare(cp, p.big2.TopPlay)

		// 不同牌型比對，不合法
		if compareResult == card.CompareResultInvalid {
			invalid = true
			continue
		}

		// 發現手牌太小，重出
		if compareResult == card.CompareResultSmaller {
			invalid = true
			continue
		}

		// 手牌大於頂牌，出牌！
		if compareResult == card.CompareResultBigger {
			invalid = false
			break
		}
	}

	if pass {
		p.big2.PassCount++
		fmt.Printf("玩家 %s PASS\n", p.GetName())
		return
	}

	if !invalid {
		p.hand.Remove(orders)
		p.big2.TopPlay = cp
		p.big2.TopPlayer = p
		p.big2.PassCount = 0
		fmt.Printf("玩家 %s 打出了 %s\n", p.GetName(), cp.String())
	}
}

func (p *player) HandSize() int {
	return p.hand.Size()
}

func (p *player) Name() {
	p.name = p.inputStrategy.Name()
}

type human struct{}

func NewHumanStrategy() InputStrategy {
	return &human{}
}

func (h human) Input() ([]int, bool) {
	reader := bufio.NewReader(os.Stdin)
	var input string
	var inputs []string

	for input == "" {
		fmt.Print("請輸入牌號：")
		input, _ = reader.ReadString('\n')
		input = strings.TrimSuffix(input, "\n")
		inputs = strings.Split(input, " ")
	}

	if inputs[0] == "-1" {
		return []int{}, true
	}

	orders := make([]int, 0, len(inputs))
	for _, input := range inputs {
		integer, _ := strconv.Atoi(input)
		orders = append(orders, integer)
	}

	return orders, false
}

func (h human) Name() string {
	var name string

	fmt.Println("請輸入你的名字：")
	_, _ = fmt.Scan(&name)
	fmt.Print("你好, " + name + "!\n")

	return name
}

type ai struct{}

func NewAIStrategy() InputStrategy {
	return &ai{}
}

func (a ai) Input() ([]int, bool) {
	//TODO implement me
	panic("implement me")
}

func (a ai) Name() string {
	name := fmt.Sprintf("%d", rand.Intn(100000))
	fmt.Print("Hello, " + name + "!\n")
	return name
}
