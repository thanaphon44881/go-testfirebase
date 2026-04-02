package adapter

import (
	"github.com/thanaphon44881/go-testfirebase/repository"
	"context"
	"strconv"
)

type userdb struct {
	db *FireDB
}

func NewuserDB(db *FireDB) repository.RepositoryUser {
	return &userdb{db: db}
}

func (u userdb) Save(user repository.User) error {
	ctx := context.Background()

	_, _, err := u.db.Client.Collection("users").Add(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (u userdb) FindAll() ([]repository.User, error) {
	ctx := context.Background()

	docs, err := u.db.Client.Collection("users").Documents(ctx).GetAll()
	if err != nil {
		return nil, err
	}

	var users []repository.User
	for _, doc := range docs {
		var user repository.User
		err := doc.DataTo(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (u userdb) FindByID(id string) (*repository.User, error) {
	ctx := context.Background()

	doc, err := u.db.Client.Collection("users").Doc(id).Get(ctx)
	if err != nil {
		return nil, err
	}

	var user repository.User

	err = doc.DataTo(&user)
	if err != nil {
		return nil, err
	}

	idInt, _ := strconv.Atoi(doc.Ref.ID)
	user.ID = idInt

	return &user, nil
}
