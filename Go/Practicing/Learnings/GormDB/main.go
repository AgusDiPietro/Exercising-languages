package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID    uint
	Name  string
	Email string
}

func main() {
	dsn := "user:password@tcp(127.0.0.1:3306)/database_name"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to the db")
	}

	db.AutoMigrate(&User{})
	fmt.Println("Migration complete")

	db.Create(&User{Name: "Agus", Email: "aguys@gmail.com"})

	var user User
	db.First(&user, 1) //Search id
	fmt.Println(user)

	db.Model(&user).Update("Email", "agusnew@gmail.com")

	db.Delete(&user, 1)

	fmt.Println("Connected to Grom")
}
