package mysql

import (
	"database/sql"
	"ecogoly/pkg/models"
	"errors"
)

type MarkerModel struct {
	DB *sql.DB
}

func (m *MarkerModel) Insert(name, desc, addr string, status, type_ int) (int, error) {
	stmt := `insert into ecologydb.marks (name,description,addr,status, type,dateCreated ) values (?, ?, ?, ?, ?, utc_timestamp())`

	result, err := m.DB.Exec(stmt, name, desc, addr, status, type_)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (m *MarkerModel) Get(id int) (*models.Marker, error) {
	stmt := `select name,description,addr from ecologydb.marks where id = ?`

	row := m.DB.QueryRow(stmt, id)

	s := &models.Marker{}

	err := row.Scan(&s.ID, &s.Name, &s.Description, &s.Address)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}

	return s, nil
}

func (m *MarkerModel) Put(id int) error {
	return nil
}
