package main

import (
	"fmt"
	"net/http"

	"github.com/kikkikkikkik/ocelot"
)

func init() {

	var o = ocelot.New()

	o.Register("get", "/", func(w http.ResponseWriter, r *http.Request) error {
		fmt.Fprintf(w, "Welcome!")
		return nil
	})

	http.Handle("/", o)
}
