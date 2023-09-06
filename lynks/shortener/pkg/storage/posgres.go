package storage

import (
	"database/sql"
	"time"

	"github.com/badimalex/go-course/lynks/shortener/pkg/urls"

	_ "github.com/lib/pq"
)

type PostgresStorage struct {
	db *sql.DB
}

func New(dataSourceName string) (*PostgresStorage, error) {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStorage{db: db}, nil
}

func (storage *PostgresStorage) Save(urlData urls.Data) error {
	query := `INSERT INTO urls (short_url, destination, creation_time) VALUES ($1, $2, $3)`
	_, err := storage.db.Exec(query, urlData.Short, urlData.Destination, urlData.CreationTime)
	return err
}

func (storage *PostgresStorage) Load(shortURL string) (urls.Data, error) {
	query := `SELECT destination, creation_time FROM urls WHERE short_url = $1`
	row := storage.db.QueryRow(query, shortURL)

	var destination string
	var creationTime time.Time
	err := row.Scan(&destination, &creationTime)
	if err != nil {
		return urls.Data{}, err
	}

	return urls.Data{
		Short:        shortURL,
		Destination:  destination,
		CreationTime: creationTime,
	}, nil
}

// Init инициализирует базу данных, создавая необходимую таблицу
func (storage *PostgresStorage) Init() error {
	query := `
	CREATE TABLE IF NOT EXISTS urls (
		short_url VARCHAR(50) PRIMARY KEY,
		destination VARCHAR(255) NOT NULL,
		creation_time TIMESTAMP NOT NULL
	)`
	_, err := storage.db.Exec(query)
	return err
}
