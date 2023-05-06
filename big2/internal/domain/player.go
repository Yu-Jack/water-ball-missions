package domain

import (
	"fmt"

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

func (p *player) HandSize() int {
	return p.hand.Size()
}

func (p *player) Name() {
	p.name = p.inputStrategy.Name()
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
		isMatchRule bool
		invalid     bool
		cp          cardPattern.Pattern
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

		if isMatchRule, invalid = p.checkRules(cp); isMatchRule {
			if invalid {
				continue
			} else {
				break
			}
		}

		if invalid = p.compareWithTable(cp); invalid {
			continue
		} else {
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

func (p *player) checkRules(cp cardPattern.Pattern) (isMatchRule, invalid bool) {
	// 第一回合玩家，一定要出梅花三
	if p.big2.TopPlay == nil && p.big2.Round == 1 {
		for _, c := range cp.GetCards() {
			if c.IsClub3() {
				return true, false
			}
		}

		// 沒有梅花三
		return true, true
	}

	// 每回合首發玩家，直接出牌
	if p.big2.TopPlay == nil {
		return true, false
	}

	return false, false
}

func (p *player) compareWithTable(cp cardPattern.Pattern) bool {
	compareResult := p.big2.Comparer.Compare(cp, p.big2.TopPlay)

	if compareResult == card.CompareResultSmaller {
		// 發現手牌太小，重出
		return true
	} else if compareResult == card.CompareResultBigger {
		// 手牌大於頂牌，出牌！
		return false
	} else {
		// 不同牌型比對，不合法
		return true
	}
}
