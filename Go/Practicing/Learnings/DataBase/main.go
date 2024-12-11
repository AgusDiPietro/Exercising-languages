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

func insertUsuario(db *sql.DB, nombre, email string) (int64, error) {
	result, err := db.Exec("INSERT INTO usuarios (nombre, email) VALUES (?, ?)", nombre, email)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func updateUserEmail(db *sql.DB, id int64, newEmail string) error {
	_, err := db.Exec("UPDATE user SET email = ? WHERE id = ?", newEmail, id)
	return err
}

func deleteUser(db *sql.DB, id int64) error {
	_, err := db.Exec("DELETE FROM usuarios WHERE id = ?", id)
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

	id, err := insertUsuario(db, "Agus", "agus@gmail.com")
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	fmt.Println("User inserted with ID: ", id)
}
