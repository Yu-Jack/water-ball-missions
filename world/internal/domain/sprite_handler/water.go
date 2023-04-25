package sprite_handler

import (
	"world/internal/domain"
	collisionStrategy "world/internal/domain/collision_strategy"
	"world/internal/domain/sprite"
)

type waterHandler struct{}

func NewWaterHandler() domain.SpriteHandler {
	return &waterHandler{}
}

func (wh *waterHandler) Match(s sprite.Sprite) bool {
	return s.GetName() == sprite.TypeWater
}

func (wh *waterHandler) Collision(spriteX, spriteY sprite.Sprite) domain.CollisionStrategy {
	if spriteY == nil {
		return &collisionStrategy.MoveCollisionStrategy{}
	}

	if spriteY.GetName() == sprite.TypeWater {
		return &collisionStrategy.FailedCollisionStrategy{}
	}

	if spriteY.GetName() == sprite.TypeFire {
		return &collisionStrategy.RemoveBothCollisionStrategy{}
	}

	if spriteY.GetName() == sprite.TypeHero {
		hero := spriteY.(*sprite.Hero)
		hero.PlusHP(10)
		return &collisionStrategy.RemoveFirstCollisionStrategy{}
	}

	return &collisionStrategy.MoveCollisionStrategy{}
}
