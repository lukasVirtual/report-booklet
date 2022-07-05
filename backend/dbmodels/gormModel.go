package dbmodels

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

	db.Raw("select name, password, role from users where name = ?", name).Scan(&user)
	if user.Name != "" && user.Password != "" {
		return true
	} else {
		return false
	}
}
