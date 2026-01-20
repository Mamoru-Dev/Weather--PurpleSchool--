package main

import (
	"flag"
	"fmt"
	"pet/weather_cli/geo"
	"pet/weather_cli/weather"
)

func main() {
	city := flag.String("city", "", "Город пользователя")
	format := flag.Int("format", 1, "Формат вывода погоды")
	flag.Parse()
	geoData, err := geo.GetMyLocation(*city)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	weatherData := weather.GetWeather(*geoData, *format)
	fmt.Println(weatherData)
}
