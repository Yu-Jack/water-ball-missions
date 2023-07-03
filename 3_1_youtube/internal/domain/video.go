package domain

import (
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
