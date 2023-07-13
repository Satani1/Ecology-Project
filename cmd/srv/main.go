package main

import (
	"database/sql"
	"ecogoly/internal"
	"ecogoly/pkg/models/mysql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"os"
)

const addr string = "localhost:9000"

func main() {
	dsn := "srv:ecoPass@/ecologydb?parseTime=true"
	//logs
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	userDB, err := openDB(dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer userDB.Close()

	markerDB, err := openDB(dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer markerDB.Close()
	//app struct
	App := &internal.Applicaton{
		ErrogLog:  errorLog,
		InfoLog:   infoLog,
		UsersDB:   &mysql.UserModel{DB: userDB},
		MarkersDB: &mysql.MarkerModel{DB: markerDB},
	}
	//Server config and router
	srv := &http.Server{
		Addr:     addr,
		ErrorLog: errorLog,
		Handler:  App.Routes(),
	}

	//launch
	infoLog.Printf("Launching server on %s", addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
