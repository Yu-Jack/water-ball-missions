package domain

type RelationSummon struct {
	beSummoned Role
	summoner   Role
}

func NewRelationSummon(beSummoned, summoner Role) RelationObserver {
	return &RelationSummon{
		beSummoned: beSummoned,
		summoner:   summoner,
	}
}

func (rs *RelationSummon) Action() {
	if rs.summoner.GetHp() > 0 {
		rs.summoner.PlusHp(30)
	}
}

func (rs *RelationSummon) GetAttackerName() string {
	return rs.summoner.GetNameWithTroop()
}
