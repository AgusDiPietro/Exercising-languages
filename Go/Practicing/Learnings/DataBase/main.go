package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func createTable(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS usuarios (
		id INT AUTO_INCREMENT,
		nombre VARCHAR(50),
		email VARCHAR(50),
		PRIMARY KEY (id)
	
	);`
	_, err := db.Exec(query)
	return err
}

func main() {
	// DSN (Data Source Name) para conectarse a la base de datos MySQL
	dsn := "usuario:contraseña@tcp(127.0.0.1:3306)/nombre_basedatos"

	db, err := sql.Open("mysql", dsn)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	//Creation of new table
	err = createTable(db)
	if err != nil {
		panic(err)
		return
	}
	fmt.Println("Table created")

}
