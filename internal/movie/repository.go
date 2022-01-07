package movie

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

func (repo *Repo) GetOne(ctx context.Context, id int) (*model.Movie, error) {
	movie := model.Movie{}

	movieStatement := `SELECT * FROM movies WHERE id = $1;`

	err := repo.db.GetContext(ctx, &movie, movieStatement, id)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		} else {
			return nil, err
		}
	}

	categories := []model.Category{}

	categoryStatement := `
		SELECT c.id, c.name, c.description
		FROM movie_category mc
		JOIN categories c on mc.category_id = c.id
		WHERE mc.movie_id = $1;
	`

	err = repo.db.SelectContext(ctx, &categories, categoryStatement, id)

	movie.Categories = categories

	if err != nil {
		return nil, err
	}

	return &movie, nil
}

func (repo *Repo) GetByCategoryId(ctx context.Context, categoryId int) ([]model.Movie, error) {
	movies := []model.Movie{}

	moviesByCategoryStatement := `
		SELECT m.* FROM movies m
		JOIN movie_category mc on m.id = mc.movie_id
		WHERE mc.category_id = $1; 
	`

	err := repo.db.SelectContext(ctx, &movies, moviesByCategoryStatement, categoryId)

	if err != nil {
		if err == sql.ErrNoRows {
			return movies, nil
		} else {
			return movies, err
		}
	}

	return movies, nil
}

func (repo *Repo) GetFavourites(ctx context.Context, userId int) ([]model.Movie, error) {
	movies := []model.Movie{}
	favouritedMoviesStatement := `
		SELECT m.id, m.name, m.plot FROM movies m
		join user_movie um on um.movie_id = m.id
		where user_id = $1 and m.is_deleted = FALSE;
	`

	err := repo.db.SelectContext(ctx, &movies, favouritedMoviesStatement, userId)

	if err != nil || len(movies) == 0 {
		return movies, err
	}

	movieIds := make([]int, 0, len(movies))
	for _, m := range movies {
		movieIds = append(movieIds, m.Id)
	}

	categoryStatement := `
		SELECT mc.movie_id, c.id, c.name
		FROM movie_category mc
		JOIN categories c on mc.category_id = c.id
		WHERE mc.movie_id IN (?);
	`

	query, args, err := sqlx.In(categoryStatement, movieIds)
	if err != nil {
		return movies, err
	}

	query = repo.db.Rebind(query)

	rows, err := repo.db.QueryContext(ctx, query, args...)
	if err != nil {
		return movies, err
	}

	movieCategories := make(map[int][]model.Category)

	for rows.Next() {
		var movieId int
		category := model.Category{}

		err := rows.Scan(&movieId, &category.Id, &category.Name)
		if err != nil {
			return movies, nil
		}

		movieCategories[movieId] = append(movieCategories[movieId], category)
	}

	moviesWithCategories := []model.Movie{}

	for _, m := range movies {
		m.Categories = movieCategories[m.Id]
		moviesWithCategories = append(moviesWithCategories, m)
	}

	return moviesWithCategories, nil
}

func (repo *Repo) Create(ctx context.Context, data CreateMovieRequest) (*model.Movie, error) {
	movie := model.Movie{}

	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	statement := `
		INSERT INTO movies(name, plot) 
		VALUES($1,$2) RETURNING *;`

	err = tx.QueryRowContext(ctx, statement, data.Name, data.Plot).Scan(
		&movie.Id,
		&movie.Name,
		&movie.Plot,
		&movie.IsDeleted,
		&movie.CreatedAt,
		&movie.UpdatedAt,
		&movie.DeletedAt,
	)

	if err != nil {
		return nil, err
	}

	for _, cId := range data.Categories {
		_, err = tx.ExecContext(ctx, `
			INSERT INTO movie_category (movie_id, category_id)
			VALUES ($1, $2);`, movie.Id, cId)

		if err != nil {
			return nil, err
		}
	}

	if err = tx.Commit(); err != nil {
		return &movie, err
	}

	return &movie, nil
}

func (repo *Repo) Update(ctx context.Context, id int, data UpdateMovieRequest) (*model.Movie, error) {
	movie := model.Movie{}

	statement := `
		UPDATE movies 
		SET name = $2, plot = $3
		WHERE id = $1 RETURNING *;`

	err := repo.db.QueryRowContext(ctx, statement, id, data.Name, data.Plot).Scan(
		&movie.Id,
		&movie.Name,
		&movie.Plot,
		&movie.IsDeleted,
		&movie.CreatedAt,
		&movie.UpdatedAt,
		&movie.DeletedAt,
	)

	if err != nil {
		return nil, err
	}

	return &movie, nil
}

func (repo *Repo) Delete(ctx context.Context, id int) (*model.Movie, error) {
	movie := model.Movie{}

	statement := `
		UPDATE movies 
		SET is_deleted = TRUE, deleted_at = $2
		WHERE id = $1 RETURNING *;`

	err := repo.db.QueryRowContext(ctx, statement, id, time.Now()).Scan(
		&movie.Id,
		&movie.Name,
		&movie.Plot,
		&movie.IsDeleted,
		&movie.CreatedAt,
		&movie.UpdatedAt,
		&movie.DeletedAt,
	)

	if err != nil {
		return nil, err
	}

	return &movie, nil
}

func (repo *Repo) AddMovieToFavourites(ctx context.Context, userId int, movieId int) (bool, error) {
	var id int

	selectStatement := `SELECT id FROM user_movie WHERE user_id = $1 AND movie_id = $2;`

	err := repo.db.Get(&id, selectStatement, userId, movieId)

	if err != nil {
		if err != sql.ErrNoRows {
			return false, err
		}
	} else {
		if id != 0 {
			return false, err
		}
	}

	statement := `
		INSERT INTO user_movie(user_id, movie_id, favourited_at) 
		VALUES($1, $2, $3);`

	_, err = repo.db.ExecContext(ctx, statement, userId, movieId, time.Now())

	if err != nil {
		return false, err
	}

	return true, err
}

func (repo *Repo) RemoveMovieFromFavourites(ctx context.Context, userId, movieId int) (bool, error) {

	statement := `
		DELETE FROM user_movie 
		WHERE user_id = $1 AND movie_id = $2;`

	res, err := repo.db.ExecContext(ctx, statement, userId, movieId)

	if err != nil {
		return false, err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return false, err
	}

	if count == 1 {
		return true, nil
	}

	return false, err
}
