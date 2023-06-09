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
	GetAllRolesExcludeSelf(troopID, roleID int) []Role
	GetAllEnemies(troopID int) []Role
	GetAllAllies(troopID int) []Role
	GetAllAlliesExcludeSelf(troopID, roleID int) []Role
	GetRole(troopID, roleID int) Role
	GetAllyTroop(id int) Troop
	Start()
	AddTroop(troop Troop)
}

func NewClientRPG() RPG {
	return &rpg{}
}

func (r *rpg) startGame() string {
	end, winner := r.End()

	for !end {
		for _, t := range r.troops {
			for i := 0; i < len(t.roles); i++ {
				r.takeRound(t.roles[i])

				if end, winner = r.End(); end {
					return winner
				}
			}
		}
	}

	return winner
}

func (r *rpg) takeRound(role *role) {
	if role.GetHp() <= 0 {
		return
	}

	role.PrintInformation()
	role.state.Do()

	if role.GetHp() <= 0 {
		return
	}

	role.TakeAction()
	role.state.CountDown()
}

func (r *rpg) Start() {
	winner := r.startGame()
	fmt.Println(winner)
}

func (r *rpg) End() (bool, string) {
	for i := 0; i < len(r.troops); i++ {
		t := r.troops[i]
		atLeastOneAlive := false

		for j := 0; j < len(t.roles); j++ {
			role := t.roles[j]

			// 暫時寫死用 英雄 去判斷是人為控制的
			if role.GetName() == StrategyNameHero.String() && role.GetHp() <= 0 {
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

func (r *rpg) GetAllRolesExcludeSelf(troopID, roleID int) []Role {
	var rs []Role

	for _, t := range r.troops {
		for _, role := range t.roles {
			if role.GetID() == roleID && role.GetTroopID() == troopID {
				continue
			}
			if role.GetHp() > 0 {
				rs = append(rs, role)
			}
		}
	}

	return rs
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

func (r *rpg) GetAllAlliesExcludeSelf(troopID, roleID int) []Role {

	var roles []Role

	for _, t := range r.troops {
		if t.id != troopID {
			continue
		}

		for index, role := range t.roles {
			if index != roleID && role.GetHp() > 0 {
				roles = append(roles, role)
			}
		}
	}

	return roles
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
