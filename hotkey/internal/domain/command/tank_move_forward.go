package command

import "hotkey/internal/domain"

type TankMoveForward struct {
	tank *domain.Tank
}

func NewTankMoveForward(tank *domain.Tank) domain.Command {
	return &TankMoveForward{tank}
}

func (t *TankMoveForward) Execute() {
	t.tank.MoveForward()
}

func (t *TankMoveForward) Undo() {
	t.tank.MoveBackward()
}
