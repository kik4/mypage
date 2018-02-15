package app

import (
	"html/template"
	"net/http"

	"github.com/kikkikkikkik/ocelot"
)

func init() {
	t := template.Must(template.ParseFiles("templates/index.tmpl"))

	var o = ocelot.New()

	o.Register("get", "/", func(w http.ResponseWriter, r *http.Request) error {
		return t.Execute(w, map[string]interface{}{
			"Language": "Golang",
			"Stars":    5,
		})
	})

	http.Handle("/", o)
}
