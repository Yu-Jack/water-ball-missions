package collision_handler

import (
	"world/internal/domain"
	collisionStrategy "world/internal/domain/collision_strategy"
	"world/internal/domain/sprite"
)

type heroHeroHandler struct{}

func NewHeroHeroHandler() domain.SpriteHandler {
	return &heroHeroHandler{}
}

func (hh *heroHeroHandler) Match(c1, c2 sprite.Sprite) bool {
	return c1.GetName() == sprite.TypeHero && c2.GetName() == sprite.TypeHero
}

func (hh *heroHeroHandler) Collision(c1, c2 sprite.Sprite) domain.CollisionStrategy {
	return &collisionStrategy.FailedCollisionStrategy{}
}
