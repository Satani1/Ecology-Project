package mysql

import (
	"database/sql"
	"ecogoly/pkg/models"
	"errors"
	"log"
)

type MarkerModel struct {
	DB *sql.DB
}

func (m *MarkerModel) Insert(name, desc, addr string) (int, error) {
	stmt := `insert into ecologydb.markers (name, description, addr) values (?, ?, ?)`

	result, err := m.DB.Exec(stmt, name, desc, addr)
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
	stmt := `select name,description,addr from ecologydb.markers where mark_id = ?`

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

func (m *MarkerModel) GetAll() (*[]models.Marker2, error) {
	rows, err := m.DB.Query("SELECT mark_id, name, latitude, longitude FROM ecologydb.markers")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	markers := []models.Marker2{}
	for rows.Next() {
		var marker models.Marker2
		err := rows.Scan(&marker.ID, &marker.Name, &marker.Latitude, &marker.Longitude)
		if err != nil {
			log.Fatal(err)
		}
		markers = append(markers, marker)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return &markers, nil
}
