package domain

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

type InputStrategy interface {
	Input() (orders []int, pass bool)
	Name() string
}

type Player interface {
	Name()
	GetName() string
	Play()
	AddCard(card *Card)
	HandSize() int
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

func (p *player) AddCard(card *Card) {
	p.hand.AddCard(card)
}

func (p *player) GetName() string {
	return p.name
}

func (p *player) Play() {
	p.hand.Sort()
	fmt.Printf("輪到 %s 了\n", p.GetName())
	p.hand.List()

	orders, pass := p.inputStrategy.Input()
	for pass && p.big2.TopPlay == nil {
		fmt.Println("你不能在新的回合中喊 PASS")
		orders, pass = p.inputStrategy.Input()
	}

	for !pass {
		cards := p.hand.PickCard(orders)

		cp := p.big2.Matcher.Match(cards)
		if cp == nil {
			fmt.Println("此牌型不合法，請再嘗試一次。")
			orders, pass = p.inputStrategy.Input()
			continue
		}

		// 每回合首發玩家，直接出牌
		if p.big2.TopPlay == nil {
			p.hand.Remove(orders)
			p.big2.TopPlay = cp
			p.big2.TopPlayer = p
			fmt.Printf("玩家 %s 打出了 %s\n", p.GetName(), cp.String())
			break
		}

		compareResult := p.big2.Comparer.Compare(cp, p.big2.TopPlay)

		// 發現手牌太小，重出
		if compareResult == CompareResultSmaller {
			fmt.Println("此牌型不合法，請再嘗試一次。")
			orders, pass = p.inputStrategy.Input()
			continue
		}

		// 手牌大於頂牌，出牌！
		if compareResult == CompareResultBigger {
			p.hand.Remove(orders)
			p.big2.TopPlay = cp
			p.big2.TopPlayer = p
			p.big2.PassCount = 0
			fmt.Printf("玩家 %s 打出了 %s\n", p.GetName(), cp.String())
			break
		}
	}

	if pass {
		p.big2.PassCount++
		fmt.Printf("玩家 %s PASS\n", p.GetName())
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
