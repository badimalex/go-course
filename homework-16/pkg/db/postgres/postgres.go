package postgres

import (
	"database/sql"
	"go-course/homework-16/pkg/db"

	_ "github.com/lib/pq"
)

type Postgres struct {
	conn *sql.DB
}

func New(conn string) (*Postgres, error) {
	db, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, err
	}
	return &Postgres{conn: db}, nil
}

func (pg *Postgres) AddMovies(movies []db.Movie) error {
	tx, err := pg.conn.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	for _, movie := range movies {
		_, err := tx.Exec(`
            INSERT INTO movies (title, release_year, studio_id, box_office, rating)
            VALUES ($1, $2, $3, $4, $5)
        `, movie.Title, movie.ReleaseYear, movie.StudioID, movie.BoxOffice, movie.Rating)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

func (pg *Postgres) DeleteMovie(movieID int) error {
	_, err := pg.conn.Exec("DELETE FROM movies WHERE id = $1", movieID)
	return err
}

func (pg *Postgres) UpdateMovie(movie db.Movie) error {
	_, err := pg.conn.Exec(`
        UPDATE movies SET title = $1, release_year = $2, studio_id = $3, box_office = $4, rating = $5 WHERE id = $6
    `, movie.Title, movie.ReleaseYear, movie.StudioID, movie.BoxOffice, movie.Rating, movie.ID)
	return err
}

func (pg *Postgres) GetMovies(studioID int) ([]db.Movie, error) {
	var rows *sql.Rows
	var err error

	if studioID == 0 {
		rows, err = pg.conn.Query("SELECT id, title, release_year, studio_id, box_office, rating FROM movies")
	} else {
		rows, err = pg.conn.Query("SELECT id, title, release_year, studio_id, box_office, rating FROM movies WHERE studio_id = $1", studioID)
	}

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movies []db.Movie
	for rows.Next() {
		var m db.Movie
		if err := rows.Scan(&m.ID, &m.Title, &m.ReleaseYear, &m.StudioID, &m.BoxOffice, &m.Rating); err != nil {
			return nil, err
		}
		movies = append(movies, m)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return movies, nil
}
