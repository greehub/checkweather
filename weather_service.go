package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func getWeather(lat, lon string) (WeatherResponse, error) {
	var weatherResponse WeatherResponse
	apiKey := os.Getenv("OPENWEATHER_API_KEY")
	url := fmt.Sprintf(apiUrl+"?lat=%s&lon=%s&units=metric&appid=%s", lat, lon, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		return weatherResponse, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return weatherResponse, fmt.Errorf(MsgApiError+": %s", resp.Status)
	}

	err = json.NewDecoder(resp.Body).Decode(&weatherResponse)
	if err != nil {
		return weatherResponse, err
	}

	return weatherResponse, nil
}
