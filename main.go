package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/juzhen/postgresql-test/mysql"
	"github.com/juzhen/postgresql-test/pq"
)

var (
	user, password, dbname string
)

func main() {
	fmt.Println(os.Args)
	if len(os.Args) != 4 {
		fmt.Println("Usage: postgresql-test $USER $PASSWD $DBNAME")
		fmt.Println("Default args will be applied")

		user = "lance"
		password = "lancepwd"
		dbname = "testdb"
	} else {
		user = os.Args[2]
		password = os.Args[3]
		dbname = os.Args[4]
	}

	pq.Init("postgresql-test", user, password, dbname)
	err := pq.Open()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("pq status: ", pq.Status())

	mysql.Init("mysql-test", user, password, dbname)
	err = mysql.Open()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("mysql status: ", mysql.Status())

	// checkTable()
	// db.Query("CREATE TABLE hello  (kind varchar(10) )")

	mux := http.NewServeMux()
	mux.HandleFunc("/pq", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("got request")
		if err = pq.Ping(); err != nil {
			fmt.Println(err)
			fmt.Fprintf(w, fmt.Sprintf("Failed to ping db: %s", err.Error()))
		} else {
			fmt.Fprintf(w, "DB status OK")
		}
	})

	mux.HandleFunc("/mysql", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("got request")
		if err = mysql.Ping(); err != nil {
			fmt.Println(err)
			fmt.Fprintf(w, fmt.Sprintf("Failed to ping db: %s", err.Error()))
		} else {
			fmt.Fprintf(w, "DB status OK")
		}
	})

	mux.HandleFunc("/github", func(w http.ResponseWriter, r *http.Request) {
		resp, err := http.Get("http://apis.github.com")
		if err != nil {
			fmt.Println("Failed to get http://apis.github.com, e: ", err.Error())
			fmt.Fprintf(w, err.Error())
		} else if bs, e := ioutil.ReadAll(resp.Body); e != nil {
			fmt.Println("Failed to read resp.Body, e: ", e.Error())
			fmt.Fprintf(w, e.Error())
		} else {
			fmt.Print("Got http://apis.github.com!", "\n", string(bs))
			fmt.Fprintf(w, string(bs))
		}
	})
	mux.HandleFunc("/baidu", func(w http.ResponseWriter, r *http.Request) {
		resp, err := http.Get("http://www.baidu.com")
		if err != nil {
			fmt.Println("Failed to get http://www.baidu.com, e: ", err.Error())
			fmt.Fprintf(w, err.Error())
		} else if bs, e := ioutil.ReadAll(resp.Body); e != nil {
			fmt.Println("Failed to read resp.Body, e: ", e.Error())
			fmt.Fprintf(w, e.Error())
		} else {
			fmt.Print("Got http://www.baidu.com!", "\n", string(bs))
			fmt.Fprintf(w, string(bs))
		}
	})

	http.ListenAndServe(":8080", mux)
}
