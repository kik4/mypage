// +build appengine

package server

import (
	"net/http"
)

func start(h http.Handler) {
	http.Handle("/", h)
}
