package repository

import (
	"context"
	"fmt"

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
	result, err := repository.db.ExecContext(
		ctx,
		"INSERT INTO movie (title, year) VALUE (?, ?)",
		movie.Title,
		movie.Year,
	)
	if err != nil {
		fmt.Println("[REPOSITORY] error =>", err)
		return domain.Movie{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		fmt.Println("[REPOSITORY] error =>", err)
		return domain.Movie{}, err
	}

	return domain.Movie{
		ID:    int(id),
		Title: movie.Title,
		Year:  movie.Year,
	}, nil
}
