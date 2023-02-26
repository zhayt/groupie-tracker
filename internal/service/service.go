package service

import (
	"encoding/json"
	"errors"
	"github.com/zhayt/groupie-tracker/internal/models"
	"io"
	"net/http"
)

func GetAll(addr string) ([]*models.Artist, error) {
	text, err := http.Get(addr)
	if err != nil {
		return nil, err
	}

	defer func() { _ = text.Body.Close() }()

	artists := make([]*models.Artist, 0)

	content, err := io.ReadAll(text.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(content, &artists)
	if err != nil {
		return nil, err
	}

	for _, artist := range artists {
		artist.Locations, err = GetLocations(artist.LocationsURL)
		if err != nil {
			return nil, err
		}
	}

	return artists, nil
}

func GetById(addr, id string) (*models.PresentAllData, error) {
	allData := models.New()

	err := unmarshal(addr+"/"+id, allData.Artist)
	if err != nil {
		return nil, err
	}

	if allData.Artist.Id == 0 {
		return nil, errors.New("not found")
	}

	err = unmarshal(allData.Artist.LocationsURL, allData.Locations)
	if err != nil {
		return nil, err
	}

	err = unmarshal(allData.Artist.ConcertDates, allData.ConcertDates)
	if err != nil {
		return nil, err
	}

	err = unmarshal(allData.Artist.Relations, allData.Relations)
	if err != nil {
		return nil, err
	}

	return allData, nil
}

func GetLocations(url string) (*models.Locations, error) {
	location := &models.Locations{}
	err := unmarshal(url, location)
	if err != nil {
		return nil, err
	}
	return location, nil
}

func unmarshal(url string, data interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer func() { _ = resp.Body.Close() }()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, data)
	if err != nil {
		return err
	}

	return err
}
