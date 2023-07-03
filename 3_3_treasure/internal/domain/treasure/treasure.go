//go:generate go-enum -f=$GOFILE --nocase

package treasure

import (
	"math/rand"

	"treasure/internal/domain"
)

// Content
/*
ENUM(
DokodemoDoor
Poison
KingsRock
SuperStar
HealingPotion
AcceleratingPotion
DevilFruit
)
*/
type Content string

type treasure struct {
	position domain.Position
	content  Content
}

func (t *treasure) GetPosition() domain.Position {
	return t.position
}

func (t *treasure) GetContent() string {
	return t.content.String()
}

var (
	contents = []Content{
		ContentSuperStar,
		ContentPoison,
		ContentAcceleratingPotion,
		ContentHealingPotion,
		ContentDevilFruit,
		ContentKingsRock,
		ContentDokodemoDoor,
	}
	probs = []float64{
		0.1,
		0.25,
		0.2,
		0.15,
		0.1,
		0.1,
		0.1,
	}
)

func GenerateRandomTreasure(p domain.Position) domain.Treasure {
	var content Content

	r := rand.Float64()
	var acc float64
	for i, p := range probs {
		acc += p
		if r < acc {
			content = contents[i]
			break
		}
	}

	if content == ContentSuperStar {
		return NewSuperStar(p)
	}

	if content == ContentHealingPotion {
		return NewHealingPotion(p)
	}

	if content == ContentAcceleratingPotion {
		return NewAcceleratingPotion(p)
	}

	if content == ContentDevilFruit {
		return NewDevilFruit(p)
	}

	if content == ContentDokodemoDoor {
		return NewDokodemoDoor(p)
	}

	if content == ContentKingsRock {
		return NewKingsRock(p)
	}

	if content == ContentPoison {
		return NewPoison(p)
	}

	return nil
}
