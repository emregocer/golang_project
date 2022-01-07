package user

import (
	"context"
	"database/sql"

	"github.com/emregocer/golang_project/internal/model"
	"github.com/jmoiron/sqlx"
)

type Repo struct {
	db *sqlx.DB
}

func NewRepo(db *sqlx.DB) *Repo {
	return &Repo{db}
}

func (repo *Repo) GetOneByUsername(ctx context.Context, username string) (*model.User, error) {
	user := model.User{}

	err := repo.db.GetContext(ctx, &user, "SELECT * FROM users WHERE username=$1", username)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		} else {
			return nil, err
		}
	}

	return &user, nil
}

func (repo *Repo) Create(ctx context.Context, username string, password string, email string) (*model.User, error) {
	user := model.User{}

	statement := `
		INSERT INTO users(username, password, email) 
		VALUES($1,$2,$3) RETURNING *`

	err := repo.db.QueryRowContext(ctx, statement, username, string(password), email).Scan(
		&user.Id,
		&user.Username,
		&user.Password,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	user.Password = ""

	return &user, nil
}

func (repo *Repo) CheckUserExists(ctx context.Context, username string, email string) (bool, error) {
	var id int
	err := repo.db.QueryRowContext(ctx, `SELECT id FROM users WHERE username=$1 OR email=$2`, username, email).Scan(&id)

	if err != nil {
		// if the record doesn't exists
		if err == sql.ErrNoRows {
			return false, nil
			// if there was an other db error
		} else {
			return true, err
		}
	}

	return true, nil
}
