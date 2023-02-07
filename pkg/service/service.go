package service

import (
	"encoding/json"
	"github.com/zhayt/groupie-tracker/pkg/models"
	"io/ioutil"
	"net/http"
)

func GetAll(addr string) ([]models.Artist, error) {
	text, err := http.Get(addr)
	if err != nil {
		return nil, err
	}

	artists := make([]models.Artist, 1)
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

func GetById(addr, id string) (*models.Artist, error) {
	text, err := http.Get(addr + "/" + id)
	if err != nil {
		return nil, err
	}
	defer text.Body.Close()

	artist := &models.Artist{}

	content, err := ioutil.ReadAll(text.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(content, &artist)

	if err != nil {
		return nil, err
	}

	return artist, nil
}
