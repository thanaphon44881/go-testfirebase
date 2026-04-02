package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"strings"

	"github.com/thanaphon44881/go-testfirebase/adapter"
	"github.com/thanaphon44881/go-testfirebase/service"
	"github.com/gofiber/fiber/v2/middleware/cors"
	firebase "firebase.google.com/go"
	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/api/option"
)

type FirebaseCred struct {
	Type        string `json:"type"`
	ProjectID   string `json:"project_id"`
	PrivateKey  string `json:"private_key"`
	ClientEmail string `json:"client_email"`
	TokenURI    string `json:"token_uri"`
}

var app *fiber.App

func init() {
	app = fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", 
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowMethods: "GET,POST,PUT,DELETE",
	}))
	projectID := os.Getenv("FIREBASE_PROJECT_ID")
	clientEmail := os.Getenv("FIREBASE_CLIENT_EMAIL")
	privateKey := os.Getenv("FIREBASE_PRIVATE_KEY")

	if privateKey == "" {
		panic("PRIVATE KEY MISSING")
	}

	privateKey = strings.ReplaceAll(privateKey, "\\n", "\n")

	cred := FirebaseCred{
		Type:        "service_account",
		ProjectID:   projectID,
		PrivateKey:  privateKey,
		ClientEmail: clientEmail,
		TokenURI:    "https://oauth2.googleapis.com/token",
	}

	credBytes, err := json.Marshal(cred)
	if err != nil {
		panic(err)
	}

	opt := option.WithCredentialsJSON(credBytes)

	conf := &firebase.Config{
		ProjectID: projectID,
	}

	fbApp, err := firebase.NewApp(context.Background(), conf, opt)
	if err != nil {
		panic(err)
	}

	fireDB := adapter.NewFireDB(fbApp)
	adpDb := adapter.NewuserDB(fireDB)
	srv := service.NewService(adpDb)
	adphttp := adapter.Newhttpuser(srv)

	app.Get("/api/users", adphttp.GetUsers)
	app.Post("/api/user", adphttp.CreatUser)
	app.Get("/api/users/:id", adphttp.GetUserByID)

	app.Get("/api/", func(c *fiber.Ctx) error {
		return c.SendString("API WORKING 🚀")
	})
}

func Handler(w http.ResponseWriter, r *http.Request) {
	adaptor.FiberApp(app)(w, r)
}
