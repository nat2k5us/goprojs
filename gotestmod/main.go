package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gorilla/mux"
)

// remove the go cache
// rm -rf ~/go/pkg/mod

func main() {
	ConnectToSQL()
	StartRestAPIServices()
}

// ConnectToSQL func
func ConnectToSQL() {
	db, err := sql.Open("mysql", "user:password@/dbname")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	// Open doesn't open a connection. Validate DSN data:
	err = db.Ping()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
}

// TestEndPoint func
func TestEndPoint(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "testing server - server is running !!!")
}

// StartRestAPIServices - starts the api servies
func StartRestAPIServices() {
	fmt.Println("Starting application")
	log.Print("Starting Go Server at http://localhost:8011")
	router := mux.NewRouter()

	router.HandleFunc("/test", TestEndPoint).Methods("GET")

	log.Fatal(http.ListenAndServe(":8011", router))
}
