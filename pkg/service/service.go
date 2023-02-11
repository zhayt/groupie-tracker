package service

import (
	"encoding/json"
	"errors"
	"github.com/zhayt/groupie-tracker/pkg/models"
	"io/ioutil"
	"net/http"
)

func GetAll(addr string) ([]*models.Artist, error) {
	text, err := http.Get(addr)
	if err != nil {
		return nil, err
	}

	artists := make([]*models.Artist, 1)
	defer text.Body.Close()

	content, err := ioutil.ReadAll(text.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(content, &artists)
	if err != nil {
		return nil, err
	}

	return artists, nil
}

func GetById(addr, id string) (*models.PresentAllData, error) {
	artist := &models.Artist{}
	location := &models.Locations{}
	concertDates := &models.ConcertDates{}
	relations := &models.Relations{}

	err := unmarshal(addr+"/"+id, artist)
	if err != nil {
		return nil, err
	}
	if artist.Id == 0 {
		return nil, errors.New("not found")
	}
	err = unmarshal(artist.Locations, location)
	if err != nil {
		return nil, err
	}

	err = unmarshal(artist.ConcertDates, concertDates)
	if err != nil {
		return nil, err
	}

	err = unmarshal(artist.Relations, relations)
	if err != nil {
		return nil, err
	}

	return &models.PresentAllData{
		Artist:       artist,
		Locations:    location,
		ConcertDates: concertDates,
		Relations:    relations,
	}, nil
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
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, data)
	if err != nil {
		return err
	}

	return err
}
