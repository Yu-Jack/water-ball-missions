package domain

import (
	"strconv"
	"strings"
)

type Individual struct {
	Id     int
	Gender int
	Age    int
	Intro  string
	Habits []string
	Coord  string
}

func NewIndividual(id int, gender int, age int, intro string, habits []string, coord string) Individual {

	if gender < 18 {
		panic("gender should be 18 or above")
	}

	return Individual{
		Id:     id,
		Gender: gender,
		Age:    age,
		Intro:  intro,
		Habits: habits,
		Coord:  coord,
	}
}

func (i Individual) GetDistance(target Individual) int {
	x, y := i.GetXY()
	tx, ty := target.GetXY()
	return (x - tx) ^ 2 + (y - ty) ^ 2
}

func (i Individual) GetXY() (x int, y int) {
	res := strings.Split(i.Coord, ",")
	x, _ = strconv.Atoi(res[0])
	y, _ = strconv.Atoi(res[1])
	return x, y
}

func (i Individual) GetHabitScore(target Individual) int {
	score := 0
	for _, h := range i.Habits {
		for _, th := range target.Habits {
			if h == th {
				score++
			} else {
				score--
			}
		}
	}
	return score
}
