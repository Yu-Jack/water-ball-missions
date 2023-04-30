package command

import "hotkey/internal/domain"

type TelecomDisconnect struct {
	telecom *domain.Telecom
}

func NewTelecomDisconnect(telecom *domain.Telecom) domain.Command {
	return &TelecomDisconnect{telecom}
}

func (t *TelecomDisconnect) Execute() {
	t.telecom.Disconnect()
}

func (t *TelecomDisconnect) Undo() {
	t.telecom.Connect()
}
