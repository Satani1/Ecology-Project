package postgres

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type PostgresRepository struct {
	DB *sql.DB
}

func NewPostgres(dsn string) (*PostgresRepository, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresRepository{db}, nil
}

func (pr *PostgresRepository) Close() {
	pr.DB.Close()
}
