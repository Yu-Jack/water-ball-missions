package command

import "hotkey/internal/domain"

type TelecomConnect struct {
	telecom *domain.Telecom
}

func NewTelecomConnect(telecom *domain.Telecom) domain.Command {
	return &TelecomConnect{telecom}
}

func (t *TelecomConnect) Execute() {
	t.telecom.Connect()
}

func (t *TelecomConnect) Undo() {
	t.telecom.Disconnect()
}
