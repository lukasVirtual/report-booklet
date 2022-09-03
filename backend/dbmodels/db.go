package dbmodels

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	db  *gorm.DB
	err error
)

type User struct {
	gorm.Model
	ID        uint64 `json: "id"`
	Parent_id uint64 `json: "parent_id"`
	Name      string `json: "name"`
	Password  string `json: "password"`
	Role      string `json: "role"`
}

type InputField struct {
	Input     string `json: "input"`
	TimeStamp string `json: "timeStamp"`
}
type TextField struct {
	gorm.Model
	CalendarDate string `json: "calendarDate"`
	*InputField
}

func InitDB() {
	fmt.Println("INIT DB")
	// dsnWindows := "root:1234@tcp(127.0.0.1:3306)/test"
	dsnLinux := "sqluser:password@tcp(localhost:3306)/test"
	db, err = gorm.Open("mysql", dsnLinux)

	if err != nil {
		panic("Error Occured when tried to open DB")
	} else {
		log.Println("successfully connected")
	}

	if err = db.DB().Ping(); err != nil {
		panic("Error from DB")
	}

	db.AutoMigrate(&User{})
	// db.AutoMigrate(&TextField{})
	// db.AutoMigrate(&InputField{})
}
