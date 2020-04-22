// Package handler contains an HTTP Cloud Function.
package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/rfiestas/forecast"
)

// getQueryKey : take a http request url query key, assign default value when not exist.
func getQueryKey(r *http.Request, key string, failure string) string {
	var value string
	keys, ok := r.URL.Query()[key]
	if !ok || len(keys[0]) < 1 {
		value = failure
	} else {
		value = string(keys[0])
	}
	return url.QueryEscape(value)
}

// ForecastAPIV1 : Forecast call
func ForecastAPIV1(w http.ResponseWriter, r *http.Request) {
	var location string
	location = forecast.GetQueryKey(r, "location", "Barcelona")
	fmt.Fprint(w, forecast.GetForecastAPIV1(url.QueryEscape(location)))
	return
}

// GetIndex : Return index.html
func GetIndex(w http.ResponseWriter, r *http.Request) {
	var location string
	location = forecast.GetQueryKey(r, "location", "Barcelona")
	data, err := ioutil.ReadFile(forecast.GetIndexAPIV1())
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
