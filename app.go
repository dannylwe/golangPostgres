package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Initialize(user, password, dbname, host, port string) {
	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", host, port, user, password, dbname)

	var err error
	a.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	
	//Ping method actually calls db and makes 100% sure params are correct
	err = a.DB.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully Connected to remote Postgres")
	a.Router = mux.NewRouter()

	// defer a.DB.Close()
}

func (a *App) Run(addr string) {}
