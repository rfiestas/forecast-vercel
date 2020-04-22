// Package forecast contains an HTTP Cloud Function.
package forecast

import (
	"net/http"
	"net/url"
)

// GetQueryKey : take a http request url query key, assign default value when not exist.
func GetQueryKey(r *http.Request, key string, failure string) string {
	var value string
	keys, ok := r.URL.Query()[key]
	if !ok || len(keys[0]) < 1 {
		value = failure
	} else {
		value = string(keys[0])
	}
	return url.QueryEscape(value)
}
