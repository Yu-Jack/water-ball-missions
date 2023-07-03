package domain

type RelationSummon struct {
	beSummoned Role
	summoner   Role
}

func NewRelationSummon(beSummoned, summoner Role) RelationSummon {
	return RelationSummon{
		beSummoned: beSummoned,
		summoner:   summoner,
	}
}

func (rs *RelationSummon) RecoverSummonerHp() {
	if rs.summoner.GetHp() > 0 {
		rs.summoner.PlusHp(30)
	}
}
