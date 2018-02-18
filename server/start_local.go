// +build !appengine

package server

import (
	"github.com/kikkikkikkik/ocelot"
)

func start(o *ocelot.Ocelot) {
	o.Start(":8080")
}
