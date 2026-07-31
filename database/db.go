package database
import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)
var DB *sql.DB
func Connect() {
	connStr := "host= localhost port=5432 user=postgres password=password dbname=todo_api sslmode=disable"

	var err error
	DB, err = sql.Open("postgres", connStr)
		if err != nil {
		log.Fatal(err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("✅ Berhasil terkoneksi ke PostgreSQL")

}