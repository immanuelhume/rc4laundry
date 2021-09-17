package tests

import (
	"database/sql"
	"fmt"
	"log"
	"testing"

	_ "github.com/lib/pq"
)

func must(t testing.TB, err error, tx *sql.Tx) {
	t.Helper()
	if err != nil {
		tx.Rollback()
		t.Fatal(err)
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
		must(t, err, tx)
		rows, err := tx.Query("SELECT schema_name FROM information_schema.schemata WHERE schema_name = $1", "rc4laundry")
		must(t, err, tx)
		defer rows.Close()
		if !rows.Next() {
			t.Error("Schema does not exist.")
		}
		err = tx.Rollback()
		must(t, err, tx)
	})

	t.Run("machine table exists", func(t *testing.T) {
		tx, err := db.Begin()
		must(t, err, tx)
		rows, err := db.Query("select (floor, position, type, last_started_at, is_in_use, approx_duration) from rc4laundry.machine where false")
		if err != nil {
			t.Fatal("Cannot query machine table.")
		}
		defer rows.Close()
		if rows.Next() {
			t.Fatal("Got rows when there should have been none.")
		}
		err = tx.Rollback()
		must(t, err, tx)
	})

	t.Run("should throw error if insert duplicate machines", func(t *testing.T) {
		tx, err := db.Begin()
		must(t, err, tx)
		stmt, err := tx.Prepare("insert into rc4laundry.machine(floor, position, type) values($1, $2, $3)")
		must(t, err, tx)
		defer stmt.Close()
		_, err = stmt.Exec(14, 0, "washer")
		must(t, err, tx)
		_, err = stmt.Exec(14, 0, "dryer")
		if err == nil {
			t.Error("Expected an error but got nil.")
		}
		err = tx.Rollback()
		must(t, err, tx)
	})

	t.Run("should insert the right default value for approx_duration if not given", func(t *testing.T) {
		cases := []struct {
			floor                  int
			position               int
			machineType            string
			expectedApproxDuration int
		}{
			{14, 0, "washer", 1800},
			{14, 2, "dryer", 2400},
		}
		tx, err := db.Begin()
		must(t, err, tx)
		stmnt, err := tx.Prepare("insert into rc4laundry.machine(floor, position, type) values ($1, $2, $3) returning approx_duration")
		defer stmnt.Close()
		must(t, err, tx)
		for _, test := range cases {
			var approxDuration int
			err := stmnt.QueryRow(test.floor, test.position, test.machineType).Scan(&approxDuration)
			must(t, err, tx)
			if approxDuration != test.expectedApproxDuration {
				t.Errorf("Expected %v got %v", test.expectedApproxDuration, approxDuration)
			}
		}
		err = tx.Rollback()
		must(t, err, tx)
	})

	t.Run("should use value for approx_duration if provided", func(t *testing.T) {
		testDuration := 420
		var approxDuration int
		tx, err := db.Begin()
		must(t, err, tx)
		err = tx.QueryRow("insert into rc4laundry.machine(floor, position, type, approx_duration) values ($1, $2, $3, $4) returning approx_duration", 14, 0, "washer", testDuration).Scan(&approxDuration)
		must(t, err, tx)
		if approxDuration != testDuration {
			t.Errorf("Expected %v got %v", testDuration, approxDuration)
		}
		err = tx.Rollback()
		must(t, err, tx)
	})
}
