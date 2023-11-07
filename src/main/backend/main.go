package backend

import (
	"fmt"
)

func printNumbers1() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
	}
}

func printLetters1() {
	for i := 'A'; i < 'A'+10; i++ {
		fmt.Printf("%c ", i)
	}
}

func print1() {
	printNumbers1()
	printLetters1()
}

func goPrint1() {
	go printNumbers1()
	go printLetters1()
}

func Add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println("I'm working")
}

//func getLatLong(city string) (*LatLong, error) {
//    endpoint := fmt.Sprintf("https://geocoding-api.open-meteo.com/v1/search?name=%s&count=1&language=en&format=json", url.QueryEscape(city))
//    resp, err := http.Get(endpoint)
//    if err != nil {
//   	 return nil, fmt.Errorf("error making request to Geo API: %w", err)
//    }
//    defer resp.Body.Close()
//
//    var response GeoResponse
//    if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
//   	 return nil, fmt.Errorf("error decoding response: %w", err)
//    }
//
//    if len(response.Results) < 1 {
//   	 return nil, errors.New("no results found")
//    }
//
//    return &response.Results[0], nil
//}
//
//type GeoResponse struct {
//    // A list of results; we only need the first one
//    Results []LatLong `json:"results"`
//}
//
//type LatLong struct {
//    Latitude  float64 `json:"latitude"`
//    Longitude float64 `json:"longitude"`
//}
//
//func getWeather(latLong LatLong) (string, error) {
//    endpoint := fmt.Sprintf("https://api.open-meteo.com/v1/forecast?latitude=%.6f&longitude=%.6f&hourly=temperature_2m", latLong.Latitude, latLong.Longitude)
//    resp, err := http.Get(endpoint)
//    if err != nil {
//   	 return "", fmt.Errorf("error making request to Weather API: %w", err)
//    }
//    defer resp.Body.Close()
//
//    body, err := io.ReadAll(resp.Body)
//    if err != nil {
//   	 return "", fmt.Errorf("error reading response body: %w", err)
//    }
//
//    return string(body), nil
//}
//
//func main() {
//    latlong, err := getLatLong("London") // you know it will rain
//    if err != nil {
//   	 log.Fatalf("Failed to get latitude and longitude: %s", err)
//    }
//    fmt.Printf("Latitude: %f, Longitude: %f\n", latlong.Latitude, latlong.Longitude)
//
//    weather, err := getWeather(*latlong)
//    if err != nil {
//   	 log.Fatalf("Failed to get weather: %s", err)
//    }
//    fmt.Printf("Weather: %s\n", weather)
//}
