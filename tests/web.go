package tests

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

// StartWebServer starts a testing HTTP server with static content
func StartWebServer(requests map[string]string) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for uri, content := range requests {
			if r.RequestURI == uri {
				_, _ = fmt.Fprintln(w, content)
			}
		}
	}))
}
