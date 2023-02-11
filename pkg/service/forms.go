package service

import (
	"errors"
	"fmt"
	"github.com/zhayt/groupie-tracker/pkg/models"
	"net/url"
	"strconv"
	"strings"
)

type Form struct {
	url.Values
	AllArtists     []*models.Artist
	SearchedArtist []*models.Artist
}

func NewArtists(data url.Values) *Form {
	return &Form{
		data,
		nil,
		nil,
	}
}

func (f *Form) Required(fields ...string) error {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			fmt.Println("imaio")
			return errors.New("this field cannot be blank")
		}
	}
	return nil
}

func (f *Form) ParseValue(str string) (string, error) {
	slice := strings.Split(strings.TrimSpace(str), " - ")
	if len(slice) > 2 {
		return "", errors.New("bad request")
	}
	return strings.ToLower(strings.TrimSpace(slice[0])), nil
}

func (f *Form) Find(toSearch string) {
	var searchedArtists []*models.Artist
	for _, artist := range f.AllArtists {
		if strings.Contains(strings.ToLower(artist.Name), toSearch) || strings.Contains(strings.ToLower(strconv.Itoa(artist.CreationDate)), toSearch) || strings.Contains(strings.ToLower(artist.FirstAlbum), toSearch) {
			searchedArtists = append(searchedArtists, artist)
			continue
		}

		for _, member := range artist.Members {
			if strings.Contains(strings.ToLower(member), toSearch) {
				searchedArtists = append(searchedArtists, artist)
				continue
			}
		}

		locations, err := GetLocations(artist.Locations)
		if err != nil {
			continue
		}

		for _, loc := range locations.Locations {
			if strings.Contains(strings.ToLower(loc), toSearch) {
				searchedArtists = append(searchedArtists, artist)
				break
			}
		}
	}

	if len(searchedArtists) == 0 {
		return
	}
	f.SearchedArtist = searchedArtists
}
