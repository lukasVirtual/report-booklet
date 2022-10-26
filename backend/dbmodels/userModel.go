package dbmodels

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	jsonhandler "www.github.com/backend/Jsonhandler"
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
	db.Unscoped().Where("name=?", name).Delete(&User{})
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

	db.Create(&textField)
}

func GetTextFieldData(date string) []TextField {
	var textField []TextField
	db.Unscoped().Raw("select calendar_date, input, time_stamp from text_fields where calendar_date=?", date).Scan(&textField)
	return textField
}

func UpdateId(nameInstructor, nameUser string) {
	var user User

	db.Find(&user, "name = ?", nameInstructor).Update("parent_id", user.Model.ID)
	fmt.Println(user.Model.ID)
	db.Table("users").Where("name = ?", nameUser).Update("parent_id", user.Model.ID)
}

func GetRelations(instructor string) []User {
	var user User
	var users []User

	db.Find(&user, "name = ?", instructor)
	// db.Debug().Find(&users, "parent_id = ?", user.Model.ID)
	db.Raw("select id, parent_id, name from users where parent_id = ? AND NOT name=?", user.Model.ID, instructor).Scan(&users)

	fmt.Println(users)
	return users
}

func DeleteUserFromInstructor(instructor string) {
	var user User
	db.Find(&user, "name = ?", instructor)

	db.Table("users").Where("parent_id = ?", user.Model.ID).Update("parent_id", 0)
}

func DelteSingleUserFromInstructor(client string) {
	db.Table("users").Where("name = ?", client).Update("parent_id", 0)
}

func FindInstructor(username string) string {
	var user User
	db.Find(&user, "name=?", username)
	var instructor User
	db.Where("parent_id=? AND role='instructor'", user.Parent_id).Find(&instructor)
	return instructor.Name
}

func ExportAsPdf() ([]byte, error) {
	var err error
	/*
	 *only needed for Windows
	 */
	if runtime.GOOS == "windows" {
		wkhtmltopdf.SetPath(`C:/Program Files/wkhtmltopdf/bin/wkhtmltopdf`)
	}

	pdf, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		return nil, err
	}

	months, _ := os.ReadDir(destinationPath + "/Dokumente/TextFieldOutput")
	/* TODO Probably need to do some sort of Quick Sort because
	currently it is sorting after the first digit [x].x.xxxx
	needs to be sorted after x.[x].xxxx
	*/
	var resString string

	resString = `<hmtl><header><h1 style="background: -webkit-linear-gradient(yellow, red); text-align:center">Report Booklet</h1></header><body>`
	for _, month := range months {
		fmt.Println(month.Name())

		opt := strings.ReplaceAll(month.Name(), ".json", "")
		output, _ := jsonhandler.ReadFromJson(opt)
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
