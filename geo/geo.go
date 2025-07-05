package geo

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type GeoData struct {
	City string `json:"city"`
}
type CityPopulation struct {
	Error bool `json:"error"`
}

func GetMyLocation(city string) (*GeoData, error) {
	if city != "" {
		if !checkCity(city) {
			panic("нет такого города")
		}
		return &GeoData{City: city}, nil
	}
	resp, err := http.Get("https://ipapi.co/json/")

	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {

		return nil, errors.New("not200")
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var geo GeoData
	err = json.Unmarshal(body, &geo)
	if err != nil {
		return nil, err
	}
	return &geo, nil
}

//https://countriesnow.space/api/v0.1/countries/population/cities

func checkCity(city string) bool {
	postBody, err := json.Marshal(map[string]string{
		"city": city,
	})
	if err != nil {
		return false
	}
	resp, err := http.Post("https://countriesnow.space/api/v0.1/countries/population/cities", "application/json", bytes.NewBuffer(postBody))
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false
	}
	var populationResponce CityPopulation
	json.Unmarshal(body, &populationResponce)
	return !populationResponce.Error
}
