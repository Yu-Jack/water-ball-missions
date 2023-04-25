package sprite_handler

import (
	"world/internal/domain"
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
		return &domain.MoveCollisionStrategy{}
	}

	if spriteY.GetName() == sprite.TypeWater {
		return &domain.FailedCollisionStrategy{}
	}

	if spriteY.GetName() == sprite.TypeFire {
		return &domain.RemoveBothCollisionStrategy{}
	}

	if spriteY.GetName() == sprite.TypeHero {
		hero := spriteY.(*sprite.Hero)
		hero.PlusHP(10)
		return &domain.RemoveFirstCollisionStrategy{}
	}

	return &domain.MoveCollisionStrategy{}
}
