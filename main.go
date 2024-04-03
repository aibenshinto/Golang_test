package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"main.go/create"
	"main.go/dbconnect"
)

func main() {
	db1 := dbconnect.Dbconnect()

	r := mux.NewRouter()

	r.HandleFunc("/VeiwAll", create.VeiwAll(db1)).Methods("GET")

	r.HandleFunc("/VInsertValues", create.InsertValues(db1))

	r.HandleFunc("/GetById", create.GetById(db1))

	r.HandleFunc("/Update", create.Update(db1))

	r.HandleFunc("/Delete", create.Delete(db1, 16))

	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
	defer db1.Close()
}
