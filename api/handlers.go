package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"

	f "github.com/rfiestas/forecast"
)

// ForecastAPIV1 : Forecast call
func ForecastAPIV1(w http.ResponseWriter, r *http.Request) {
	var location string
	location = f.GetQueryKey(r, "location", "Barcelona")
	fmt.Fprint(w, f.GetForecastAPIV1(url.QueryEscape(location)))
	return
}

// GetIndex : Return index.html
func GetIndex(w http.ResponseWriter, r *http.Request) {
	var location string
	location = f.GetQueryKey(r, "location", "Barcelona")
	data, err := ioutil.ReadFile(f.GetIndexAPIV1())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	w.Header().Set("Cache-Control", "max-age=86400, public")
	fmt.Fprint(w, strings.Replace(string(data), "#LOCATION#", location, -1))
	return
}

// GetRobots : Return robots.txt
func GetRobots(w http.ResponseWriter, r *http.Request) {
	var data string
	data = "User-agent: *\nDisallow: \n"
	fmt.Fprint(w, data)
	return
}
