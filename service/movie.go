package service

import (
	"context"

	"github.com/hendriam/movie-service/model/domain"
	"github.com/hendriam/movie-service/model/web"
	"github.com/hendriam/movie-service/repository"
)

type MovieService struct {
	movieRepository repository.MovieRepository
}

func NewMovieService(movieRepository repository.MovieRepository) MovieService {
	return MovieService{movieRepository: movieRepository}
}

func (service *MovieService) Create(ctx context.Context, movieRequest web.MovieCreateRequestBody) (domain.Movie, error) {
	movie := domain.Movie{
		Title: movieRequest.Title,
		Year:  movieRequest.Year,
	}

	responseCreate, err := service.movieRepository.Insert(ctx, movie)
	if err != nil {
		return responseCreate, err
	}

	return responseCreate, nil
}

func (service *MovieService) Update(ctx context.Context, movieRequest web.MovieUpdateRequestBody, movieId int) (domain.Movie, error) {
	movie := domain.Movie{
		Title: movieRequest.Title,
		Year:  movieRequest.Year,
	}

	_, err := service.movieRepository.Update(ctx, movie, movieId)
	if err != nil {
		return movie, err
	}

	movie.ID = movieId

	return movie, nil
}

func (service *MovieService) Delete(ctx context.Context, movieId int) error {
	err := service.movieRepository.Delete(ctx, movieId)
	if err != nil {
		return err
	}
	return nil
}

func (service *MovieService) FindById(ctx context.Context, movieId int) (domain.Movie, error) {
	responseFindId, err := service.movieRepository.FindById(ctx, movieId)
	if err != nil {
		return responseFindId, err
	}

	return responseFindId, nil
}
