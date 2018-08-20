package Service

import (
	"../Model"
	"../Repositories"
)

// MovieService handles some of the CRUID operations of the movie datamodel.
// It depends on a movie repository for its actions.
// It's here to decouple the data source from the higher level compoments.
// As a result a different repository type can be used with the same logic without any aditional changes.
// It's an interface and it's used as interface everywhere
// because we may need to change or try an experimental different domain logic at the future.
type MovieService interface {
	GetAll() []Model.HomeModel
	GetByID(id int64) (Model.HomeModel, bool)
	DeleteByID(id int64) bool
	UpdatePosterAndGenreByID(id int64, poster string, genre string) (Model.HomeModel, error)
}

// NewMovieService returns the default movie service.
func NewMovieService(repo Repositories.MovieRepository) MovieService {
	return &movieService{
		repo: repo,
	}
}

type movieService struct {
	repo Repositories.MovieRepository
}

// GetAll returns all movies.
func (s *movieService) GetAll() []Model.HomeModel {
	return s.repo.SelectMany(func(_ Model.HomeModel) bool {
		return true
	}, -1)
}

// GetByID returns a movie based on its id.
func (s *movieService) GetByID(id int64) (Model.HomeModel, bool) {
	return s.repo.Select(func(m Model.HomeModel) bool {
		return m.ID == id
	})
}

// UpdatePosterAndGenreByID updates a movie's poster and genre.
func (s *movieService) UpdatePosterAndGenreByID(id int64, poster string, genre string) (Model.HomeModel, error) {
	// update the movie and return it.
	return s.repo.InsertOrUpdate(Model.HomeModel{
		ID: id,
		// Poster: poster,
		// Genre:  genre,
	})
}

// DeleteByID deletes a movie by its id.
//
// Returns true if deleted otherwise false.
func (s *movieService) DeleteByID(id int64) bool {
	return s.repo.Delete(func(m Model.HomeModel) bool {
		return m.ID == id
	}, 1)
}
