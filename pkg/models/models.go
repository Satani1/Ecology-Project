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
	Username    string
	DateCreated time.Time
	Rating      int
	Email       string
	Password    string
}
