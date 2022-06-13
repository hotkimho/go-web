package main

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	_ "go-sql-driver/mysql"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/book/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		//
		fmt.Println(r.Cookies())
	})

	db, err := sql.Open("mysql", "root:!gktrlagh33@(127.0.0.1:3306)/go_user?parseTime=true")
	if err != nil {
		fmt.Println(err)
	}
	err = db.Ping()
	http.ListenAndServe(":80", router)
}
