package user

import (
	"context"

	"github.com/emregocer/golang_project/internal/model"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repo *Repo
}

func NewService(repo *Repo) *Service {
	return &Service{repo: repo}
}

func (service *Service) Create(ctx context.Context, username string, password string, email string) (*model.User, error) {
	exists, err := service.repo.CheckUserExists(ctx, username, email)

	if err != nil {
		return nil, err
	}

	if exists {
		return &model.User{}, err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.MinCost,
	)

	if err != nil {
		return nil, err
	}

	var user *model.User

	user, err = service.repo.Create(ctx, username, string(hashedPassword), email)
	if err != nil {
		return nil, err
	}

	return user, err
}

func (service *Service) Login(ctx context.Context, username string, password string) (*model.User, error) {
	user, err := service.repo.GetOneByUsername(ctx, username)

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, nil
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, err
	} else {
		return user, nil
	}
}
