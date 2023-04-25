package sprite_handler

import (
	"world/internal/domain"
	collisionStrategy "world/internal/domain/collision_strategy"
	"world/internal/domain/sprite"
)

type fireHandler struct{}

func NewFireHandler() domain.SpriteHandler {
	return &fireHandler{}
}

func (fh *fireHandler) Match(s sprite.Sprite) bool {
	return s.GetName() == sprite.TypeFire
}

func (fh *fireHandler) Collision(spriteX, spriteY sprite.Sprite) domain.CollisionStrategy {
	if spriteY == nil {
		return &collisionStrategy.MoveCollisionStrategy{}
	}

	if spriteY.GetName() == sprite.TypeFire {
		return &collisionStrategy.FailedCollisionStrategy{}
	}

	if spriteY.GetName() == sprite.TypeWater {
		return &collisionStrategy.RemoveBothCollisionStrategy{}
	}

	if spriteY.GetName() == sprite.TypeHero {
		hero := spriteY.(*sprite.Hero)
		hero.MinusHP(10)

		if hero.GetHP() <= 0 {
			return &collisionStrategy.RemoveBothCollisionStrategy{}
		}

		return &collisionStrategy.RemoveFirstCollisionStrategy{}
	}

	return &collisionStrategy.MoveCollisionStrategy{}
}
