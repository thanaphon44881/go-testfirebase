package service

import (
	"TestUser/repository"
)

type ServiceUser interface {
	Creat(user repository.User) error
	GetUsers() ([]repository.User, error)
	GetUserByID(id string) (*repository.User, error)
}

type repo struct {
	repo repository.RepositoryUser
}

func NewService(r repository.RepositoryUser) ServiceUser {
	return repo{repo: r}
}

func (r repo) Creat(user repository.User) error {
	err := r.repo.Save(user)
	if err != nil {
		return err
	}
	return nil
}

func (s repo) GetUsers() ([]repository.User, error) {
	return s.repo.FindAll()
}

func (s repo) GetUserByID(id string) (*repository.User, error) {
	return s.repo.FindByID(id)
}
