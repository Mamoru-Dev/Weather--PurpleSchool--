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

type CityPopulationResponse struct {
	Error bool `json:"error"`
}

func GetMyLocation(city string) (*GeoData, error) {
	// Если город передан => Используем его
	if city != "" {
		if CheckCity(city) {
			return &GeoData{
				City: city,
			}, nil
		} else {
			panic("Такого города нет")
		}
	}

	// Если город не передан => Ищем его по IP
	res, err := http.Get("https://ipapi.co/json/")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, errors.New("NOT200")
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var geo GeoData
	json.Unmarshal(body, &geo)
	return &geo, nil
}

func CheckCity(city string) bool {
	// Создаем объект body для Post-запроса
	postBody, _ := json.Marshal(map[string]string{"city": city})

	// Post(url, кодировка, body)
	apiUrl := "https://countriesnow.space/api/v0.1/countries/population/cities"
	resp, err := http.Post(apiUrl, "application/json", bytes.NewBuffer(postBody))
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false
	}
	var populationResponse CityPopulationResponse
	json.Unmarshal(body, &populationResponse)
	return !populationResponse.Error
}
