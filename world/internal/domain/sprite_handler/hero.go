package sprite_handler

import (
	"world/internal/domain"
	collisionStrategy "world/internal/domain/collision_strategy"
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
		return &collisionStrategy.MoveCollisionStrategy{}
	}

	if spriteY.GetName() == sprite.TypeHero {
		return &collisionStrategy.FailedCollisionStrategy{}
	}

	hero := spriteX.(*sprite.Hero)

	if spriteY.GetName() == sprite.TypeWater {
		hero.PlusHP(10)
	}

	if spriteY.GetName() == sprite.TypeFire {
		hero.MinusHP(10)
	}

	if hero.GetHP() <= 0 {
		return &collisionStrategy.RemoveBothCollisionStrategy{}
	}

	return &collisionStrategy.MoveCollisionStrategy{}
}
