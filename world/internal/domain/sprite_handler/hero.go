package sprite_handler

import (
	"world/internal/domain"
	"world/internal/domain/sprite"
)

type heroHandler struct{}

func NewHeroHandler() domain.SpriteHandler {
	return &heroHandler{}
}

func (hh *heroHandler) Match(s sprite.Sprite) bool {
	return s.GetName() == sprite.TypeHero
}

func (hh *heroHandler) Collision(spriteX, spriteY sprite.Sprite) domain.CollisionStrategy {
	if spriteY == nil {
		return &domain.MoveCollisionStrategy{}
	}

	if spriteY.GetName() == sprite.TypeHero {
		return &domain.FailedCollisionStrategy{}
	}

	hero := spriteX.(*sprite.Hero)

	if spriteY.GetName() == sprite.TypeWater {
		hero.PlusHP(10)
	}

	if spriteY.GetName() == sprite.TypeFire {
		hero.MinusHP(10)
	}

	if hero.GetHP() <= 0 {
		return &domain.RemoveFirstCollisionStrategy{}
	}

	return &domain.MoveCollisionStrategy{}
}
