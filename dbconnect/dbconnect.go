package dbconnect

import (
	"database/sql"
	"fmt"

	"main.go/create"
)

func Dbconnect() *sql.DB {
	db, err := sql.Open("postgres", "postgres://vefceiis:pzTKYphme3DcRyDg3ertxItmrtc44HiQ@rain.db.elephantsql.com/vefceiis")
	if err != nil {
		panic(err)
	}
	
	// Check if the connection is successful
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to the database!")

	err = create.CreateTable(db)
	if err != nil {
		panic(err)
	}
	return db
}
