package dbmodels

type User struct {
	ID       uint64 `json: "id"`
	Name     string `json: "name"`
	Password string `json: "password"`
}

func CreateUser(user *User) error {
	statement := `insert into users(name, password) values(?, ?)`

	_, err := db.Exec(statement, user.Name, user.Password)

	return err

}

func GetUser(id string) (User, error) {
	var user User

	statement := `select * from users where id = ?`

	rows, err := db.Query(statement, id)

	if err != nil {
		return User{}, err
	}

	for rows.Next() {
		err = rows.Scan(&user.ID, &user.Name, &user.Password)
	}

	return user, nil
}
