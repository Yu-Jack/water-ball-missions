package domain

type Macro struct {
	commands []Command
}

func NewMarco(commands []Command) Command {
	return &Macro{commands: commands}
}

func (m *Macro) Execute() {
	for _, c := range m.commands {
		c.Execute()
	}
}

func (m *Macro) Undo() {
	for i := len(m.commands) - 1; i >= 0; i-- {
		m.commands[i].Undo()
	}
}
