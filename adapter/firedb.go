package adapter

import (
	"context"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
)

type FireDB struct {
	Client *firestore.Client
}

func NewFireDB(app *firebase.App) *FireDB {
	client, err := app.Firestore(context.Background())
	if err != nil {
		panic(err)
	}

	return &FireDB{
		Client: client,
	}
}
