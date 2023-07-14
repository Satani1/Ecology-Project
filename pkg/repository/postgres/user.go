package postgres

import (
	"ecogoly/pkg/models"
	"log"
)

func (pr *PostgresRepository) InsertUser(user models.User) error {
	stmt := `insert into ecodb.users (uid, name, password) values ($1, $2, $3)`

	_, err := pr.DB.Exec(stmt, user.UID, user.Name, user.Password)
	if err != nil {
		return err
	}
	return nil
}

func (pr *PostgresRepository) GetUserByName(name string) (*models.User, error) {
	stmt := `select uid, name, password from ecodb.users where name = $1`

	row := pr.DB.QueryRow(stmt, name)

	var user models.User

	if err := row.Scan(&user.UID, &user.Name, &user.Password); err != nil {
		log.Fatalln(err)
		return nil, err
	}

	if err := row.Err(); err != nil {
		log.Fatalln(err)
		return nil, err
	}
	return &user, nil
}
