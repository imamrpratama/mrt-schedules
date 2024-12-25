package station

import (
	"encoding/json"
	"net/http"
	"time"
	"your_project_path/client" // Update the import path
)

type Service interface {
	GetAllStation() (response []StationResponse, err error)
}

type service struct {
	client *http.Client
}

func NewService() Service {
	return &service{
		client: &http.Client{
			Timeout: time.Second * 10, // Timeout after 10 seconds
		},
	}
}

func (s *service) GetAllStation() (response []StationResponse, err error) {
	url := "https://jakartamrt.co.id/id/val/stasiuns"

	// Hit URL
	byteResponse, err := client.DoRequest(s.client, url)
	if err != nil {
		return nil, err
	}

	var stations []Station
	err = json.Unmarshal(byteResponse, &stations)
	if err != nil {
		return nil, err
	}

	// Map response
	for _, item := range stations {
		response = append(response, StationResponse{
			Id:   item.Id,
			Name: item.Name,
		})
	}

	return response, nil
}
