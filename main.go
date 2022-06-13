package main

import (
	"fmt"
	"net/http"
)

func testFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "test Function\n")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello Function\n")
	})
	mux.Handle("/test", http.HandlerFunc(testFunc))
	http.ListenAndServe("127.0.0.1:80", mux)
}

//
//import (
//	"database/sql"
//	"fmt"
//	_ "github.com/go-sql-driver/mysql"
//	"time"
//)
//
//var databaseURL string = "root:!gktrlagh33@(127.0.0.1:3306)/go_user?parseTime=true"
//
//func main() {
//
//	db, err := sql.Open("mysql", databaseURL)
//	if err != nil {
//		fmt.Println(err)
//	}
//	username := "kimho"
//	password := "rlagh33"
//	createAt := time.Now()
//
//	result, err := db.Exec(`INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)`, username, password, createAt)
//	if err != nil {
//		fmt.Println(err)
//	}
//	userID, err := result.LastInsertId()
//	fmt.Println(userID)
//}
