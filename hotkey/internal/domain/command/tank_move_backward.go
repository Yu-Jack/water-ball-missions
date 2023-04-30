package command

import "hotkey/internal/domain"

type TankMoveBackward struct {
	tank *domain.Tank
}

func NewTankMoveBackward(tank *domain.Tank) domain.Command {
	return &TankMoveBackward{tank}
}

func (t *TankMoveBackward) Execute() {
	t.tank.MoveBackward()
}

func (t *TankMoveBackward) Undo() {
	t.tank.MoveForward()
}
