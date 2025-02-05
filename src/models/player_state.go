package models

import "time"

type PlayerState struct {
	Name       string
	Value      string
	LastActive time.Time
}
