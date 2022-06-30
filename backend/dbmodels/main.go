package dbmodels

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitDB() {
	fmt.Println("Database init")
	var err error
	db, err = sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/test")

	if err != nil {
		panic(err)
	}

	err2 := db.Ping()

	// user := &User{
	// 	Name:     "Peteer",
	// 	Password: "1234",
	// }

	// err = CreateUser(user)

	// _, err = db.Query("insert into users(name, password) values('gsdfsd', 1234) ")

	// if err != nil {
	// 	panic(err)
	// }
	if err2 != nil {
		panic(err2)
	}

	// defer db.Close()

	fmt.Println("db is connected successfully")
}
