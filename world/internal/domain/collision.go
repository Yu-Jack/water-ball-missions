package domain

import "fmt"

type collisionStrategy struct{}
type RemoveFirstCollisionStrategy struct{ collisionStrategy }
type RemoveBothCollisionStrategy struct{ collisionStrategy }
type MoveCollisionStrategy struct{ collisionStrategy }
type FailedCollisionStrategy struct{ collisionStrategy }

type CollisionStrategy interface {
	Action(world *World, p1, p2 int)
}

func (s *RemoveFirstCollisionStrategy) Action(world *World, p1, p2 int) {
	world.RemoveSprite(p1)
}

func (s *RemoveBothCollisionStrategy) Action(world *World, p1, p2 int) {
	world.RemoveSprite(p1)
	world.RemoveSprite(p2)
}

func (s *MoveCollisionStrategy) Action(world *World, p1, p2 int) {
	world.MoveSprite(p1, p2)
}

func (s *FailedCollisionStrategy) Action(world *World, p1, p2 int) {
	fmt.Println("failed")
}
