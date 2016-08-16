// Package status is middleware for returning status code for requests
package status

import (
	"net/http"

	"github.com/mholt/caddy/caddyhttp/httpserver"
)

// Status is a middleware to return status code for request
type Status struct {
	Rules map[string]int
	Next  httpserver.Handler
}

// ServeHTTP implements the httpserver.Handler interface
func (status Status) ServeHTTP(w http.ResponseWriter, r *http.Request) (int, error) {
	if statusCode, ok := status.Rules[r.URL.Path]; ok {
		return statusCode, nil
	}

	return status.Next.ServeHTTP(w, r)
}
