package pq

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var (
	Host, User, Password, DBName string
	db                           *sql.DB
)

func Init(host, user, password, dbname string) {
	Host = host
	User = user
	Password = password
	DBName = dbname

}
func Open() error {
	connStr := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", User, Password, Host, DBName)
	fmt.Println(connStr)
	var err error
	db, err = sql.Open("postgres", connStr)
	return err
}

func Status() interface{} {
	return db.Stats()
}

func Ping() error {
	_, err := db.Query("SELECT * FROM information_schema.tables")
	// _, err := db.Query("SELECT * FROM hello")
	// _, err := db.Query("CREATE TABLE hello  (kind varchar(10) )")
	return err
}
