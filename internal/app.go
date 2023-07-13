package internal

import (
	"ecogoly/pkg/models/mysql"
	"log"
)

type Applicaton struct {
	ErrogLog  *log.Logger
	InfoLog   *log.Logger
	UsersDB   *mysql.UserModel
	MarkersDB *mysql.MarkerModel
}
