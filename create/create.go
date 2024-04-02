package create

import (
	"database/sql"
	"fmt"
	"net/http"
)

func CreateTable(db *sql.DB) error {
	// SQL command to create a table
	sqlStatement := `
        CREATE TABLE IF NOT EXISTS users (
            id SERIAL PRIMARY KEY,
            username VARCHAR(50),
            email VARCHAR(100)  
        )
    `
	// Executing the SQL command
	_, err := db.Exec(sqlStatement)
	if err != nil {
		return err
	}
	fmt.Println("Table created successfully")
	return nil
}

func InsertValues(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sqlStatement := `
	INSERT INTO users (username, email)
	VALUES ($1, $2)
`
		_, err := db.Exec(sqlStatement, "dfdfsdfafgsfs", "gfagraddfsf")
		if err != nil {
			panic(err)
		}
		fmt.Fprintf(w, "Inserted successfully")
	}

}

func GetById(db *sql.DB, id int) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sqlStatement := `
	SELECT username, email FROM users WHERE id = $1
	`
		var username, email string
		db.QueryRow(sqlStatement, id).Scan(&id, &username, &email)
		fmt.Printf("ID: %d, Username: %s, Email: %s\n", id, username, email)
	}
}

func Delete(db *sql.DB, id int) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sqlStatement := `
		DELETE FROM users WHERE id=$1
		`

		db.Exec(sqlStatement, id)
		fmt.Println("Row deleted successfully")

	}
}

func Update(db *sql.DB, id int, newEmail string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sqlStatement := `
		UPDATE users SET  email=$2 WHERE id=$1
		`

		db.Exec(sqlStatement, id, newEmail)
		fmt.Println("Row updated successfully")

	}
}

func VeiwAll(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sqlStatement := `
	SELECT * FROM users 
	`

		rows, _ := db.Query(sqlStatement)
		defer rows.Close()
		for rows.Next() {
			var id int
			var username, email string
			rows.Scan(&id, &username, &email)

			fmt.Fprintf(w, "ID: %d, Username: %s, Email: %s\n", id, username, email)
		}
	}
}
