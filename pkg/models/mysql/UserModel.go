package mysql

import (
	"database/sql"
	"ecogoly/pkg/models"
	"errors"
)

type UserModel struct {
	DB *sql.DB
}

func (m *UserModel) Insert(firstName, lastName string) (int, error) {
	stmt := `insert into users (firstName, lastName, dateCreated) values (?, ?,utc_timestamp())`

	result, err := m.DB.Exec(stmt, firstName, lastName)
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
	stmt := `select id, firstName,lastName, email from users where id = ?`

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

func (m *UserModel) Put(id int) error {
	return nil
}
