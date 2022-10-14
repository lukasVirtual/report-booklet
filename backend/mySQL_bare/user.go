package dbmodels

// import "fmt"

// type User struct {
// 	ID       uint64 `json: "id"`
// 	Name     string `json: "name"`
// 	Password string `json: "password"`
// 	Role     string `json: "role"`
// }

// func CreateUser(user *User) error {
// 	statement := `insert into users(name, password, role) values(?, ?, ?)`

// 	_, err := db.Exec(statement, user.Name, user.Password, user.Role)

// 	return err

// }

// func GetUser(name string) (User, error) {
// 	var user User

// 	statement := `select name, password, role from users where name=? limit 1`
// 	rows, err := db.Query(statement, name)
// 	if err != nil {
// 		return User{}, err
// 	}

// 	for rows.Next() {
// 		err = rows.Scan(&user.Name, &user.Password, &user.Role)
// 		fmt.Println(err)
// 	}

// 	return user, nil
// }

// func CheckUser(name string, user *User) bool {
// 	statement := `select name, password from users where name=? limit 1`
// 	rows, err := db.Query(statement, name)
// 	if err != nil {
// 		return false
// 	}

// 	for rows.Next() {
// 		err = rows.Scan(&user.Name, &user.Password)
// 		if err != nil {
// 			return false
// 		}
// 	}

// 	return true
// }
