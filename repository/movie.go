package repository

import (
	"context"
	"errors"

	frameworks "github.com/hendriam/movie-service/framework"
	"github.com/hendriam/movie-service/model/domain"
)

type MovieRepository struct {
	db frameworks.Database
}

func NewMovieRepository(db frameworks.Database) MovieRepository {
	return MovieRepository{db: db}
}

func (repository *MovieRepository) Insert(ctx context.Context, movie domain.Movie) (domain.Movie, error) {
	insert, err := repository.db.ExecContext(
		ctx,
		"INSERT INTO movies (title, year) VALUE (?, ?)",
		movie.Title,
		movie.Year,
	)
	if err != nil {
		return movie, err
	}

	id, err := insert.LastInsertId()
	if err != nil {
		return movie, err
	}

	movie.ID = int(id)

	return movie, nil
}

func (repository *MovieRepository) Update(ctx context.Context, movie domain.Movie, movieId int) (int64, error) {
	update, err := repository.db.ExecContext(
		ctx,
		"UPDATE movies SET title=?, year=? WHERE id=?",
		movie.Title,
		movie.Year,
		movieId,
	)
	if err != nil {
		return 0, err
	}

	row, _ := update.RowsAffected()

	return row, nil
}

func (repository *MovieRepository) Delete(ctx context.Context, movieId int) error {
	_, err := repository.db.ExecContext(
		ctx,
		"DELETE FROM movies WHERE id=?",
		movieId,
	)
	if err != nil {
		return err
	}

	return nil
}

func (repository *MovieRepository) FindById(ctx context.Context, movieId int) (domain.Movie, error) {
	movie := domain.Movie{}
	err := repository.db.QueryRowContext(
		ctx,
		"SELECT * FROM movies WHERE id=?",
		movieId,
	).Scan(&movie.ID, &movie.Title, &movie.Year)
	if err != nil {
		return movie, errors.New("movie not found")
	}

	return movie, nil
}
