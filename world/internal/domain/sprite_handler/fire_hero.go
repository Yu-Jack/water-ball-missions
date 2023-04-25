package sprite_handler

import (
	"fmt"

	"world/internal/domain"
	collisionStrategy "world/internal/domain/collision_strategy"
	"world/internal/domain/sprite"
)

type fireHeroHandler struct{}

func NewFireHeroHandler() domain.SpriteHandler {
	return &fireHeroHandler{}
}

func (fh *fireHeroHandler) Match(c1, c2 sprite.Sprite) bool {
	return (c1.GetName() == sprite.TypeFire && c2.GetName() == sprite.TypeHero) ||
		(c1.GetName() == sprite.TypeHero && c2.GetName() == sprite.TypeFire)
}

func (fh *fireHeroHandler) Collision(c1, c2 sprite.Sprite) domain.CollisionStrategy {

	fmt.Println("Fire")
	fmt.Println("Fire")
	fmt.Println("Fire")
	fmt.Println("Fire")

	// c1 是 fire
	if c1.GetName() == sprite.TypeFire {
		hero := c2.(*sprite.Hero)
		hero.MinusHP(10)
		if hero.GetHP() <= 0 {
			return &collisionStrategy.RemoveBothCollisionStrategy{}
		}

		return &collisionStrategy.RemoveFirstCollisionStrategy{}
	}

	// c1 是 hero
	if c1.GetName() == sprite.TypeHero {
		hero := c1.(*sprite.Hero)
		hero.MinusHP(10)

		if hero.GetHP() <= 0 {
			return &collisionStrategy.RemoveBothCollisionStrategy{}
		}

		return &collisionStrategy.MoveCollisionStrategy{}
	}

	return &collisionStrategy.FailedCollisionStrategy{}
}
