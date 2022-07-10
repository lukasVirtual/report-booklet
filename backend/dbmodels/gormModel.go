package dbmodels

import (
	"fmt"
	"log"
)

func CreateUser(user *User) {
	db.Create(&user)
}

//TODO Only Scan Name, Password and Role
func GetUser(name string) User {
	var userStruct User
	db.Where("name = ?", name).Find(&userStruct)
	return userStruct
}

//TODO Only Scan Name, Password and Role
func CheckUser(name string, user *User) bool {
	log.Println(name)
	err := db.Raw("select name, password, role from users where name = ?", name).Scan(&user)
	if err != nil {
		fmt.Println("error is from here", err)
	}
	if user.Name != "" && user.Password != "" {
		return true
	} else {
		return false
	}
}

func GetAllData() []User {
	var user []User
	db.Raw("select name, role from users").Scan(&user)
	return user
}
