package main

import (
	"context"
	"database/sql"
	"ecogoly/internal"
	mysql2 "ecogoly/pkg/repository/mysql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

const addr string = "localhost:9000"

func main() {
	dsn := "root:YaPoc290302@/ecologydb?parseTime=true"
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

	//dsn = "postgres://postgres:12345678@postgres/postgres?sslmode=disable"
	//db, err := postgres.NewPostgres(dsn)
	//if err != nil {
	//	errorLog.Fatal(err)
	//}
	//defer db.Close()

	//app struct
	App := &internal.Applicaton{
		ErrogLog:  errorLog,
		InfoLog:   infoLog,
		UsersDB:   &mysql2.UserModel{DB: userDB},
		MarkersDB: &mysql2.MarkerModel{DB: markerDB},
	}

	//Server config and router
	srv := &http.Server{
		Addr:     addr,
		ErrorLog: errorLog,
		Handler:  App.Routes(),
	}

	//launch
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal("Cant start the server", err)
		}
		log.Println("Launch server on addr", addr)
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Cant shutdown the server", err)
	}

	log.Println("Shutdown...")
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
