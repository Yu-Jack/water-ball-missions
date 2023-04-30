package domain

import "fmt"

type Command interface {
	Execute()
	Undo()
}

type Keyboard struct {
	hotkey     map[string]Command
	controller *MainController
}

func NewKeyboard() *Keyboard {
	return &Keyboard{
		hotkey: make(map[string]Command),
	}
}

func (k *Keyboard) AddController(controller *MainController) {
	k.controller = controller
}

func (k *Keyboard) Register(name string, command Command) {
	if name[0]-'a' < 0 || 'z'-name[0] > 25 {
		fmt.Printf("%s is not a valid hotkey name.\n", name)
		return
	}

	k.hotkey[name] = command
}

func (k *Keyboard) Reset() {
	k.hotkey = make(map[string]Command)
}

func (k *Keyboard) Press(key string) {
	if c, ok := k.hotkey[key]; ok {
		c.Execute()
		k.controller.AddExecuteHistory(c)
	} else {
		fmt.Println("hotkey is not found.")
	}
}
