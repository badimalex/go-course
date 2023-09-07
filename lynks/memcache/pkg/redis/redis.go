package redis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

const expiryDuration = 24 * time.Hour

type Data struct {
	Short        string    `json:"shortUrl"`
	Destination  string    `json:"destination"`
	CreationTime time.Time `json:"-"`
}

type Storage struct {
	client *redis.Client
}

func New(addr string, password string, db int) (*Storage, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}

	return &Storage{
		client: rdb,
	}, nil
}

func (storage *Storage) Save(urlData Data) error {
	err := storage.client.Set(context.Background(), urlData.Short, urlData.Destination, expiryDuration).Err()
	if err != nil {
		return err
	}

	return nil
}

func (storage *Storage) Load(shortURL string) (Data, error) {
	destURL, err := storage.client.Get(context.Background(), shortURL).Result()
	if err != nil {
		return Data{}, err
	}

	return Data{
		Short:       shortURL,
		Destination: destURL,
	}, nil
}
