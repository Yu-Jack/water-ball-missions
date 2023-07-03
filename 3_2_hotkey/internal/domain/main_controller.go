package domain

import "fmt"

type MainController struct {
	Keyboard                *Keyboard
	commandExecuteHistories []Command
	commandUndoHistories    []Command
}

func NewMainController() *MainController {
	return &MainController{
		commandExecuteHistories: []Command{},
		commandUndoHistories:    []Command{},
	}
}

func (c *MainController) AddKeyboard(keyboard *Keyboard) {
	c.Keyboard = keyboard
}

func (c *MainController) AddExecuteHistory(command Command) {
	c.commandExecuteHistories = append(c.commandExecuteHistories, command)
}

func (c *MainController) Undo() {
	if len(c.commandExecuteHistories) == 0 {
		fmt.Println("Nothing to undo.")
		return
	}

	pop := c.commandExecuteHistories[len(c.commandExecuteHistories)-1]
	c.commandExecuteHistories = c.commandExecuteHistories[:len(c.commandExecuteHistories)-1]
	pop.Undo()

	c.commandUndoHistories = append(c.commandUndoHistories, pop)
}

func (c *MainController) Redo() {
	if len(c.commandUndoHistories) == 0 {
		fmt.Println("Nothing to redo.")
		return
	}

	pop := c.commandUndoHistories[len(c.commandUndoHistories)-1]
	c.commandUndoHistories = c.commandUndoHistories[:len(c.commandUndoHistories)-1]
	pop.Execute()

	c.commandExecuteHistories = append(c.commandExecuteHistories, pop)
}
