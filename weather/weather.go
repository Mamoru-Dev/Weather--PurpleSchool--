package weather

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"pet/weather_cli/geo"
)

func GetWeather(geo geo.GeoData, format int) string {
	// Используем библиотеку url для формирования запроса к API
	baseUrl, err := url.Parse("https://wttr.in/" + geo.City)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	// Добавляем к нашему baseUrl query-параметр
	params := url.Values{}
	params.Add("format", fmt.Sprint(format)) // <-- Добавляем параметр format
	baseUrl.RawQuery = params.Encode()

	// Делаем запрос к API wttr.in
	resp, err := http.Get(baseUrl.String())
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	return string(body)
}
