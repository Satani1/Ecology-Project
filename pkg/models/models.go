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

type Marker struct {
	ID          int
	Name        string
	Description string
	Address     string
	Latitude    float64
	Longitude   float64
	Type        int
	Status      int
	DateCreated time.Time
}
