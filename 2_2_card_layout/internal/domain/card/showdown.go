package card

import (
	"fmt"

	"card.layout/internal/domain"
)

type showdown struct {
	rank int
	suit int
}

const (
	Rank2 = iota + 1
	Rank3
	Rank4
	Rank5
	Rank6
	Rank7
	Rank8
	Rank9
	Rank10
	RankJ
	RankQ
	RankK
	RankA
)

const (
	SuitClub = iota + 1
	SuitDiamond
	SuitHeart
	SuitSpade
)

var (
	Ranks = []int{
		Rank2, Rank3, Rank4, Rank5,
		Rank6, Rank7, Rank8, Rank9,
		Rank10, RankJ, RankQ, RankK, RankA,
	}
	Suits = []int{
		SuitClub, SuitDiamond, SuitHeart, SuitSpade,
	}
)

func NewShowdown() domain.Card {
	return &showdown{}
}

// Compare returns 0 (smaller), 1 (bigger)
func (s *showdown) Compare(card domain.Card) int {
	c := card.(*showdown)

	if s.rank < c.rank {
		return 0
	}

	if s.rank == c.rank {
		if s.suit < c.suit {
			return 0
		}
	}

	return 1
}

func (s *showdown) GenerateDeck() []domain.Card {
	cards := make([]domain.Card, 0, 52)

	for _, r := range Ranks {
		for _, s := range Suits {
			cards = append(cards, &showdown{rank: r, suit: s})
		}
	}

	return cards
}

// for convenience of debugging
func (s *showdown) String() string {
	rank := ""

	switch s.rank {
	case Rank2:
		rank = "2"
	case Rank3:
		rank = "3"
	case Rank4:
		rank = "4"
	case Rank5:
		rank = "5"
	case Rank6:
		rank = "6"
	case Rank7:
		rank = "7"
	case Rank8:
		rank = "8"
	case Rank9:
		rank = "9"
	case Rank10:
		rank = "10"
	case RankJ:
		rank = "J"
	case RankQ:
		rank = "Q"
	case RankK:
		rank = "K"
	case RankA:
		rank = "A"
	}

	suit := ""
	switch s.suit {
	case SuitClub:
		suit = "梅花"
	case SuitDiamond:
		suit = "菱形"
	case SuitHeart:
		suit = "方塊"
	case SuitSpade:
		suit = "黑桃"
	}

	return fmt.Sprintf("%s%s", suit, rank)
}
