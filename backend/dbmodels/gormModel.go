package dbmodels

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

type JsonData struct {
	Id    int    `json: "id"`
	Date  string `json: "date"`
	Input string `json: "Input"`
	Time  string `json: "time"`
	Rows  int    `json: "rows"`
}

type StatusJson struct {
	Date   string `json: "date"`
	Status string `json: "status"`
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

	db.Raw("select name, password, role from users where name = ?", name).Scan(&user)

	return true
}

func GetAllData() []User {
	var user []User
	db.Raw("select name, parent_id, role from users").Scan(&user)
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

func SaveAsJson(id int, date, input, time string, rows int) {
	m := []JsonData{}
	obj := JsonData{Id: id, Date: date, Input: input, Time: time, Rows: rows}
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
				if obj.Time == "00:00" {
					output[i].Input = obj.Input
					output[i].Rows = obj.Rows
				} else if obj.Input == "" {
					output[i].Time = obj.Time
				} else if obj.Time != "00:00" && obj.Input != "" {
					output[i].Input = obj.Input
					output[i].Time = obj.Time
				}
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
	//fmt.Println(local)
	for i, _ := range local {
		local[i].Id = i
	}
	rawData, _ := json.MarshalIndent(local, " ", " ")

	ioutil.WriteFile(destinationPath+"/Documents/TextFieldOutput/"+date+".json", rawData, 0755)

}

type Form struct {
	Status string
}

func WriteStatusJson(date, status string) {
	if _, err := ioutil.ReadDir(destinationPath + "/Documents/AbsenceStatus/"); err != nil {
		fmt.Println("Creating Directory...")
		os.Mkdir(destinationPath+"/Documents/AbsenceStatus/", 0755)
	}
	rec := Form{Status: status}
	data, _ := json.MarshalIndent(rec, " ", " ")

	ioutil.WriteFile(destinationPath+"/Documents/AbsenceStatus/"+date+".json", data, 0755)
}

func ReadStatusJson(date string) (StatusJson, error) {
	data, err := ioutil.ReadFile(destinationPath + "/Documents/AbsenceStatus/" + date + ".json")
	if err != nil {
		return StatusJson{}, err
	}
	output := StatusJson{}
	err = json.Unmarshal(data, &output)
	if err != nil {
		log.Fatalf("error unmarshal json: %v", err)
	}

	return output, nil
}

func WriteQualificationsJson(qualis []interface{}, date, endpoint string) {
	if _, err := ioutil.ReadDir(destinationPath + endpoint); err != nil {
		fmt.Println("Creating Directory...")
		os.Mkdir(destinationPath+endpoint, 0755)
	}

	data, _ := json.MarshalIndent(qualis, " ", " ")
	ioutil.WriteFile(destinationPath+endpoint+date+".json", data, 0755)
}

type QualificationFormReturn struct {
	Title string
	State bool
}

func ReadQualificationsJson(date, endpoint string) []QualificationFormReturn {
	data, _ := ioutil.ReadFile(destinationPath + endpoint + date + ".json")

	var qualis []QualificationFormReturn
	_ = json.Unmarshal(data, &qualis)

	return qualis
}

func UpdateId(nameInstructor, nameUser string) {
	var user User

	db.Debug().Find(&user, "name = ?", nameInstructor)
	fmt.Println(user.Model.ID)
	db.Debug().Table("users").Where("name = ?", nameUser).Update("parent_id", user.Model.ID)
}

func GetRelations(instructor string) []User {
	var user User
	var users []User

	db.Debug().Find(&user, "name = ?", instructor)
	// db.Debug().Find(&users, "parent_id = ?", user.Model.ID)
	db.Raw("select id, parent_id, name from users where parent_id = ?", user.Model.ID).Scan(&users)

	fmt.Println(users)
	return users
}

func DeleteUserFromInstructor(instructor string) {
	var user User
	db.Debug().Find(&user, "name = ?", instructor)

	db.Debug().Table("users").Where("parent_id = ?", user.Model.ID).Update("parent_id", 0)
}

func DelteSingleUserFromInstructor(client string) {
	db.Debug().Table("users").Where("name = ?", client).Update("parent_id", 0)
}

func ExportAsPdf() ([]byte, error) {
	var err error

	pdf, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		return nil, err
	}

	months, _ := ioutil.ReadDir(destinationPath + "/Documents/TextFieldOutput")

	/* TODO
	Probably need to do some sort of Quick Sort because
	currently it is sorting after the first digit [x].x.xxxx
	needs to be sorted after x.[x].xxxx
	*/

	var resString string
	resString = `<hmtl><header><h1 style="background: -webkit-linear-gradient(yellow, red); text-align:center">Report Booklet</h1></header><body>`
	for _, month := range months {
		fmt.Println(month.Name())

		opt := strings.ReplaceAll(month.Name(), ".json", "")
		output, _ := ReadFromJson(opt)
		resString += `<div style="margin-top:180px; border:none">`
		resString += fmt.Sprintf(`<div style="margin:auto;width: 800px;height: 80px;border: 1px solid black"><h1 style="text-align:center; justify-content: center">%s</h1></div>`, opt)
		resString += `<div style="margin:auto;min-height: 100px; width: 800px; border: 1px black">`
		for _, out := range output {
			resString += fmt.Sprintf(`<li>%s</li>`, out.Input)
		}
	}
	resString += "</div>"
	resString += "</div>"
	resString += "</body></html>"
	page := wkhtmltopdf.NewPageReader(strings.NewReader(resString))
	pdf.AddPage(page)

	pdf.MarginLeft.Set(0)
	pdf.MarginRight.Set(0)
	pdf.Dpi.Set(300)
	pdf.PageSize.Set(wkhtmltopdf.PageSizeA4)
	pdf.Orientation.Set(wkhtmltopdf.OrientationLandscape)

	err = pdf.Create()
	if err != nil {
		return nil, err
	}
	return pdf.Bytes(), nil
}
