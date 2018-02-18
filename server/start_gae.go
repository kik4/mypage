// +build appengine

package server

import (
	"net/http"

	"github.com/kikkikkikkik/ocelot"
)

func start(o *ocelot.Ocelot) {
	http.Handle("/", o)
}
