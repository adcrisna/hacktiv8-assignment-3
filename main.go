package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"text/template"
)

const (
	min = 1
	max = 100
)

type WeatherData struct {
	Water  int    `json:"water"`
	Wind   int    `json:"wind"`
	Status string `json:"status"`
}

func main() {
	http.HandleFunc("/weather", Weather)
	fmt.Println("please open http://localhost:3000/weather")
	http.ListenAndServe(":3000", nil)
}

func Weather(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		file, _ := os.ReadFile("data.json")
		hasil := WeatherData{}
		_ = json.Unmarshal(file, &hasil)

		water := rand.Intn(max-min) + min
		wind := rand.Intn(max-min) + min
		weatherStatus := "Status Default"

		if water < 5 {
			weatherStatus = "AMAN"
		} else if water >= 6 && water == 8 {
			weatherStatus = "SIAGA"
		} else if water > 8 {
			weatherStatus = "BAHAYA"
		} else if wind < 6 {
			weatherStatus = "AMAN"
		} else if wind >= 7 && wind == 15 {
			weatherStatus = "SIAGA"
		} else if wind > 15 {
			weatherStatus = "BAHAYA"
		}

		updatedData := WeatherData{
			Water:  water,
			Wind:   wind,
			Status: weatherStatus,
		}

		json, _ := json.Marshal(&updatedData)
		_ = ioutil.WriteFile("data.json", json, os.ModePerm)

		tmpl, err := template.ParseFiles("view.html")

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, updatedData)
		return

	}

	http.Error(w, "Invalid Method", http.StatusBadRequest)
}
