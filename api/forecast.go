// Package handler contains an HTTP Cloud Function.
package handler

import (
	"fmt"
	"net/http"
	"net/url"

	p "github.com/rfiestas/forecast"
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
	location = getQueryKey(r, "location", "Barcelona")
	fmt.Fprint(w, p.GetForecastAPIV1(url.QueryEscape(location)))
	return
}
