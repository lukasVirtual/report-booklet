package dbmodels

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type JsonData struct {
	Id    int    `json: "id"`
	Date  string `json: "date"`
	Input string `json: "Input"`
	Time  string `json: "time"`
}
type GlobalWriteJsonStructure struct {
	Path  string      `json: "path"`
	Date  string      `json: "date"`
	Input interface{} `json: "input"`
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

func SaveAsJson(id int, date, input, time string) {
	m := []JsonData{}
	obj := JsonData{Id: id, Date: date, Input: input, Time: time}
	m = append(m, obj)
	data, err := json.MarshalIndent(m, " ", " ")
	if err != nil {
		fmt.Println("Something went wrong saving json")
	}

	if _, err := ioutil.ReadDir(destinationPath + "/Documents/TextFieldOutput/"); err != nil {
		fmt.Println("Creating Directory...")
		os.Mkdir(destinationPath+"/Documents/TextFieldOutput/", 0755)
	}

	if _, err := ioutil.ReadFile(destinationPath + "/Documents/TextFieldOutput/" + date + ".json"); err != nil {
		fmt.Println("Creating File...")
		ioutil.WriteFile(destinationPath+"/Documents/TextFieldOutput/"+date+".json", data, 0645)

	} else {
		local := []JsonData{}
		output, _ := ReadFromJson(date)
		allowed := true
		for i, v := range output {
			if v.Id == obj.Id {
				output[i].Input = obj.Input
				output[i].Time = obj.Time
				allowed = false
				break
			}
		}
		local = append(local, output[:]...)
		if allowed {
			local = append(local, obj)
		}
		// fmt.Println(local)
		appendData, _ := json.MarshalIndent(local, " ", " ")

		ioutil.WriteFile(destinationPath+"/Documents/TextFieldOutput/"+date+".json", appendData, 0755)
	}

}

func ReadFromJson(date string) ([]JsonData, error) {
	file, err := ioutil.ReadFile(destinationPath + "/Documents/TextFieldOutput/" + date + ".json")
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

func RemoveFromJson(date string, index int) {
	output, err := ReadFromJson(date)
	if err != nil {
		fmt.Println(err)
	}
	var local []JsonData
	if len(output) >= index {
		output = append(output[:index], output[index+1:]...)
	}
	local = append(local, output[:]...)
	fmt.Println(local)
	for i, _ := range local {
		local[i].Id = i
	}
	rawData, _ := json.MarshalIndent(local, " ", " ")

	ioutil.WriteFile(destinationPath+"/Documents/TextFieldOutput/"+date+".json", rawData, 0755)

}

func GlobalWriteJson(path, date string, input ...interface{}) {
	obj := GlobalWriteJsonStructure{Input: input}
	rawData, err := json.MarshalIndent(obj, " ", " ")
	if err != nil {
		fmt.Println(err)
	}
	err = ioutil.WriteFile(path+date+".json", rawData, 0755)
	if err != nil {
		log.Fatalln("wrong: ", err)
	}
}
