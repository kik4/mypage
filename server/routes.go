package server

import (
	"html/template"
	"net/http"
	"time"

	"github.com/kikkikkikkik/ocelot"
)

// Init server
func Init() {
	// load templates
	tIndex := template.Must(template.ParseFiles("./content/templates/index.tmpl"))

	// new Framework
	var o = ocelot.New()

	// register routes
	o.Register("get", "/", func(w http.ResponseWriter, r *http.Request) error {
		w.Header().Set("cache-control", "public, max-age=3")

		return tIndex.Execute(w, map[string]interface{}{
			"Now": time.Now().String(),
		})
	})

	// start server
	start(o)
}
