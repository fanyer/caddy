package status

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mholt/caddy/caddyhttp/httpserver"
)

func TestStatus(t *testing.T) {
	status := Status{
		Rules: map[string]int{
			"/foo":   http.StatusNotFound,
			"bar":    http.StatusServiceUnavailable,
			"teapot": http.StatusTeapot,
		},
		Next: httpserver.HandlerFunc(urlPrinter),
	}

	tests := []struct {
		path           string
		statusExpected bool
		status         int
	}{
		{"/foo", true, http.StatusNotFound},
		{"bar", true, http.StatusServiceUnavailable},
		{"teapot", true, http.StatusTeapot},
		{"/foobar", false, 0},
	}

	for i, test := range tests {
		req, err := http.NewRequest("GET", test.path, nil)
		if err != nil {
			t.Fatalf("Test %d: Could not create HTTP request: %v",
				i, err)
		}

		rec := httptest.NewRecorder()
		actualStatus, err := status.ServeHTTP(rec, req)

		if err != nil {
			t.Fatalf("Test %d: Serving request failed with error %v",
				i, err)
		}

		if test.statusExpected {
			if test.status != actualStatus {
				t.Errorf("Test %d: Expected status code %d, got %d",
					i, test.status, actualStatus)
			}
			if rec.Body.String() != "" {
				t.Errorf("Test %d: Expected empty body, got '%s'",
					i, rec.Body.String())
			}
		} else if rec.Body.String() != test.path {
			t.Errorf("Test %d: Expected body '%s', got '%s'",
				i, test.path, rec.Body.String())
		}
	}
}

func urlPrinter(w http.ResponseWriter, r *http.Request) (int, error) {
	fmt.Fprint(w, r.URL.String())
	return 0, nil
}
