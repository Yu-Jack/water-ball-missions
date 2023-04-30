package domain

import "fmt"

type Telecom struct{}

func NewTelecom() *Telecom {
	return &Telecom{}
}

func (t *Telecom) Connect() {
	fmt.Println("Connect")
}

func (t *Telecom) Disconnect() {
	fmt.Println("Disconnect")
}
