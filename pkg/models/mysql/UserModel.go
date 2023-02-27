package mysql

import (
	"database/sql"
	"ecogoly/pkg/models"
)

type UserModel struct {
	DB *sql.DB
}

func (m *UserModel) Insert(firstName, lastName string) (int, error) {
	return 0, nil
}

func (m *UserModel) Get(id int) (*models.User, error) {
	return nil, nil
}

func (m *UserModel) Put(id int) error {
	return nil
}
