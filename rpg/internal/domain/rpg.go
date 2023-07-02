package domain

type rpg struct {
	troops []*troop
}

func (r *rpg) GetRole(troopID, roleID int) Role {
	for _, t := range r.troops {
		if t.id == troopID {
			for index, role := range t.roles {
				if index == roleID {
					return role
				}
			}
		}
	}

	return nil
}

type RPG interface {
	GetAllEnemies(troopID int) []Role
	GetAllAllies(troopID int) []Role
	GetRole(troopID, roleID int) Role
	GetAllyTroop(id int) Troop
	Start()
	AddTroop(troop Troop)
}

func NewClientRPG() RPG {
	return &rpg{}
}

func (r *rpg) Start() {
	for !r.End() {
		for _, t := range r.troops {
			for i := 0; i < len(t.roles); i++ {
				role := t.roles[i]

				if role.GetHp() <= 0 {
					continue
				}

				role.state.Do()
				role.state.CountDown()
				role.TakeAction()
			}
		}
	}
}

func (r *rpg) End() bool {
	// TODO: HERO dies or troops all die
	return false
}

func (r *rpg) AddTroop(t Troop) {
	r.troops = append(r.troops, t.(*troop))
}

func (r *rpg) GetAllEnemies(troopID int) []Role {

	for _, t := range r.troops {
		if t.id != troopID {
			return t.GetAllRole()
		}
	}

	return []Role{}
}

func (r *rpg) GetAllAllies(troopID int) []Role {

	for _, t := range r.troops {
		if t.id == troopID {
			return t.GetAllRole()
		}
	}

	return []Role{}
}

func (r *rpg) GetAllyTroop(troopID int) Troop {
	var troop Troop

	for _, t := range r.troops {
		t := t
		if t.id == troopID {
			troop = t
			break
		}
	}

	return troop
}
