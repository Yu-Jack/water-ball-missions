package domain

type troop struct {
	id    int
	roles []*role
	rpg   RPG
}

type Troop interface {
	CreateNewRole(
		name string,
		hp, mp, str int,
		state State,
		skills []Skill,
		strategy ActionStrategy,
	)
	AddRole(r Role)
}

func NewTroop(id int, rpg RPG) Troop {
	return &troop{
		id:    id,
		roles: []*role{},
		rpg:   rpg,
	}
}

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
	t.roles = append(t.roles, role)
}

func (t *troop) CreateNewRole(name string, hp, mp, str int, state State, skills []Skill, strategy ActionStrategy) {
	role := &role{
		name:           name,
		troopID:        t.id,
		ID:             len(t.roles),
		hp:             hp,
		mp:             mp,
		str:            str,
		extraStr:       0,
		skills:         skills,
		state:          state,
		actionAble:     true,
		actionStrategy: strategy,
		rpg:            t.rpg,
	}

	role.SetState(state)

	t.roles = append(t.roles, role)
}
