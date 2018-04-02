package mysql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
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
	connStr := fmt.Sprintf("%s:%s@tcp(%s)/%s", User, Password, Host, DBName)
	fmt.Println(connStr)
	var err error
	db, err = sql.Open("mysql", connStr)
	return err
}

func Status() interface{} {
	return db.Stats()
}

func Ping() error {
	_, err := db.Query("SHOW TABLES")
	// _, err := db.Query("SELECT * FROM hello")
	// _, err := db.Query("CREATE TABLE hello  (kind varchar(10) )")
	return err
}
