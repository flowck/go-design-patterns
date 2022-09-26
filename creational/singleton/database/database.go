package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
	"strconv"
	"sync"
)

type Database struct {
	name               string
	port               int
	host               string
	maxIdleConnections int
	instance           *sql.DB
}

func (d *Database) GetInstance() *sql.DB {
	return d.instance
}

var db *Database
var once sync.Once

func GetDatabaseInstance() (*Database, error) {
	var err error
	var port int
	var instance *sql.DB
	var maxIdleConnections int

	if port, err = strconv.Atoi(os.Getenv("DB_PORT")); err != nil {
		return nil, err
	}

	if maxIdleConnections, err = strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONNECTIONS")); err != nil {
		return nil, err
	}

	if instance, err = sql.Open("postgres", "some-db-uri"); err != nil {
		return nil, err
	}

	// Run this func only once
	once.Do(func() {
		if db == nil {
			instance.SetMaxIdleConns(maxIdleConnections)

			db = &Database{
				name:               os.Getenv("DB_NAME"),
				port:               port,
				host:               os.Getenv("DB_HOST"),
				maxIdleConnections: maxIdleConnections,
				instance:           instance,
			}
			fmt.Println("Db instance has been created")
		}
	})

	return db, nil
}
