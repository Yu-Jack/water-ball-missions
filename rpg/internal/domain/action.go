package domain

type ActionStrategy interface {
	S1(skillsIDs []int) int
	S2(availableIDs []int, limit int, list string) []int
}
