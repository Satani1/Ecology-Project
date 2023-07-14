package main

import (
	"context"
	"ecogoly/internal"
	"ecogoly/pkg/repository/postgres"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

const addr string = "localhost:9000"

func main() {
	//read config variables
	cfg := LoadEnvVariables()

	//logs
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	//open db
	dbURL := fmt.Sprintf("postgres://%s:%s@/%s?sslmode=disable", cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresDB)

	infoLog.Println("Connecting to postgresDB")
	appDB, err := postgres.NewPostgres(dbURL)
	if err != nil {
		log.Fatalln("Cant connect to postgres:", err)
	}
	defer appDB.Close()
	infoLog.Println("Successful connect to postgresDB")

	//app struct
	App := &internal.Applicaton{
		ErrogLog:   errorLog,
		InfoLog:    infoLog,
		Repository: appDB,
		Secret:     cfg.Secret,
	}
	//Server config and router
	srv := &http.Server{
		Addr:     cfg.ServerAddr,
		ErrorLog: errorLog,
		Handler:  App.Routes(),
	}

	//running http server
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalln("Cant start the server:", err)
		}
		log.Println("Run the server", srv.Addr)
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalln("Cant shutdown the server", err)
	}

	log.Println("Shutdown the server.")
}
