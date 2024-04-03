package create

import (
	"database/sql"
	"fmt"
	"html/template"
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

		tpl := template.Must(template.ParseFiles("templates/insert.html"))
		tpl.Execute(w, nil)

		if r.Method == http.MethodPost {
			username := r.FormValue("username")
			email := r.FormValue("email")

			sqlStatement := `
	INSERT INTO users (username, email)
	VALUES ($1, $2)
`
			_, err := db.Exec(sqlStatement, username, email)
			if err != nil {
				panic(err)
			}
			fmt.Fprintf(w, "Inserted successfully")
		}

	}

}

func GetById(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		tpl := template.Must(template.ParseFiles("templates/GetById.html"))
		tpl.Execute(w, nil)

		if r.Method == http.MethodPost {
			id := r.FormValue("ID")

			sqlStatement := `
	SELECT username, email FROM users WHERE id = $1
	`
			var username, email string
			db.QueryRow(sqlStatement, id).Scan(&username, &email)
			fmt.Fprintf(w, "ID: %s, Username: %s, Email: %s\n", id, username, email)
		}
	}
}

func Delete(db *sql.DB, id int) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		tpl := template.Must(template.ParseFiles("templates/Delete.html"))
		tpl.Execute(w, nil)

		if r.Method == http.MethodPost {
			id := r.FormValue("ID")
			sqlStatement := `
		DELETE FROM users WHERE id=$1
		`

			db.Exec(sqlStatement, id)
			fmt.Println("Row deleted successfully")
		}
	}
}

func Update(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		tpl := template.Must(template.ParseFiles("templates/update.html"))
		tpl.Execute(w, nil)
		if r.Method == http.MethodPost {
			id := r.FormValue("ID")
			newEmail := r.FormValue("email")
			sqlStatement := `
		UPDATE users SET  email=$2 WHERE id=$1
		`

			db.Exec(sqlStatement, id, newEmail)
			fmt.Println("Row updated successfully")
		}
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
