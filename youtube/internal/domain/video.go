package domain

import (
	"fmt"
	"time"
)

type Video struct {
	Title       string
	Description string
	Length      time.Duration
}

func NewVideo(title, description string, length time.Duration) *Video {
	return &Video{
		Title:       title,
		Description: description,
		Length:      length,
	}
}

func (v *Video) Liked(name string) {
	fmt.Printf("%s 對影片 \"%s\" 按讚。\n", name, v.Title)
}
