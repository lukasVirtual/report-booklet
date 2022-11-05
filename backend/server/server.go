package server

import (
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/websocket/v2"
	_ "github.com/gofiber/websocket/v2"
	// "www.github.com/backend/mongodb"
)

var (
	store    *session.Store
	AUTH_KEY string = "authenticated"
	USER_ID  string = "ID"
)

type User struct {
	Name     string `json: "name"`
	Password string `json: "password"`
	Role     string `json: "role"`
}

type SubmittedData struct {
	Name string        `json: "name"`
	Date string        `json: "date"`
	Data []interface{} `json: "data"`
}

type QualificationForm struct {
	Date           string        `json: "date"`
	Qualifications []interface{} `json: "qualifications"`
}

func Init() {
	router := fiber.New()

	store = session.New(session.Config{
		CookieHTTPOnly: true,
		Expiration:     time.Hour * 5,
	})

	router.Use(NewMiddleware(), cors.New((cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "*",
		AllowHeaders:     "Access-Control-Allow-Origin, Content-Type, Origin, Accept",
	})))

	router.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	//TODO: Make a better REST API.

	router.Post("/api/register", Register)
	router.Post("/api/login", Login)
	router.Post("/api/logout", Logout)
	router.Post("/api/delete", DeleteElem)
	router.Post("/api/insertinput", InsertInputField)
	router.Post("/api/inserttext", InsertTextField)
	router.Post("/api/writeJson", WriteJson)
	router.Post("/api/readJson", ReadJson)
	router.Post("/api/removeJson", RemoveJson)
	router.Post("/api/getTextFieldData", GetTextFieldData)
	router.Post("/api/saveSubmittedData", SaveSubmittedData)
	router.Post("/api/retrieveSubmittedData", RetrieveSubmittedData)
	router.Post("/api/writeStatusJson", WriteStatus)
	router.Post("/api/readStatusJson", ReadStatus)
	router.Post("/api/assigne", AssignUsertoInstructor)
	router.Post("/api/delteUserInstructor", DelteUserInstructor)
	router.Post("/api/removeUserInstructor", RemoveSingleUserFromInstructor)
	router.Post("/api/returnUsers", GetUserBelongInstructor)
	router.Post("/api/insertQualifications", InsertQualifications)
	router.Post("/api/getQualifications", GetQualifications)
	router.Post("/api/insertcurriculum", InsertCurriculum)
	router.Post("/api/readcurriculum", GetCurriculum)
	router.Post("/api/findInstructor", GetInstructor)
	router.Post("/api/readJsonMonth", ReadJsonForMonth)
	router.Get("/api/statuscheck", StatusCheck)
	router.Get("/api/dataware", DataWare)
	// router.Get("/ws/helloworld", ReadMessages)
	router.Get("/api/createdpdf", CreatePdf)
	router.Get("/api/user", GetUserData)

	router.Listen("127.0.0.1:5000")

}

func NewMiddleware() fiber.Handler {
	return AuthMiddleware
}

func AuthMiddleware(c *fiber.Ctx) error {
	sess, err := store.Get(c)
	if strings.Split(c.Path(), "/")[1] == "api" || strings.Split(c.Path(), "/")[1] == "ws" {
		return c.Next()
	}
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Nicht Angemeldet",
		})
	}

	if sess.Get(AUTH_KEY) == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Nicht Angemeldet",
		})
	}

	return c.Next()
}
