package domain

import (
	"fmt"
)

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
	end, winner := r.End()

	defer func() {
		fmt.Println(winner)
	}()

	for !end {
		for _, t := range r.troops {
			for i := 0; i < len(t.roles); i++ {
				role := t.roles[i]

				if role.GetHp() <= 0 {
					continue
				}

				role.state.Do()
				role.state.CountDown()
				role.TakeAction()

				if end, winner = r.End(); end {
					return
				}
			}
		}
	}
}

func (r *rpg) End() (bool, string) {
	for i := 0; i < len(r.troops); i++ {
		t := r.troops[i]
		atLeastOneAlive := false

		for j := 0; j < len(t.roles); j++ {
			role := t.roles[j]

			if role.GetName() == "英雄" && role.GetHp() < 0 {
				return true, "你失敗了！"
			}

			if role.GetHp() > 0 {
				atLeastOneAlive = true
				break
			}
		}

		if atLeastOneAlive {
			continue
		} else {
			return true, "你獲勝了！"
		}
	}

	return false, ""
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