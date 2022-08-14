package dbmodels

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type JsonData struct {
	Id    int    `json: "id"`
	Date  string `json: "date"`
	Input string `json: "Input"`
	Time  string `json: "time"`
}

var (
	destinationPath, _ = os.UserHomeDir()
)

func CreateUser(user *User) {
	db.Create(&user)
}

func GetUser(name string) User {
	var userStruct User
	db.Where("name = ?", name).Find(&userStruct)
	return userStruct
}

func CheckUser(name string, user *User) bool {
	if user == nil || name == "" {
		return false
	}

	db.Debug().Raw("select name, password, role from users where name = ?", name).Scan(&user)

	return true
}

func GetAllData() []User {
	var user []User
	db.Raw("select name, role from users").Scan(&user)
	return user
}

func DeleteItemAt(name string) {
	db.Debug().Unscoped().Where("name=?", name).Delete(&User{})
}

func InsertInputField(text, time string) InputField {
	inputField := InputField{
		Input:     text,
		TimeStamp: time,
	}

	return inputField
}

func InsertTextField(date string, inputField *InputField) {
	textField := TextField{
		CalendarDate: date,
		InputField:   inputField,
	}

	db.Debug().Create(&textField)
}

func GetTextFieldData(date string) []TextField {
	var textField []TextField
	db.Debug().Unscoped().Raw("select calendar_date, input, time_stamp from text_fields where calendar_date=?", date).Scan(&textField)
	return textField
}

//TODO Currently using map, problem with that it saves old data which we dont want. Just need to add new Data regardless of existing data.

func SaveAsJson(id int, date, input, time string) {
	m := []JsonData{}
	obj := JsonData{Id: id, Date: date, Input: input, Time: time}
	m = append(m, obj)
	data, err := json.MarshalIndent(m, " ", " ")
	if err != nil {
		fmt.Println("Something went wrong saving json")
	}

	if _, err := ioutil.ReadDir(destinationPath + "/Documents/calendarOutput/"); err != nil {
		fmt.Println("Creating Directory...")
		os.Mkdir(destinationPath+"/Documents/calendarOutput/", 0755)
	}

	if _, err := ioutil.ReadFile(destinationPath + "/Documents/calendarOutput/" + date + ".json"); err != nil {
		fmt.Println("Creating File...")
		ioutil.WriteFile(destinationPath+"/Documents/calendarOutput/"+date+".json", data, 0645)

	} else {
		file, err := os.OpenFile(destinationPath+"/Documents/calendarOutput/"+date+".json", os.O_RDWR|os.O_CREATE, 0644)
		if err != nil {
			fmt.Println(err)
		}

		local := []JsonData{}
		output, _ := ReadFromJson(date)
		local = append(local, output[:]...)
		local = append(local, obj)
		// fmt.Println(local)
		appendData, _ := json.MarshalIndent(local, " ", " ")

		// fmt.Println("obj: ", obj)
		if _, err := file.Write(appendData); err != nil {
			fmt.Println(err)
		}

		defer file.Close()
	}

}

func ReadFromJson(date string) ([]JsonData, error) {
	file, err := ioutil.ReadFile(destinationPath + "/Documents/calendarOutput/" + date + ".json")
	if err != nil {
		return nil, err
	}

	output := []JsonData{}

	_ = json.Unmarshal([]byte(file), &output)

	// for _, v := range output {
	// 	fmt.Println(v)
	// }
	return output, nil
}
