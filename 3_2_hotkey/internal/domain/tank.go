package domain

import "fmt"

type Tank struct{}

func NewTank() *Tank {
	return &Tank{}
}

func (t *Tank) MoveForward() {
	fmt.Println("MoveTankForward")
}

func (t *Tank) MoveBackward() {
	fmt.Println("MoveTankBackward")
}
