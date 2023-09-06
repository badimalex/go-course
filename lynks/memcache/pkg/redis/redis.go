package redis

import (
	"context"
	"time"

	"github.com/badimalex/go-course/lynks/shortener/pkg/urls"

	"github.com/go-redis/redis/v8"
)

const expiryDuration = 24 * time.Hour

type RedisStorage struct {
	client *redis.Client
}

func NewRedisStorage(addr string, password string, db int) (*RedisStorage, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}

	return &RedisStorage{
		client: rdb,
	}, nil
}

func (storage *RedisStorage) Save(urlData urls.Data) error {
	err := storage.client.Set(context.Background(), urlData.Short, urlData.Destination, expiryDuration).Err()
	if err != nil {
		return err
	}

	// сохраните другие данные urlData, если необходимо

	return nil
}

func (storage *RedisStorage) Load(shortURL string) (urls.Data, error) {
	destURL, err := storage.client.Get(context.Background(), shortURL).Result()
	if err != nil {
		return urls.Data{}, err
	}

	// Здесь нужно будет получить другие данные, если они сохраняются

	return urls.Data{
		Short:       shortURL,
		Destination: destURL,
		// CreationTime: ..., // установите это, если сохраняете время создания
	}, nil
}
