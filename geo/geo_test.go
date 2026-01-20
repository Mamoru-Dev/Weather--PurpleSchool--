package geo_test

import (
	"pet/weather_cli/geo"
	"testing"
)

func TestGetMyLocation(t *testing.T) {
	// Arrange - подготовка, expected результат, данные для функции
	city := "London"
	expected := geo.GeoData{
		City: "London",
	}

	// Act - выполнение функции с заданными данными
	got, err := geo.GetMyLocation(city)

	// Assert - проверка результата с expected
	if err != nil {
		t.Error("Ошибка получения города")
	}
	if got.City != expected.City {
		t.Errorf("Ожидалось %v, получено %v", expected.City, got.City)
	}
}

func TestGetMyLocationNoCity(t *testing.T) {
	city := "Londonfsfsd"
	_, err := geo.GetMyLocation(city)
	if err != geo.ErrNoCity {
		t.Errorf("Ожидалось %v, получено %v", geo.ErrNoCity, err)
	}
}
