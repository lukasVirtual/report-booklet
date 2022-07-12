package server

import (
	"fmt"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	_ "github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/session"
	"golang.org/x/crypto/bcrypt"
	"www.github.com/backend/dbmodels"
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

	router.Post("/api/register", Register)
	router.Post("/api/login", Login)
	router.Post("/api/logout", Logout)
	router.Get("/api/statuscheck", StatusCheck)
	router.Get("/api/dataware", DataWare)
	router.Post("/api/delete", DeleteElem)

	router.Get("/api/user", GetUserData)

	router.Listen(":5000")

}

func NewMiddleware() fiber.Handler {
	return AuthMiddleware
}

func AuthMiddleware(c *fiber.Ctx) error {
	sess, err := store.Get(c)
	if strings.Split(c.Path(), "/")[1] == "api" {
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

func Register(c *fiber.Ctx) error {
	c.Accepts("application/json")

	var data User
	fmt.Println("data", data.Name)
	err := c.BodyParser(&data)

	fmt.Println("data", data)

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

	fmt.Println("data model", data.Name)
	if !dbmodels.CheckUser(data.Name, &user) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "1" + err.Error(),
		})
	}
	fmt.Println("user", user.Name, user.Password, user.Role)
	fmt.Println("user and data: ", user.Password, data.Password)
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
	fmt.Println("session Error: ", sessErr)
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
	fmt.Println(auth)
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

	var user dbmodels.User

	user = dbmodels.GetUser(fmt.Sprint(userId))
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
		"message": "Successfull Delted",
	})
}
