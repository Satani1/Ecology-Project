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
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Address     string  `json:"address"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"Longitude"`
	Type        int
	Status      int
	DateCreated time.Time
}
type Marker2 struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
