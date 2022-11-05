package server

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	jsonhandler "www.github.com/backend/Jsonhandler"
	"www.github.com/backend/dbmodels"
	serverfiles "www.github.com/backend/server_files"
)

func Register(c *fiber.Ctx) error {
	c.Accepts("application/json")

	var data User
	// fmt.Println("data", data.Name)
	err := c.BodyParser(&data)

	// fmt.Println("data", data)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "1" + err.Error(),
		})
	}

	password, cryptErr := bcrypt.GenerateFromPassword([]byte(data.Password), 14)
	if cryptErr != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "2",
		})
	}

	user := dbmodels.User{
		Name:     data.Name,
		Password: string(password),
		Role:     data.Role,
	}
	dbmodels.CreateUser(&user)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "registered",
	})
}

func Login(c *fiber.Ctx) error {
	var data User

	err := c.BodyParser(&data)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "the pointer " + err.Error(),
		})
	}

	var user dbmodels.User

	if !dbmodels.CheckUser(data.Name, &user) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "1" + err.Error(),
		})
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Nicht Regestriert",
		})
	}

	sess, sessErr := store.Get(c)

	if sessErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error: " + err.Error(),
		})
	}

	sess.Set(AUTH_KEY, true)
	sess.Set(USER_ID, user.Name)
	sessErr = sess.Save()
	if sessErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error: " + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "logged in",
	})
}

func Logout(c *fiber.Ctx) error {
	sess, err := store.Get(c)

	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "no session available" + err.Error(),
		})
	}

	err = sess.Destroy()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error: " + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "no session available",
	})
}

func StatusCheck(c *fiber.Ctx) error {
	sess, err := store.Get(c)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Not Authorized",
		})
	}

	auth := sess.Get(AUTH_KEY)
	if auth != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "authenticated",
		})
	} else {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "not authorized",
		})
	}
}

func GetUserData(c *fiber.Ctx) error {
	sess, err := store.Get(c)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Not Authorized",
		})
	}

	if sess.Get(AUTH_KEY) == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Not Authorized",
		})
	}

	userId := sess.Get(USER_ID)
	if userId == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Not Authorized",
		})
	}

	var user dbmodels.User = dbmodels.GetUser(fmt.Sprint(userId))
	// if err != nil {
	// 	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
	// 		"message": "Not Authorized",
	// 	})
	// }

	return c.Status(fiber.StatusOK).JSON(user)
}

func DataWare(c *fiber.Ctx) error {
	_, err := store.Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error: " + err.Error(),
		})
	}
	value := dbmodels.GetAllData()

	return c.Status(fiber.StatusOK).JSON(value)
}

func DeleteElem(c *fiber.Ctx) error {
	var data User
	err := c.BodyParser(&data)

	if err != nil {
		fmt.Println("Something went wrong")
	}

	dbmodels.DeleteItemAt(data.Name)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Successfully Deleted",
	})
}

var inpField dbmodels.InputField

func InsertInputField(c *fiber.Ctx) error {
	var input dbmodels.InputField
	err := c.BodyParser(&input)

	if err != nil {
		fmt.Println("Could not read input data")
	}

	inpField = dbmodels.InsertInputField(input.Input, input.TimeStamp)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Successfully inserted InputField",
	})
}

func InsertTextField(c *fiber.Ctx) error {
	var textField dbmodels.TextField
	err := c.BodyParser(&textField)
	if err != nil {
		fmt.Println("Could not read TextField data")
	}

	dbmodels.InsertTextField(textField.CalendarDate, &inpField)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Successfully inserted TextField",
	})
}

func GetTextFieldData(c *fiber.Ctx) error {
	var textField dbmodels.TextField
	err := c.BodyParser(&textField)

	if err != nil {
		fmt.Println("Could not read Date")
	}

	dates := dbmodels.GetTextFieldData(textField.CalendarDate)

	return c.Status(fiber.StatusOK).JSON(dates)
}

func WriteJson(c *fiber.Ctx) error {
	var jdata jsonhandler.JsonData
	err := c.BodyParser(&jdata)

	if err != nil {
		fmt.Println("Could not read data")
	}
	//fmt.Println(jdata.Date)
	jsonhandler.SaveAsJson(jdata.Id, jdata.Date, jdata.Input, jdata.Time, jdata.Rows)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Successfully created Json",
	})
}

func ReadJson(c *fiber.Ctx) error {
	var jsonOutput jsonhandler.JsonData
	err := c.BodyParser(&jsonOutput)

	if err != nil {
		fmt.Println(err)
	}

	// fmt.Println(jsonOutput.Date)

	output, _ := jsonhandler.ReadFromJson(jsonOutput.Date)
	// fmt.Println(output)
	// if err != nil {
	// 	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
	// 		"message": err,
	// 	})
	// }
	return c.Status(fiber.StatusOK).JSON(output)
}

type RemoveJsonStruct struct {
	Index int    `json: "index"`
	Date  string `json: "date"`
}

func RemoveJson(c *fiber.Ctx) error {
	var jsonOutput RemoveJsonStruct
	err := c.BodyParser(&jsonOutput)

	if err != nil {
		log.Fatalf("something went wrong parsing json: %v", err)
	}

	jsonhandler.RemoveFromJson(jsonOutput.Date, jsonOutput.Index)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "successfully removed from Json",
	})

}

func InsertQualifications(c *fiber.Ctx) error {
	var qualis QualificationForm
	err := c.BodyParser(&qualis)
	if err != nil {
		log.Fatalf("something went wrong parsing json: %v", err)
	}
	jsonhandler.WriteQualificationsJson(qualis.Qualifications, qualis.Date, "/Dokumente/Qualifications/")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "successfully inserted data into mongodb",
	})
}

func GetQualifications(c *fiber.Ctx) error {
	var qualis QualificationForm
	err := c.BodyParser(&qualis)
	if err != nil {
		log.Fatalf("something went wrong parsing json: %v", err)
	}
	output := jsonhandler.ReadQualificationsJson(qualis.Date, "/Dokumente/Qualifications/")
	return c.Status(fiber.StatusOK).JSON(output)
}

func WriteStatus(c *fiber.Ctx) error {
	var inputData jsonhandler.StatusJson
	err := c.BodyParser(&inputData)
	if err != nil {
		log.Fatalf("Could not read data: %v", err)
	}

	jsonhandler.WriteStatusJson(inputData.Date, inputData.Status)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "successfully inserted status",
	})
}

func ReadStatus(c *fiber.Ctx) error {
	var inputData jsonhandler.StatusJson
	err := c.BodyParser(&inputData)
	if err != nil {
		log.Fatalf("Could not read data: %v", err)
	}

	out, _ := jsonhandler.ReadStatusJson(inputData.Date)

	return c.Status(fiber.StatusOK).JSON(out)
}

type UserInstructor struct {
	UsersName       string `json: "usersName`
	InstructorsName string `json: "instructorsName`
}

func AssignUsertoInstructor(c *fiber.Ctx) error {
	var data UserInstructor
	err := c.BodyParser(&data)
	if err != nil {
		log.Fatalf("Could not parse data from json: %v", err)
	}

	dbmodels.UpdateId(data.InstructorsName, data.UsersName)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "successfully updated user",
	})
}

func GetUserBelongInstructor(c *fiber.Ctx) error {
	var data UserInstructor
	err := c.BodyParser(&data)
	if err != nil {
		log.Fatalf("Could not parse data from json: %v", err)
	}

	users := dbmodels.GetRelations(data.InstructorsName)

	return c.Status(fiber.StatusOK).JSON(users)
}

func DelteUserInstructor(c *fiber.Ctx) error {
	var data UserInstructor

	err := c.BodyParser(&data)
	if err != nil {
		log.Fatalf("Could not parse data from json: %v", err)
	}

	dbmodels.DeleteUserFromInstructor(data.InstructorsName)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "successfully deleted users relation to instructor",
	})
}

func RemoveSingleUserFromInstructor(c *fiber.Ctx) error {
	var data UserInstructor
	err := c.BodyParser(&data)
	if err != nil {
		log.Fatalf("Could not parse data from json: %v", err)
	}

	dbmodels.DelteSingleUserFromInstructor(data.UsersName)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "successfully deleted user relation to instructor",
	})
}

// var ReadMessages = websocket.New(func(c *websocket.Conn) {
// 	var (
// 		mt  int
// 		msg []byte
// 		err error
// 	)

// 	for {
// 		if mt, msg, err = c.ReadMessage(); err != nil {
// 			log.Println("read: ", err)
// 			break
// 		}

// 		log.Printf("message recieved: %s", msg)

// 		if err = c.WriteMessage(mt, msg); err != nil {
// 			log.Println("write:", err)
// 			break
// 		}
// 	}
// })

func CreatePdf(c *fiber.Ctx) error {
	data, err := dbmodels.ExportAsPdf()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not Export",
		})
	}

	return c.Status(fiber.StatusOK).JSON(data)
}

func InsertCurriculum(c *fiber.Ctx) error {
	var data QualificationForm
	err := c.BodyParser(&data)
	if err != nil {
		log.Fatalf("Could not parse data from json: %v", err)
	}

	jsonhandler.WriteQualificationsJson(data.Qualifications, data.Date, "/Dokumente/Curriculum/")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Successfully saved Curriculum",
	})
}

func GetCurriculum(c *fiber.Ctx) error {
	var data QualificationForm
	err := c.BodyParser(&data)
	if err != nil {
		log.Fatalf("Could not parse data from json: %v", err)
	}

	out := jsonhandler.ReadQualificationsJson(data.Date, "/Dokumente/Curriculum/")

	return c.Status(fiber.StatusOK).JSON(out)
}

func GetInstructor(c *fiber.Ctx) error {
	var data User
	err := c.BodyParser(&data)
	if err != nil {
		log.Fatalf("Could not parse data from json: %v", err)
	}

	instructor := dbmodels.FindInstructor(data.Name)

	return c.Status(fiber.StatusOK).JSON(instructor)
}

func ReadJsonForMonth(c *fiber.Ctx) error {
	var data jsonhandler.JsonData
	err := c.BodyParser(&data)
	if err != nil {
		log.Fatalf("Could not parse data from json: %v", err)
	}

	output, err := jsonhandler.ReadJsonMonth(data.Date)
	if err != nil {
		log.Fatalf("Error while reading Json for the Month %s, error: %v", data.Date, err)
	}

	return c.Status(fiber.StatusOK).JSON(output)
}

func SaveSubmittedData(c *fiber.Ctx) error {
	var data SubmittedData
	err := c.BodyParser(&data)
	if err != nil {
		log.Fatalf("Could not parse data from json: %v", err)
	}
	// data.Data = "{ 'hello': 'world'}"
	// fmt.Println(data.Data)
	// servercom.SaveSubmittedData(data.Name, data.Date, data.Data)
	serverfiles.SaveSubmittedData(data.Name, data.Date, data.Data)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "successfully inserted data",
	})
}

func RetrieveSubmittedData(c *fiber.Ctx) error {
	var data SubmittedData
	err := c.BodyParser(&data)
	if err != nil {
		log.Fatalf("Could not parse data from json: %v", err)
	}

	// output := servercom.RetrieveSubmittedData(data.Name, data.Date)
	output := serverfiles.RetrieveSubmittedData(data.Name, data.Date)

	return c.Status(fiber.StatusOK).JSON(output)
}
