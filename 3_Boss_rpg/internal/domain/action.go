//go:generate go-enum -f=$GOFILE --nocase

package domain

type ActionStrategy interface {
	S1(skillsIDs []int) int
	S2(availableIDs []int, limit int, list string) []int
}

// StrategyName
/*
ENUM(
Hero = "英雄"
AI = "AI"
)
*/
type StrategyName string
