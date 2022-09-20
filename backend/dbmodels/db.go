package dbmodels

import (
	"fmt"
	"log"
	"runtime"

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
	// fmt.Println(runtime.GOOS == "windows")
	var dsn string
	switch runtime.GOOS {
	case "windows":
		dsn = "root:1234@tcp(127.0.0.1:3306)/test"
		break
	case "linux":
		dsn = "sqluser:password@tcp(localhost:3306)/test"
		break
	default:
		dsn = "sqluser:password@tcp(localhost:3306)/test"
		break
	}
	// dsnWindows := "root:1234@tcp(127.0.0.1:3306)/test"
	// dsnLinux := "sqluser:password@tcp(localhost:3306)/test"
	db, err = gorm.Open("mysql", dsn+"?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic("Error Occured when tried to open DB")
	} else {
		log.Println("successfully connected")
	}

	if err = db.DB().Ping(); err != nil {
		panic("Error from DB")
	}

	/*
		instead type varchar timestamp(6)
	*/
	db.AutoMigrate(&User{})
}
