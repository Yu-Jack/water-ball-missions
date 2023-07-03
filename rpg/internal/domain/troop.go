package domain

type troop struct {
	id    int
	roles []*role
	rpg   RPG
}

type Troop interface {
	GetID() int
	AddRole(r Role)
}

func NewTroop(id int, rpg RPG) Troop {
	return &troop{
		id:    id,
		roles: []*role{},
		rpg:   rpg,
	}
}

func (t *troop) GetID() int { return t.id }
func (t *troop) GetAllRole() []Role {
	var available []Role

	for _, r := range t.roles {
		r := r
		if r.hp > 0 {
			available = append(available, r)
		}
	}

	return available
}

func (t *troop) AddRole(r Role) {
	role := r.(*role)
	role.rpg = t.rpg
	role.ID = len(t.roles)
	role.troopID = t.id

	t.roles = append(t.roles, role)
}
