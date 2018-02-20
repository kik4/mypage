// +build !appengine

package server

import "net/http"

func start(h http.Handler) {
	http.ListenAndServe(":8080", h)
}
