package urls

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"time"
)

// Data структура для хранения информации о короткой и оригинальной ссылке
type Data struct {
	Short        string    `json:"shortUrl"`
	Destination  string    `json:"destination"`
	CreationTime time.Time `json:"-"`
}

// Storage интерфейс для работы с базой данных
type Storage interface {
	Save(urlData Data) error
	Load(shortURL string) (Data, error)
}

// Service структура, хранящая ссылки и базу данных
type Service struct {
	storage Storage
}

// New создает новый URLService
func New(storage Storage) *Service {
	return &Service{storage: storage}
}

// Create создает новую короткую ссылку
func (service *Service) Create(destination string) (Data, error) {
	if destination == "" {
		return Data{}, errors.New("invalid destination url")
	}

	shortURL, err := generateUrl()
	if err != nil {
		return Data{}, err
	}

	urlData := Data{
		Short:        shortURL,
		Destination:  destination,
		CreationTime: time.Now(),
	}

	err = service.storage.Save(urlData)
	if err != nil {
		return Data{}, err
	}

	return urlData, nil
}

// Get возвращает оригинальную ссылку на основе короткой
func (service *Service) Get(short string) (Data, error) {
	return service.storage.Load(short)
}

func generateUrl() (string, error) {
	const len = 6
	randBytes := make([]byte, len)
	_, err := rand.Read(randBytes)
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(randBytes)[:len], nil
}
