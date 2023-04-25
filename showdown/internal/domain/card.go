package domain

import "fmt"

type Card struct {
	rank int
	suit int
}

const (
	RANK2 = iota + 1
	RANK3
	RANK4
	RANK5
	RANK6
	RANK7
	RANK8
	RANK9
	RANK10
	RANKJ
	RANKQ
	RANKK
	RANKA
)

const (
	SuitClub = iota + 1
	SuitDiamond
	SuitHeart
	SuitSpade
)

func (c *Card) Bigger(card Card) bool {
	if c.rank < card.rank {
		return false
	}

	if c.rank == card.rank {
		if c.suit < card.suit {
			return false
		}
	}

	return true
}

// for convenience of debugging
func (c *Card) String() string {
	rank := ""

	switch c.rank {
	case RANK2:
		rank = "2"
	case RANK3:
		rank = "3"
	case RANK4:
		rank = "4"
	case RANK5:
		rank = "5"
	case RANK6:
		rank = "6"
	case RANK7:
		rank = "7"
	case RANK8:
		rank = "8"
	case RANK9:
		rank = "9"
	case RANK10:
		rank = "10"
	case RANKJ:
		rank = "J"
	case RANKQ:
		rank = "Q"
	case RANKK:
		rank = "K"
	case RANKA:
		rank = "A"
	}

	suit := ""
	switch c.suit {
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
