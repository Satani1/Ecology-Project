package models

import (
	"errors"
	"time"
)

var ErrNoRecord = errors.New("models: no matching entry was found")

type User struct {
	ID          int
	FirstName   string
	LastName    string
	DateCreated time.Time
}
