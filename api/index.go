package handler

import (
	"context"
	"net/http"
	"os"
	"strings"
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

	projectID := os.Getenv("FIREBASE_PROJECT_ID")
	clientEmail := os.Getenv("FIREBASE_CLIENT_EMAIL")
	privateKey := os.Getenv("FIREBASE_PRIVATE_KEY")

	conf := &firebase.Config{
		ProjectID: projectID,
	}

	opt := option.WithCredentialsJSON([]byte(fmt.Sprintf(`{
		"type": "service_account",
		"project_id": "%s",
		"private_key": "%s",
		"client_email": "%s",
		"token_uri": "https://oauth2.googleapis.com/token"
	}`, projectID, privateKey, clientEmail)))

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