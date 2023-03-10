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
	stmt := `insert into ecologydb.markers (name, description, address) values (?, ?, ?)`

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

func (m *MarkerModel) Get(id int) (*models.Marker2, error) {
	stmt := `select name, description, address from ecologydb.markers where mark_id = ?`

	row := m.DB.QueryRow(stmt, id)

	s := &models.Marker2{}

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
	rows, err := m.DB.Query("SELECT mark_id, name, description, address, status FROM ecologydb.markers")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	markers := []models.Marker2{}
	for rows.Next() {
		var marker models.Marker2
		err := rows.Scan(&marker.ID, &marker.Name, &marker.Description, &marker.Address, &marker.Status)
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

func (m *MarkerModel) UpdateMarkerToWork(id int) error {
	stmt := `update ecologyDB.markers set status = "В работе" where mark_id = ? and status = "Новая";`

	_, err := m.DB.Exec(stmt, id)
	if err != nil {
		return err
	}

	return nil
}

func (m *MarkerModel) Delete(id int) error {
	stmt := `delete from ecologydb.markers where mark_id = ?`

	_, err := m.DB.Exec(stmt, id)
	if err != nil {
		return err
	}

	return nil
}

func (m *MarkerModel) GetPhotoPath(id int) (string, error) {
	stmt := `select pathToPhoto from ecologydb.markers where mark_id = ?`

	row := m.DB.QueryRow(stmt, id)

	s := &models.Marker2{}

	err := row.Scan(&s.PathToPhoto)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", models.ErrNoRecord
		} else {
			return "", err
		}
	}
	return s.PathToPhoto, nil
}

func (m *MarkerModel) PutPhotoPath(path string, id int) (error) {
	stmt := `update ecologydb.markers set pathToPhoto = ? where mark_id = ?`

	_, err := m.DB.Exec(stmt, path, id)
	if err != nil {
		return err
	}
	return nil
}
