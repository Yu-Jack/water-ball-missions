// Code generated by go-enum DO NOT EDIT.
// Version:
// Revision:
// Build Date:
// Built By:

package treasure

import (
	"errors"
	"fmt"
	"strings"
)

const (
	// ContentDokodemoDoor is a Content of type DokodemoDoor.
	ContentDokodemoDoor Content = "DokodemoDoor"
	// ContentPoison is a Content of type Poison.
	ContentPoison Content = "Poison"
	// ContentKingsRock is a Content of type KingsRock.
	ContentKingsRock Content = "KingsRock"
	// ContentSuperStar is a Content of type SuperStar.
	ContentSuperStar Content = "SuperStar"
	// ContentHealingPotion is a Content of type HealingPotion.
	ContentHealingPotion Content = "HealingPotion"
	// ContentAcceleratingPotion is a Content of type AcceleratingPotion.
	ContentAcceleratingPotion Content = "AcceleratingPotion"
	// ContentDevilFruit is a Content of type DevilFruit.
	ContentDevilFruit Content = "DevilFruit"
)

var ErrInvalidContent = errors.New("not a valid Content")

// String implements the Stringer interface.
func (x Content) String() string {
	return string(x)
}

// String implements the Stringer interface.
func (x Content) IsValid() bool {
	_, err := ParseContent(string(x))
	return err == nil
}

var _ContentValue = map[string]Content{
	"DokodemoDoor":       ContentDokodemoDoor,
	"dokodemodoor":       ContentDokodemoDoor,
	"Poison":             ContentPoison,
	"poison":             ContentPoison,
	"KingsRock":          ContentKingsRock,
	"kingsrock":          ContentKingsRock,
	"SuperStar":          ContentSuperStar,
	"superstar":          ContentSuperStar,
	"HealingPotion":      ContentHealingPotion,
	"healingpotion":      ContentHealingPotion,
	"AcceleratingPotion": ContentAcceleratingPotion,
	"acceleratingpotion": ContentAcceleratingPotion,
	"DevilFruit":         ContentDevilFruit,
	"devilfruit":         ContentDevilFruit,
}

// ParseContent attempts to convert a string to a Content.
func ParseContent(name string) (Content, error) {
	if x, ok := _ContentValue[name]; ok {
		return x, nil
	}
	// Case insensitive parse, do a separate lookup to prevent unnecessary cost of lowercasing a string if we don't need to.
	if x, ok := _ContentValue[strings.ToLower(name)]; ok {
		return x, nil
	}
	return Content(""), fmt.Errorf("%s is %w", name, ErrInvalidContent)
}