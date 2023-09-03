package postgres

import (
	"database/sql"
	"go-course/homework-16/pkg/db"
	"os"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

var mock sqlmock.Sqlmock
var conn *sql.DB

func TestMain(m *testing.M) {
	conn, mock, _ = NewMock()
	defer conn.Close()

	code := m.Run()

	os.Exit(code)
}

func NewMock() (*sql.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	return db, mock, err
}

func TestAddMovies(t *testing.T) {
	pg := &Postgres{conn: conn}

	movies := []db.Movie{
		{
			ID:          1,
			Title:       "Test Movie",
			ReleaseYear: 2020,
			StudioID:    1,
			BoxOffice:   500000,
			Rating:      "PG-13",
		},
	}

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO movies").WithArgs(movies[0].Title, movies[0].ReleaseYear, movies[0].StudioID, movies[0].BoxOffice, movies[0].Rating).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err := pg.AddMovies(movies)
	assert.NoError(t, err)
}
func TestDeleteMovie(t *testing.T) {
	pg := &Postgres{conn: conn}

	movieID := 1
	mock.ExpectBegin()
	mock.ExpectExec("DELETE FROM movies").WithArgs(movieID).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err := pg.DeleteMovie(movieID)
	assert.NoError(t, err)
}

func TestUpdateMovie(t *testing.T) {
	pg := &Postgres{conn: conn}

	movie := db.Movie{
		ID:          1,
		Title:       "Updated Movie",
		ReleaseYear: 2021,
		StudioID:    1,
		BoxOffice:   600000,
		Rating:      "PG-13",
	}

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE movies").WithArgs(movie.Title, movie.ReleaseYear, movie.StudioID, movie.BoxOffice, movie.Rating, movie.ID).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err := pg.UpdateMovie(movie)
	assert.NoError(t, err)
}

func TestMovies(t *testing.T) {
	pg := &Postgres{conn: conn}

	rows := sqlmock.NewRows([]string{"id", "title", "release_year", "studio_id", "box_office", "rating"}).
		AddRow(1, "Test Movie", 2020, 1, 500000, 8.0)

	mock.ExpectQuery("SELECT id, title, release_year, studio_id, box_office, rating FROM movies").WillReturnRows(rows)

	movies, err := pg.Movies(0)
	assert.NoError(t, err)
	assert.Len(t, movies, 1)
	assert.Equal(t, "Test Movie", movies[0].Title)
}
