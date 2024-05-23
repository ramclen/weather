package main

import (
    "encoding/json"
    "math/rand"
    "net/http"
    "time"
)

type WeatherData struct {
    Cod      int    `json:"cod"`
    CityID   int    `json:"city_id"`
    CalcTime float64 `json:"calctime"`
    Result   struct {
        Month         int `json:"month"`
        Temperature   Statistics `json:"temp"`
        Pressure      Statistics `json:"pressure"`
        Humidity      Statistics `json:"humidity"`
        Wind          Statistics `json:"wind"`
        Precipitation Statistics `json:"precipitation"`
        Clouds        Statistics `json:"clouds"`
        SunshineHours float64    `json:"sunshine_hours"`
    } `json:"result"`
}

type Statistics struct {
    Min       float64 `json:"min"`
    Max       float64 `json:"max"`
    Median    float64 `json:"median"`
    Mean      float64 `json:"mean"`
    P25       float64 `json:"p25"`
    P75       float64 `json:"p75"`
    StDev     float64 `json:"st_dev"`
    Num       int     `json:"num"`
}

func randomStats(min, max float64) Statistics {
    mean := min + rand.Float64()*(max-min)
    return Statistics{
        Min:    min,
        Max:    max,
        Median: min + rand.Float64()*(max-min),
        Mean:   mean,
        P25:    mean - mean*0.1,
        P75:    mean + mean*0.1,
        StDev:  rand.Float64() * 5,
        Num:    rand.Intn(1000) + 3000,
    }
}

func randomWeatherStats() WeatherData {
	rand.Seed(time.Now().UnixNano())
    return WeatherData{
        Cod:      200,
        CityID:   5400075,
        CalcTime: rand.Float64(),
        Result: struct {
            Month         int `json:"month"`
            Temperature   Statistics `json:"temp"`
            Pressure      Statistics `json:"pressure"`
            Humidity      Statistics `json:"humidity"`
            Wind          Statistics `json:"wind"`
            Precipitation Statistics `json:"precipitation"`
            Clouds        Statistics `json:"clouds"`
            SunshineHours float64    `json:"sunshine_hours"`
        }{
            Month:         rand.Intn(12) + 1,
            Temperature:   randomStats(260, 300),
            Pressure:      randomStats(980, 1040),
            Humidity:      randomStats(10, 100),
            Wind:          randomStats(0, 20),
            Precipitation: randomStats(0, 5),
            Clouds:        randomStats(0, 100),
            SunshineHours: rand.Float64() * 120,
        },
    }
}

func weatherHandler(w http.ResponseWriter, r *http.Request) {
    data := randomWeatherStats()

	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

    // Handle preflight requests for CORS
    if r.Method == "OPTIONS" {
        w.WriteHeader(http.StatusOK)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(data)
}

func main() {
    http.HandleFunc("/weather", weatherHandler)
    http.ListenAndServe(":8080", nil)
}