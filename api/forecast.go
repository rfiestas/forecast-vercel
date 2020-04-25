// Package handler contains an HTTP Cloud Function.
package handler

import (
	"net/http"

	p "github.com/rfiestas/forecast"
)

// Handler is the ForecastAPIV1, Forecast call
func Handler(w http.ResponseWriter, r *http.Request) {
	p.ForecastAPIV1(w, r)
	return
}
