package domain

type RelationCurse struct {
	beCursed Role
	curse    Role
}

func NewRelationCurse(beCursed, curse Role) RelationObserver {
	return &RelationCurse{
		beCursed: beCursed,
		curse:    curse,
	}
}

func (rs *RelationCurse) GetName() string {
	return rs.curse.GetNameWithTroop()
}

func (rs *RelationCurse) Action() {
	if rs.curse.GetHp() > 0 {
		rs.curse.PlusHp(rs.beCursed.GetMp())
	}
}
