package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/joho/godotenv"
)

func validateInput(query url.Values) (string, string, error) {
	lat := query.Get("lat")
	lon := query.Get("lon")
	if lat == "" || lon == "" {
		log.Print(MsgMissingParameters)
		return lat, lon, fmt.Errorf(MsgMissingParameters)
	}
	return lat, lon, nil
}

func weatherHandler(w http.ResponseWriter, r *http.Request) {
	lat, lon, err := validateInput(r.URL.Query())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	weatherResp, err := getWeather(lat, lon)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if weatherResp.Main.Temp == 0 && len(weatherResp.Weather) == 0 {
		log.Print(MsgResponseAttributes)
		return
	}

	condition := "moderate"
	if weatherResp.Main.Temp < 285 {
		condition = "cold"
	} else if weatherResp.Main.Temp > 297 {
		condition = "hot"
	}

	result := map[string]string{
		"condition":   condition,
		"temperature": fmt.Sprintf("%.2f°C", weatherResp.Main.Temp),
		"description": weatherResp.Weather[0].Description,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)

}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	http.HandleFunc("/checkweather", weatherHandler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Starting server on port %s...", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
