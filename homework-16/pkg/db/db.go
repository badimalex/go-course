package db

type Movie struct {
	ID          int
	Title       string
	ReleaseYear int
	StudioID    int
	BoxOffice   float64
	Rating      string
}

type Interface interface {
	AddMovies(movies []Movie) error
	DeleteMovie(id int) error
	UpdateMovie(movie Movie) error
	Movies(studioID int) ([]Movie, error)
}
