package repository

type User struct {
	ID    int    `json:"id" firestore:"id"`
	Name  string `json:"name" firestore:"name"`
	Lname string `json:"lname" firestore:"lname"`
	Email string `json:"email" firestore:"email"`
	Phon  int    `json:"phon" firestore:"phon"`
}

type RepositoryUser interface {
	Save(user User) error
	FindAll() ([]User, error)
	FindByID(id string) (*User, error)
}
