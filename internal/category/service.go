package category

import (
	"context"

	"github.com/emregocer/golang_project/internal/model"
)

type Service struct {
	repo *Repo
}

func NewService(repo *Repo) *Service {
	return &Service{repo: repo}
}

func (c *Service) GetOne(ctx context.Context, id int) (*model.Category, error) {
	category, err := c.repo.GetOne(ctx, id)

	return category, err
}

func (c *Service) Get(ctx context.Context) ([]model.Category, error) {
	categories, err := c.repo.Get(ctx)

	return categories, err
}

func (c *Service) Create(ctx context.Context, req CreateCategoryRequest) (*model.Category, error) {
	category, err := c.repo.Create(ctx, req)

	return category, err
}

func (c *Service) Update(ctx context.Context, id int, req UpdateCategoryRequest) (*model.Category, error) {
	category, err := c.repo.Update(ctx, id, req)

	return category, err
}

func (c *Service) Delete(ctx context.Context, id int) (*model.Category, error) {
	category, err := c.repo.Delete(ctx, id)

	return category, err
}
