package movie

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

func (c *Service) GetOne(ctx context.Context, id int) (*model.Movie, error) {
	movie, err := c.repo.GetOne(ctx, id)

	return movie, err
}

func (c *Service) GetByCategoryId(ctx context.Context, categoryId int) ([]model.Movie, error) {
	movies, err := c.repo.GetByCategoryId(ctx, categoryId)

	return movies, err
}

func (c *Service) GetFavourites(ctx context.Context, userId int) ([]model.Movie, error) {
	movies, err := c.repo.GetFavourites(ctx, userId)

	return movies, err
}

func (c *Service) Create(ctx context.Context, req CreateMovieRequest) (*model.Movie, error) {
	movie, err := c.repo.Create(ctx, req)

	return movie, err
}

func (c *Service) Update(ctx context.Context, id int, req UpdateMovieRequest) (*model.Movie, error) {
	movie, err := c.repo.Update(ctx, id, req)

	return movie, err
}

func (c *Service) Delete(ctx context.Context, id int) (*model.Movie, error) {
	movie, err := c.repo.Delete(ctx, id)

	return movie, err
}

func (c *Service) AddMovieToFavourites(ctx context.Context, userId int, movieId int) (bool, error) {
	res, err := c.repo.AddMovieToFavourites(ctx, userId, movieId)

	return res, err
}

func (c *Service) RemoveMovieFromFavourites(ctx context.Context, userId int, movieId int) (bool, error) {
	res, err := c.repo.RemoveMovieFromFavourites(ctx, userId, movieId)

	return res, err
}
