package database

import (
	"log"
	"sync"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rcmendes/crud-example-go/pkg/shared/logs"
)

var once sync.Once

var DB *sqlx.DB

func CreateTables() {
	InitDB()

	schema := `CREATE TABLE services (
		id text PRIMARY KEY, 
		name text,
		description text NULL,
		created_at TIMESTAMP,
		updated_at TIMESTAMP
		);`

	// execute a query on the server
	_, err := DB.Exec(schema)
	if err != nil {
		log.Fatal(err)
	}

}

func InitDB() {
	once.Do(func() {
		db, err := sqlx.Open("sqlite3", ":memory:")
		if err != nil {
			log.Fatal(err)
		}

		if err := db.Ping(); err != nil {
			logs.Error(err)
		}

		DB = db
	})
}
