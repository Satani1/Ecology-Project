package models

import (
	"errors"
)

var ErrNoRecord = errors.New("models: no matching entry was found")

type User struct {
	UID      string `json:"uid"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Marks    int    `json:"rating"`
}

type Marker struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Address     string `json:"address"`
	Status      string `json:"status"`
	PathToPhoto string `json:"pathToPhoto"`
	FromUserID  int    `json:"fromUserID"`
}

type Rating struct {
	New    int `json:"new"`
	InWork int `json:"inWork"`
	Done   int `json:"done"`
}
