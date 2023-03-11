package mysql

import (
	"database/sql"
	"ecogoly/pkg/models"
	"errors"
)

type MarkerModel struct {
	DB *sql.DB
}

func (m *MarkerModel) Insert(name, desc, addr, pathTo string, userId int) (int, error) {
	stmt := `insert into ecologydb.markers (name,description, address, pathToPhoto, fromUserID) values (?, ?, ?, ?, ?)`

	result, err := m.DB.Exec(stmt, name, desc, addr, pathTo, userId)
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
	rows, err := m.DB.Query("SELECT mark_id, name, description, address, status, pathToPhoto FROM ecologydb.markers where status <> 'На проверке'")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	markers := []models.Marker2{}
	for rows.Next() {
		var marker models.Marker2
		err := rows.Scan(&marker.ID, &marker.Name, &marker.Description, &marker.Address, &marker.Status, &marker.PathToPhoto)
		if err != nil {
			return nil, err
		}
		markers = append(markers, marker)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return &markers, nil
}

func (m *MarkerModel) UpdateMarkerToWork(id, userID int) error {
	stmt := `update ecologyDB.markers set status = "В работе", userCleanID = ? where mark_id = ? and status = "Новая";`

	_, err := m.DB.Exec(stmt, userID, id)
	if err != nil {
		return err
	}

	return nil
}

func (m *MarkerModel) UpdateMarkerToCheck(id int) error {
	stmt := `update ecologyDB.markers set status = "На проверке" where mark_id = ? and (status = "Новая" or status = "В работе")`

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

func (m *MarkerModel) CountUserMarks(id int) (*models.Rating, error) {
	//queries
	stmt := `select count(*) from ecologydb.markers where fromUserId = ? and status = 'Новая'`
	stmt2 := `select count(*) from ecologydb.markers where fromUserId = ? and status = 'В работе'`
	stmt3 := `select count(*) from ecologydb.markers where fromUserId = ? and status = 'На проверке'`
	//exes
	row := m.DB.QueryRow(stmt, id)
	row2 := m.DB.QueryRow(stmt2, id)
	row3 := m.DB.QueryRow(stmt3, id)

	s := &models.Rating{}

	//errors handling
	err := row.Scan(&s.New)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}

	err = row2.Scan(&s.InWork)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}

	err = row3.Scan(&s.Done)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}

	return s, nil
}
