package db

import (
	"database/sql"
	"fmt"
	"io/ioutil"

	_ "github.com/lib/pq" // postgres
)

// Database contains objects for database communication.
type Database struct {
	*sql.DB
}

// New create new database instance.
func New(config *Config) (*Database, error) {
	psqlInfo := fmt.Sprintf("host=%v port=%v user=%v "+
		"password=%v dbname=%v sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.DbName)

	db, _ := sql.Open("postgres", psqlInfo)
	err := db.Ping()
	if err != nil {
		return nil, fmt.Errorf("Error to connect to database")
	}

	return &Database{db}, nil
}

//RunFileQuery Run sql schema
func (db *Database) RunFileQuery(file string) {
	query, err := ioutil.ReadFile("./" + file)
	if err != nil {
		panic(err)
	}
	if _, err := db.Exec(string(query)); err != nil {
		panic(err)
	}
}
