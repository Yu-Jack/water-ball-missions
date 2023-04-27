//go:generate go-enum -f=$GOFILE --nocase

package card

import (
	"fmt"
)

type Card struct {
	rank int
	suit int
}

// CompareResult
/*
ENUM(
Smaller = 0
Bigger = 1
Equal = 2
Invalid = 3
)
*/
type CompareResult int

const (
	Rank3 = iota + 1
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
	Rank2
)

const (
	SuitClub = iota + 1
	SuitDiamond
	SuitHeart
	SuitSpade
)

func NewCard(rank string, suit string) *Card {
	var r, s int

	if suit == "S" {
		s = SuitSpade
	}
	if suit == "D" {
		s = SuitDiamond
	}
	if suit == "H" {
		s = SuitHeart
	}
	if suit == "C" {
		s = SuitClub
	}
	if rank == "A" {
		r = RankA
	}
	if rank == "2" {
		r = Rank2
	}
	if rank == "3" {
		r = Rank3
	}
	if rank == "4" {
		r = Rank4
	}
	if rank == "5" {
		r = Rank5
	}
	if rank == "6" {
		r = Rank6
	}
	if rank == "7" {
		r = Rank7
	}
	if rank == "8" {
		r = Rank8
	}
	if rank == "9" {
		r = Rank9
	}
	if rank == "10" {
		r = Rank10
	}
	if rank == "J" {
		r = RankJ
	}
	if rank == "Q" {
		r = RankQ
	}
	if rank == "K" {
		r = RankK
	}

	return &Card{
		rank: r,
		suit: s,
	}
}

// Compare returns 0 (smaller), 1 (bigger)
func (c *Card) Compare(card *Card) CompareResult {
	if c.rank < card.rank {
		return CompareResultSmaller
	}

	if c.rank == card.rank {
		if c.suit < card.suit {
			return CompareResultSmaller
		}
	}

	return CompareResultBigger
}

// CompareRank returns 0 (smaller), 1 (bigger), 2(equal)
func (c *Card) CompareRank(card *Card) CompareResult {
	if c.rank < card.rank {
		return CompareResultSmaller
	}
	if c.rank > card.rank {
		return CompareResultBigger
	}
	return CompareResultEqual
}

// DiffRank return diff rank
func (c *Card) DiffRank(card *Card) int {
	return c.rank - card.rank
}

func (c *Card) IsClub3() bool {
	return c.rank == Rank3 && c.suit == SuitClub
}

// for convenience of debugging
func (c *Card) String() string {
	rank := ""

	switch c.rank {
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
	switch c.suit {
	case SuitClub:
		suit = "C"
	case SuitDiamond:
		suit = "D"
	case SuitHeart:
		suit = "H"
	case SuitSpade:
		suit = "S"
	}

	return fmt.Sprintf("%s[%s]", suit, rank)
}
