package sprite_handler

import (
	"world/internal/domain"
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
		return &domain.MoveCollisionStrategy{}
	}

	if spriteY.GetName() == sprite.TypeFire {
		return &domain.FailedCollisionStrategy{}
	}

	if spriteY.GetName() == sprite.TypeWater {
		return &domain.RemoveBothCollisionStrategy{}
	}

	if spriteY.GetName() == sprite.TypeHero {
		hero := spriteY.(*sprite.Hero)
		hero.MinusHP(10)

		if hero.GetHP() <= 0 {
			return &domain.RemoveBothCollisionStrategy{}
		}

		return &domain.RemoveFirstCollisionStrategy{}
	}

	return &domain.MoveCollisionStrategy{}
}
