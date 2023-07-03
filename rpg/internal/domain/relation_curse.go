package domain

type RelationCurse struct {
	beCursed Role
	curse    Role
}

func NewRelationCurse(beCursed, curse Role) RelationCurse {
	return RelationCurse{
		beCursed: beCursed,
		curse:    curse,
	}
}

func (rs *RelationCurse) RecoverCurseHp() {
	if rs.curse.GetHp() > 0 {
		rs.curse.PlusHp(rs.beCursed.GetMp())
	}
}
