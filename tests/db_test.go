package tests

import (
	"database/sql"
	"fmt"
	"log"
	"testing"

	_ "github.com/lib/pq"
)

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func TestDB(t *testing.T) {
	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", "postgres", "postgres", "rc4laundry_test")
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	t.Run("schema exists", func(t *testing.T) {
		tx, err := db.Begin()
		must(err)
		rows, err := tx.Query("SELECT schema_name FROM information_schema.schemata WHERE schema_name = $1", "rc4laundry")
		must(err)
		defer rows.Close()
		if !rows.Next() {
			t.Error("Schema does not exist.")
		}
		err = tx.Rollback()
		must(err)
	})

	t.Run("machine_in_use table exists", func(t *testing.T) {
		tx, err := db.Begin()
		must(err)
		rows, err := db.Query("select (floor, position, type, started_at) from rc4laundry.machine_in_use where false")
		if err != nil {
			t.Error("Cannot query machine_in_use table.")
			log.Fatal(err)
		}
		if rows.Next() {
			t.Fatal("Got rows when there should have been none.")
		}
		err = tx.Rollback()
		must(err)
	})

	t.Run("should throw error if insert duplicate machines", func(t *testing.T) {
		tx, err := db.Begin()
		must(err)
		stmt, err := tx.Prepare("insert into rc4laundry.machine_in_use(floor, position, type) values($1, $2, $3)")
		must(err)
		defer stmt.Close()
		_, err = stmt.Exec(14, 0, "washer")
		must(err)
		_, err = stmt.Exec(14, 0, "dryer")
		if err == nil {
			t.Error("Expected an error but got nil.")
		}
		err = tx.Rollback()
		must(err)
	})
}
