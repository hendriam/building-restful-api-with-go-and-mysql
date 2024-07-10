package service

import (
	"context"
	"fmt"

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

func (service *MovieService) CreateMovie(ctx context.Context, movieRequest web.MovieCreateRequestBody) (web.MovieCreateResponseBody, error) {
	movie := domain.Movie{
		Title: movieRequest.Title,
		Year:  movieRequest.Year,
	}

	responseCreateMovie, err := service.movieRepository.Insert(ctx, movie)
	if err != nil {
		fmt.Println("[SERVICE] error =>", err)
		return web.MovieCreateResponseBody{}, err
	}

	return web.MovieCreateResponseBody{Code: 201, Message: "Oke", Data: responseCreateMovie}, nil
}
