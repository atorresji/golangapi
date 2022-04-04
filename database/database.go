package database

import (
	"database/sql"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
)

const connectionString = "Data Source=LAPTOP-N81T3EAE;Initial Catalog=GoLangAPI;Integrated Security=True;"

var db *sql.DB

func OpenConnection() {
	connection, err := sql.Open("mssql", connectionString)
	if err != nil {
		fmt.Println("Is not possible to connect to database")
		panic(err)
	}
	db = connection
}

func CloseConnection() {
	db.Close()
}

func QueryStatement(query string) *sql.Rows {

	OpenConnection()

	rows, _ := db.Query(query)

	CloseConnection()

	return rows
}

func NonQueryStatement(query string) {
	OpenConnection()
	db.Exec(query)
	CloseConnection()
}

func Insert2(query string) {
	OpenConnection()
	db.Exec(query)
	CloseConnection()
}

func Insert(query string) int64 {

	rows, executeErr := db.Exec(query)
	defer db.Close()

	if executeErr != nil {
		fmt.Println("Something went wrong when executing statement")
	}

	id, insertErr := rows.RowsAffected()
	if insertErr != nil {
		fmt.Println("Unable to execute Non Query")
	}

	return id
}

func Ping() {
	if err := db.Ping(); err != nil {
		panic(err)
	}
	//fmt.Println("All good")
}
