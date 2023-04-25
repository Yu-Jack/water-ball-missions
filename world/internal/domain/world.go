package domain

import (
	"fmt"
	"strconv"

	"world/internal/domain/sprite"
)

type World struct {
	length   int
	position []sprite.Sprite
	handler  Handler
}

func NewWorld(handler Handler) *World {
	length := 30

	return &World{
		length:   length,
		position: make([]sprite.Sprite, length, length),
		handler:  handler,
	}
}

func (w *World) Start() {

	for i := 0; i < 10; i++ {
		if i%3 == 0 {
			w.position[i] = sprite.NewWater()
		}
		if i%3 == 1 {
			w.position[i] = sprite.NewFire()
		}
		if i%3 == 2 {
			w.position[i] = sprite.NewHero()
		}
	}

	for {
		fmt.Println("=========")
		for i, p := range w.position {
			fmt.Println(i, p)
		}
		fmt.Println("=========")

		var position1 string
		var position2 string
		fmt.Print("How do you want to move?: ")
		_, _ = fmt.Scanf("%s %s", &position1, &position2)

		p1, _ := strconv.Atoi(position1)
		p2, _ := strconv.Atoi(position2)

		x, y := w.position[p1], w.position[p2]

		w.handler.Collision(x, y).Action(w, p1, p2)
	}
}

func (w *World) RemoveSprite(order int) {
	w.position[order] = nil
}

func (w *World) MoveSprite(p1, p2 int) {
	w.position[p2] = w.position[p1]
	w.position[p1] = nil
}
