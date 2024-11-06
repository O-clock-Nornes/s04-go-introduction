package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Coordinates struct {
	Lat float32
	Lon float32
}

func getCoordFromCity(city string) (Coordinates, error) {
	url := fmt.Sprintf("http://api.openweathermap.org/geo/1.0/direct?q=%s&appid=%s", city, token)
	resp, err := http.Get(url)
	if err != nil {
		return Coordinates{}, err
	}
	body, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return Coordinates{}, err
	}
	var resultats []Coordinates
	json.Unmarshal(body, &resultats)
	return resultats[0], nil
}
