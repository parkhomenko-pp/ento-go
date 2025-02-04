package models

import "time"

type UserState struct {
	State      string
	LastActive time.Time
}
