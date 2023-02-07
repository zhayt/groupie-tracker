package service

import (
	"encoding/json"
	"github.com/zhayt/groupie-tracker/pkg/models"
	"io/ioutil"
	"net/http"
)

func GetAllArtists() ([]models.Artist, error) {
	text, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
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
