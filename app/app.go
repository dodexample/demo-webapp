package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
  "net/http"
)

func main() {
	host := os.Getenv("PSQL_HOST")
	password := os.Getenv("PSQL_PASSWORD")
	dbname := os.Getenv("PSQL_DB")
  port := os.Getenv("PSQL_PORT")
	user := os.Getenv("PSQL_USER")
	psqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s "+
			"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		res, err := db.Query("select concat('hello from postgres backend pid ', pg_backend_pid());")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "DB Query failed; %s\n", err.Error())
			return
		}
    defer res.Close()
		for res.Next() {
			var msg string
			if err := res.Scan(&msg); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, "DB Result Scan failed; %s\n", err.Error())
				return
			}
			fmt.Fprintf(w, "%s\n", msg)
		}
	})
	panic(http.ListenAndServe(":8080", nil))
}
