package p

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	"os"
)

var apixuKey = os.Getenv("apixu_key")
const days = "7"
const apiURL = "http://api.apixu.com"

// forecastStruct : http://api.apixu.com/v1/forecast.json struct
// Autogenerated from https://mholt.github.io/json-to-go/
type forecastStruct struct {
	Location struct {
		Name           string  `json:"name"`
		Region         string  `json:"region"`
		Country        string  `json:"country"`
		Lat            float64 `json:"lat"`
		Lon            float64 `json:"lon"`
		TzID           string  `json:"tz_id"`
		LocaltimeEpoch int     `json:"localtime_epoch"`
		Localtime      string  `json:"localtime"`
	} `json:"location"`
	Current struct {
		LastUpdatedEpoch int     `json:"last_updated_epoch"`
		LastUpdated      string  `json:"last_updated"`
		TempC            float64 `json:"temp_c"`
		TempF            float64 `json:"temp_f"`
		IsDay            int     `json:"is_day"`
		Condition        struct {
			Text string `json:"text"`
			Icon string `json:"icon"`
			Code int    `json:"code"`
		} `json:"condition"`
		WindMph    float64 `json:"wind_mph"`
		WindKph    float64 `json:"wind_kph"`
		WindDegree int     `json:"wind_degree"`
		WindDir    string  `json:"wind_dir"`
		PressureMb float64 `json:"pressure_mb"`
		PressureIn float64 `json:"pressure_in"`
		PrecipMm   float64 `json:"precip_mm"`
		PrecipIn   float64 `json:"precip_in"`
		Humidity   int     `json:"humidity"`
		Cloud      int     `json:"cloud"`
		FeelslikeC float64 `json:"feelslike_c"`
		FeelslikeF float64 `json:"feelslike_f"`
		VisKm      float64 `json:"vis_km"`
		VisMiles   float64 `json:"vis_miles"`
		Uv         float64 `json:"uv"`
		GustMph    float64 `json:"gust_mph"`
		GustKph    float64 `json:"gust_kph"`
	} `json:"current"`
	Forecast struct {
		Forecastday []struct {
			Date      string `json:"date"`
			DateEpoch int    `json:"date_epoch"`
			Day       struct {
				MaxtempC      float64 `json:"maxtemp_c"`
				MaxtempF      float64 `json:"maxtemp_f"`
				MintempC      float64 `json:"mintemp_c"`
				MintempF      float64 `json:"mintemp_f"`
				AvgtempC      float64 `json:"avgtemp_c"`
				AvgtempF      float64 `json:"avgtemp_f"`
				MaxwindMph    float64 `json:"maxwind_mph"`
				MaxwindKph    float64 `json:"maxwind_kph"`
				TotalprecipMm float64 `json:"totalprecip_mm"`
				TotalprecipIn float64 `json:"totalprecip_in"`
				AvgvisKm      float64 `json:"avgvis_km"`
				AvgvisMiles   float64 `json:"avgvis_miles"`
				Avghumidity   float64 `json:"avghumidity"`
				Condition     struct {
					Text string `json:"text"`
					Icon string `json:"icon"`
					Code int    `json:"code"`
				} `json:"condition"`
				Uv float64 `json:"uv"`
			} `json:"day"`
			Astro struct {
				Sunrise  string `json:"sunrise"`
				Sunset   string `json:"sunset"`
				Moonrise string `json:"moonrise"`
				Moonset  string `json:"moonset"`
			} `json:"astro"`
		} `json:"forecastday"`
	} `json:"forecast"`
}

// conditionList : map apixu weather conditions to Yahoo icon
// check https://erikflowers.github.io/weather-icons/api-list.html 
var conditionList = map[string]string {
	"Sunny":"32",
	"Clear":"31",
	"Partly cloudy":"30",
	"Cloudy":"26",
	"Overcast":"26",
	"Mist":"20",
	"Patchy rain possible":"11",
	"Patchy snow possible":"13",
	"Patchy sleet possible":"12",
	"Patchy freezing drizzle possible":"25",
	"Thundery outbreaks possible":"3",
	"Blowing snow":"5",
	"Blizzard":"43",
	"Fog":"20",
	"Freezing fog":"20",
	"Patchy light drizzle":"11",
	"Light drizzle":"11",
	"Freezing drizzle":"11",
	"Heavy freezing drizzle":"11",
	"Patchy light rain":"11",
	"Light rain":"11",
	"Moderate rain at times":"11",
	"Moderate rain":"11",
	"Heavy rain at times":"17",
	"Heavy rain":"17",
	"Light freezing rain":"11",
	"Moderate or heavy freezing rain":"17",
	"Light sleet":"11",
	"Moderate or heavy sleet":"17",
	"Patchy light snow":"13",
	"Light snow":"13",
	"Patchy moderate snow":"18",
	"Moderate snow":"18",
	"Patchy heavy snow":"18",
	"Heavy snow":"18",
	"Ice pellets":"25",
	"Light rain shower":"13",
	"Moderate or heavy rain shower":"18",
	"Torrential rain shower":"18",
	"Light sleet showers":"13",
	"Moderate or heavy sleet showers":"18",
	"Light snow showers":"13",
	"Moderate or heavy snow showers":"18",
	"Light showers of ice pellets":"18",
	"Moderate or heavy showers of ice pellets":"18",
	"Patchy light rain with thunder":"37",
	"Moderate or heavy rain with thunder":"3",
	"Patchy light snow with thunder":"39",
	"Moderate or heavy snow with thunder":"3",
}

// weatherConditionsYahooMap : convert from apixu weather conditions to Yahoo icon
// check https://erikflowers.github.io/weather-icons/api-list.html 
func weatherConditionsYahooMap(condition string) string {
	api := "yahoo"

	if value, present := conditionList[condition]; present {
      		return (fmt.Sprintf("%s-%s", api, value))
   	}
	return (fmt.Sprintf("%s-32", api)) // Other condition is sunny :)
}

// forecastAPIV1 : forecast struct convert 
func forecastAPIV1(forecast forecastStruct) ForecastAPIV1Struct {
    var res ForecastAPIV1Struct
    lastUpdated, _ := time.Parse("2006-01-02 15:04", forecast.Current.LastUpdated)
    res.Location.Name = forecast.Location.Name
    res.Current.LastUpdated = lastUpdated
    res.Current.TempC = forecast.Current.TempC
	res.Current.Condition.Text = forecast.Current.Condition.Text
	res.Current.Condition.Icon = weatherConditionsYahooMap(forecast.Current.Condition.Text)
	res.Current.PrecipMm = forecast.Current.PrecipMm
	res.Current.WindKph = forecast.Current.WindKph
    res.Current.WindDir = forecast.Current.WindDir
	
    for k, v := range forecast.Forecast.Forecastday {
		if (k==0){continue}
		forecastdayDate, _ := time.Parse("2006-01-02", v.Date)
		var forecastday ForecastdayAPIV1Struct
        forecastday.Date = forecastdayDate
		forecastday.Day.Condition.Text = v.Day.Condition.Text
		forecastday.Day.Condition.Icon = weatherConditionsYahooMap(v.Day.Condition.Text)
        forecastday.Day.AvgtempC = v.Day.AvgtempC
        forecastday.Day.MaxtempC = v.Day.MaxtempC
        forecastday.Day.MintempC = v.Day.MintempC
        
        res.Forecast.Forecastday = append (res.Forecast.Forecastday, forecastday)
    }
    
    return res
}

// GetForecastAPIV1 : Get Forecast from apixu
func GetForecastAPIV1 (query string) string {
	
	if len(apixuKey) == 0 {
		log.Fatal("apixu_key variable not defined")
	}
	
	url := fmt.Sprintf("%s/v1/forecast.json?key=%s&q=%s&days=%s", apiURL, apixuKey, query, days)

	spaceClient := http.Client{
		Timeout: time.Second * 2, // Maximum of 2 secs
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "forecast-test")

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	forecast := forecastStruct{}
	jsonErr := json.Unmarshal(body, &forecast)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	resJSON, err := json.Marshal(forecastAPIV1(forecast))
	if err != nil {
		log.Fatal(err)
	}
	return (string(resJSON))
}

// GetIndexAPIV1 : Get index.html
func GetIndexAPIV1 () string {
	return ("templates/indexV1.html")
}