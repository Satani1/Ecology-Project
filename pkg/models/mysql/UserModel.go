package mysql

import (
	"database/sql"
	"ecogoly/pkg/models"
	"errors"
)

type UserModel struct {
	DB *sql.DB
}

func (m *UserModel) Insert(firstName, lastName, email string) (int, error) {
	stmt := `insert into ecologydb.users (firstName, lastName, email) values (?, ?, ?)`

	result, err := m.DB.Exec(stmt, firstName, lastName, email)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (m *UserModel) Get(id int) (*models.User, error) {
	stmt := `select user_id, firstName, lastName, email from ecologydb.users where user_id = ?`

	row := m.DB.QueryRow(stmt, id)

	s := &models.User{}

	err := row.Scan(&s.ID, &s.FirstName, &s.LastName, &s.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}

	return s, nil
}
