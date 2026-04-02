package handler

import (
	"context"
	"net/http"
	"os"

	"TestUser/adapter"
	"TestUser/service"

	firebase "firebase.google.com/go"
	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/api/option"
)

var app *fiber.App

func init() {
	app = fiber.New()
	cred := os.Getenv("FIREBASE_CREDENTIALS")
	opt := option.WithCredentialsJSON([]byte(cred))

	fbApp, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		panic(err)
	}

	fireDB := adapter.NewFireDB(fbApp)
	adpDb := adapter.NewuserDB(fireDB)
	srv := service.NewService(adpDb)
	adphttp := adapter.Newhttpuser(srv)

	app.Post("/user", adphttp.CreatUser)
	app.Get("/users", adphttp.GetUsers)
	app.Get("/users/:id", adphttp.GetUserByID)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("API WORKING 🚀")
	})
}

func Handler(w http.ResponseWriter, r *http.Request) {
	adaptor.FiberApp(app)(w, r)
}
