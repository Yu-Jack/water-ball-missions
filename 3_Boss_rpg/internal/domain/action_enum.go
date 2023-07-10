// Code generated by go-enum DO NOT EDIT.
// Version:
// Revision:
// Build Date:
// Built By:

package domain

import (
	"errors"
	"fmt"
	"strings"
)

const (
	// StrategyNameHero is a StrategyName of type Hero.
	StrategyNameHero StrategyName = "英雄"
	// StrategyNameAI is a StrategyName of type AI.
	StrategyNameAI StrategyName = "AI"
)

var ErrInvalidStrategyName = errors.New("not a valid StrategyName")

// String implements the Stringer interface.
func (x StrategyName) String() string {
	return string(x)
}

// String implements the Stringer interface.
func (x StrategyName) IsValid() bool {
	_, err := ParseStrategyName(string(x))
	return err == nil
}

var _StrategyNameValue = map[string]StrategyName{
	"英雄": StrategyNameHero,
	"AI": StrategyNameAI,
	"ai": StrategyNameAI,
}

// ParseStrategyName attempts to convert a string to a StrategyName.
func ParseStrategyName(name string) (StrategyName, error) {
	if x, ok := _StrategyNameValue[name]; ok {
		return x, nil
	}
	// Case insensitive parse, do a separate lookup to prevent unnecessary cost of lowercasing a string if we don't need to.
	if x, ok := _StrategyNameValue[strings.ToLower(name)]; ok {
		return x, nil
	}
	return StrategyName(""), fmt.Errorf("%s is %w", name, ErrInvalidStrategyName)
}