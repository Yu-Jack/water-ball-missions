package skill

import "rpg/internal/domain"

type onePunchHandler struct {
	handler OnePunchHandler
	next    OnePunchHandler
}

type OnePunchHandler interface {
	match(current, target domain.Role) bool
	execute(current, target domain.Role)
}

func NewOnePunchHandler(handler, next OnePunchHandler) OnePunchHandler {
	return &onePunchHandler{
		handler: handler,
		next:    next,
	}
}

func (o onePunchHandler) execute(current, target domain.Role) {
	if o.handler.match(current, target) {
		o.handler.execute(current, target)
	} else if o.next != nil {
		o.next.execute(current, target)
	}
}

// for compiler
func (o onePunchHandler) match(current, target domain.Role) bool { return false }
