package handler

import (
	"context"
	"net/http"
	"os"

	"github.com/thanaphon44881/go-testfirebase/adapter"
	"github.com/thanaphon44881/go-testfirebase/service"

	firebase "firebase.google.com/go"
	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/api/option"
)

var app *fiber.App

func init() {
	app = fiber.New()

	cred := os.Getenv("FIREBASE_CREDENTIALS")
	if cred == "" {
		panic("FIREBASE_CREDENTIALS is missing")
	}

	opt := option.WithCredentialsJSON([]byte(cred))

	conf := &firebase.Config{
		ProjectID: "apitest-2b035",
	}

	fbApp, err := firebase.NewApp(context.Background(), conf, opt)
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