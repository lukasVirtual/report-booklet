package dbmodels

import (
	"fmt"
	"io/fs"
	"os"
	"runtime"
	"sort"
	"strconv"
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
	db.Table("users").Where("name = ?", nameUser).Update("parent_id", user.Model.ID)
}

func GetRelations(instructor string) []User {
	var user User
	var users []User

	db.Find(&user, "name = ?", instructor)
	db.Raw("select id, parent_id, name from users where parent_id = ? AND NOT name=?", user.Model.ID, instructor).Scan(&users)

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

	c1 := make(chan []fs.DirEntry)
	c2 := make(chan []fs.DirEntry)
	if runtime.GOOS == "windows" {
		wkhtmltopdf.SetPath(`C:/Program Files/wkhtmltopdf/bin/wkhtmltopdf`)
	}

	pdf, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		return nil, err
	}
	go func() {
		months, _ := os.ReadDir(destinationPath + "/Dokumente/TextFieldOutput")
		c1 <- months
	}()

	go func() {
		absence, _ := os.ReadDir(destinationPath + "/Dokumente/AbsenceStatus")
		c2 <- absence
	}()

	months := <-c1
	absence := <-c2
	close(c1)
	close(c2)

	sort.Slice(months, func(i, j int) bool {
		sl1 := strings.Split(months[i].Name(), ".")
		sl2 := strings.Split(months[j].Name(), ".")
		sl1Int, _ := strconv.Atoi(sl1[1])
		sl2Int, _ := strconv.Atoi(sl2[1])
		return sl1Int < sl2Int
	})

	var resString string

	resString = `<hmtl><header><h1 style="background: -webkit-linear-gradient(yellow, red); text-align:center">Report Booklet</h1></header><body>`
	for i, month := range months {
		fmt.Println(month.Name())
		absenceVal := ""
		if i < len(absence) {
			absenceVal = strings.ReplaceAll(absence[i].Name(), ".json", "")
		} else {
			absenceVal = ""
		}
		status, _ := jsonhandler.ReadStatusJson(absenceVal)
		opt := strings.ReplaceAll(month.Name(), ".json", "")
		output, _ := jsonhandler.ReadFromJson(opt)
		resString += `<div style="margin-top:180px; border:none">`
		resString += fmt.Sprintf(`<div style="margin:auto;width: 800px;height: 80px;border: 1px solid black"><h1 style="text-align:center; justify-content: center">%s <span style="float:right">%s</span></h1></div>`, opt, status.Status)
		resString += `<div style="margin:auto;min-height: 100px; width: 800px; border: 1px black">`
		for _, out := range output {
			resString += fmt.Sprintf(`<li>%s <span style="float:right">%s</span></li>`, out.Input, out.Time)
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
