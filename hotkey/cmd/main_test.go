package main

import (
	"fmt"

	"hotkey/internal/domain"
	"hotkey/internal/domain/command"
)

func Example() {
	controller := domain.NewMainController()
	keyboard := domain.NewKeyboard()

	controller.AddKeyboard(keyboard)
	keyboard.AddController(controller)

	tank := domain.NewTank()
	telecom := domain.NewTelecom()

	controller.Undo()
	controller.Redo()

	controller.Keyboard.Register(
		"z",
		command.NewTelecomConnect(telecom),
	)

	controller.Keyboard.Register(
		"a",
		command.NewTankMoveForward(tank),
	)

	controller.Keyboard.Press("z")
	controller.Keyboard.Press("a")
	controller.Keyboard.Reset()
	controller.Keyboard.Press("a")

	controller.Keyboard.Register(
		"h",
		command.NewTankMoveForward(tank),
	)
	controller.Keyboard.Register(
		"z",
		domain.NewMarco([]domain.Command{
			command.NewTankMoveBackward(tank),
			command.NewTankMoveForward(tank),
			command.NewTankMoveBackward(tank),
			command.NewTelecomConnect(telecom),
		}),
	)
	controller.Keyboard.Press("z")
	fmt.Println("Undo:")
	controller.Undo()
	fmt.Println("Redo:")
	controller.Redo()
	fmt.Println("Undo:")
	controller.Undo()
	controller.Undo()
	controller.Undo()
	controller.Keyboard.Press("h")

	// Output:
	// Nothing to undo.
	// Nothing to redo.
	// Connect
	// MoveTankForward
	// hotkey is not found.
	// MoveTankBackward
	// MoveTankForward
	// MoveTankBackward
	// Connect
	// Undo:
	// Disconnect
	// MoveTankForward
	// MoveTankBackward
	// MoveTankForward
	// Redo:
	// MoveTankBackward
	// MoveTankForward
	// MoveTankBackward
	// Connect
	// Undo:
	// Disconnect
	// MoveTankForward
	// MoveTankBackward
	// MoveTankForward
	// MoveTankBackward
	// Disconnect
	// MoveTankForward
}
