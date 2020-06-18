package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//check if everything is installed

	db, err := sql.Open("mysql", "test:password@/")
	createDB(db)
	db.Close()
	db, err = sql.Open("mysql", "test:password@/c2server")

	InitStore(&dbStore{db: db})

	if err != nil {
		panic(err)
	}

	r := newRouter()

	http.ListenAndServe(":8080", r)

}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}
