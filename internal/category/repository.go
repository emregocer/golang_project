package category

import (
	"context"
	"database/sql"
	"time"

	"github.com/emregocer/golang_project/internal/model"
	"github.com/jmoiron/sqlx"
)

type Repo struct {
	db *sqlx.DB
}

func NewRepo(db *sqlx.DB) *Repo {
	return &Repo{db}
}

func (repo *Repo) GetOne(ctx context.Context, id int) (*model.Category, error) {
	category := model.Category{}

	err := repo.db.GetContext(ctx, &category, "SELECT * FROM categories WHERE id=$1", id)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		} else {
			return nil, err
		}
	}

	return &category, nil
}

func (repo *Repo) Get(ctx context.Context) ([]model.Category, error) {
	categories := []model.Category{}

	err := repo.db.SelectContext(ctx, &categories, "SELECT * FROM categories")

	if err != nil {
		if err == sql.ErrNoRows {
			return categories, nil
		} else {
			return categories, err
		}
	}

	return categories, nil
}

func (repo *Repo) Create(ctx context.Context, data CreateCategoryRequest) (*model.Category, error) {
	category := model.Category{}

	statement := `
		INSERT INTO categories(name, description) 
		VALUES($1,$2) RETURNING *`

	err := repo.db.QueryRowContext(ctx, statement, data.Name, data.Description).Scan(
		&category.Id,
		&category.Name,
		&category.Description,
		&category.IsDeleted,
		&category.CreatedAt,
		&category.UpdatedAt,
		&category.DeletedAt,
	)

	if err != nil {
		return nil, err
	}

	return &category, nil
}

func (repo *Repo) Update(ctx context.Context, id int, data UpdateCategoryRequest) (*model.Category, error) {
	category := model.Category{}

	statement := `
		UPDATE categories 
		SET name = $2, description = $3
		WHERE id = $1 RETURNING *`

	err := repo.db.QueryRowContext(ctx, statement, id, data.Name, data.Description).Scan(
		&category.Id,
		&category.Name,
		&category.Description,
		&category.IsDeleted,
		&category.CreatedAt,
		&category.UpdatedAt,
		&category.DeletedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return &category, nil
}

func (repo *Repo) Delete(ctx context.Context, id int) (*model.Category, error) {
	category := model.Category{}

	statement := `
		UPDATE categories 
		SET is_deleted = TRUE, deleted_at = $2
		WHERE id = $1 RETURNING *`

	err := repo.db.QueryRowContext(ctx, statement, id, time.Now()).Scan(
		&category.Id,
		&category.Name,
		&category.Description,
		&category.IsDeleted,
		&category.CreatedAt,
		&category.UpdatedAt,
		&category.DeletedAt,
	)

	if err != nil {
		return nil, err
	}

	return &category, nil
}
