package cache

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Service struct {
	apiUrl string
}

func New(apiUrl string) *Service {
	return &Service{
		apiUrl: apiUrl,
	}
}

type createRequest struct {
	Destination string `json:"destination"`
	Short       string `json:"short"`
}

type getResponse struct {
	Destination string `json:"destination"`
}

func (c *Service) Create(dest, short string) error {
	data, err := json.Marshal(createRequest{Destination: dest, Short: short})

	if err != nil {
		return err
	}

	resp, err := http.Post(c.apiUrl, "application/json", bytes.NewBuffer(data))

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil
}

func (c *Service) Get(short string) (*getResponse, error) {
	resp, err := http.Get(c.apiUrl + "/" + short)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var res getResponse
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
