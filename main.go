package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

var (
	host, user, password, dbname string
	db                           *sql.DB
)

func main() {
	fmt.Println(os.Args)
	if len(os.Args) != 5 {
		fmt.Println("Usage: postgresql-test $HOST $USER $PASSWD $DBNAME")
		fmt.Println("Default args will be applied")

		host = "postgresql-test"
		user = "lance"
		password = "lancepwd"
		dbname = "testdb"
	} else {
		host = os.Args[1]
		user = os.Args[2]
		password = os.Args[3]
		dbname = os.Args[4]
	}

	connStr := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", user, password, host, dbname)
	fmt.Println(connStr)
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("connected", db.Stats())

	// checkTable()
	// db.Query("CREATE TABLE hello  (kind varchar(10) )")

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("got request")
		if err = pingDB(); err != nil {
			fmt.Println(err)
			fmt.Fprintf(w, fmt.Sprintf("Failed to ping db: %s", err.Error()))
		} else {
			fmt.Fprintf(w, "DB status OK")
		}
	})
	http.ListenAndServe(":8080", mux)
}

func pingDB() error {
	_, err := db.Query("SELECT * FROM information_schema.tables")
	// _, err := db.Query("SELECT * FROM hello")
	// _, err := db.Query("CREATE TABLE hello  (kind varchar(10) )")
	return err

}

func checkTable() {
	db.Query("CREATE TABLE hello  (kind varchar(10) )")
}
