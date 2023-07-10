package game

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"rpg/internal/domain"
	"rpg/internal/domain/action"
	"rpg/internal/domain/skill"
	"rpg/internal/domain/state"
)

func StartRPG() {
	rpg := domain.NewClientRPG()
	count := 0
	reader := bufio.NewReader(os.Stdin)

	for count < 2 {
		var input string
		var troopID int

		input, _ = reader.ReadString('\n')
		input = strings.TrimSuffix(input, "\n")

		re := regexp.MustCompile(`#軍隊-(\d+)-開始`)
		match := re.FindStringSubmatch(input)

		if len(match) > 1 {
			troopID, _ = strconv.Atoi(match[1])
		} else {
			fmt.Println("未找到數字，請重新輸入一遍")
			continue
		}

		count++

		t := domain.NewTroop(troopID, rpg)

		for {
			input, _ = reader.ReadString('\n')
			input = strings.TrimSuffix(input, "\n")

			if checkInputEnding(input, troopID) {
				break
			}

			infos := strings.Split(input, " ")

			name := infos[0]
			hp, _ := strconv.Atoi(infos[1])
			mp, _ := strconv.Atoi(infos[2])
			str, _ := strconv.Atoi(infos[3])

			skills := []domain.Skill{skill.NewBasicAttack()}

			if len(infos) > 4 {
				skills = append(skills, getSkills(infos[4:])...)
			}

			strategy := action.NewAiI()

			if name == domain.StrategyNameHero.String() {
				strategy = action.NewHero(reader)
			}

			t.AddRole(domain.NewRole(
				name, hp, mp, str, state.NewNormalState(),
				skills,
				strategy,
			))
		}

		rpg.AddTroop(t)
	}

	rpg.Start()
}

func getSkills(inputs []string) []domain.Skill {
	var skills []domain.Skill

	for i := 0; i < len(inputs); i++ {
		str := inputs[i]
		var s domain.Skill

		if str == domain.SkillNameFireBall.String() {
			s = skill.NewFireBall()
		}

		if str == domain.SkillNameWaterBall.String() {
			s = skill.NewWaterBall()
		}

		if str == domain.SkillNameSelfHealing.String() {
			s = skill.NewSelfHealing()
		}

		if str == domain.SkillNamePetrochemical.String() {
			s = skill.NewPetrochemical()
		}

		if str == domain.SkillNamePoison.String() {
			s = skill.NewPoison()
		}

		if str == domain.SkillNameSummon.String() {
			s = skill.NewSummon()
		}

		if str == domain.SkillNameSelfExplosion.String() {
			s = skill.NewSelfExplosion()
		}

		if str == domain.SkillNameCheerUp.String() {
			s = skill.NewCheerUp()
		}

		if str == domain.SkillNameCurse.String() {
			s = skill.NewCurse()
		}

		if str == domain.SkillNameOnePunch.String() {
			s = skill.NewOnePunch(
				skill.NewOnePunchHandler(
					skill.NewOnePunchI(), skill.NewOnePunchHandler(
						skill.NewOnePunchII(), skill.NewOnePunchHandler(
							skill.NewOnePunchIII(), skill.NewOnePunchHandler(
								skill.NewOnePunchIV(), nil,
							),
						),
					),
				),
			)
		}

		skills = append(skills, s)
	}

	return skills
}

func checkInputEnding(input string, targetID int) bool {
	re := regexp.MustCompile(`#軍隊-(\d+)-結束`)

	match := re.FindStringSubmatch(input)
	if len(match) > 1 {
		troopID, _ := strconv.Atoi(match[1])
		if troopID == targetID {
			return true
		}
	}

	return false
}
