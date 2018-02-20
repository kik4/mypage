package server

import (
	"html/template"
	"net/http"
	"time"

	"github.com/pressly/chi"
)

// Init server
func Init() {
	// load templates
	tIndex := template.Must(template.ParseFiles("./content/templates/index.tmpl"))

	// new router
	r := chi.NewRouter()

	// register routes
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("cache-control", "public, max-age=3")

		tIndex.Execute(w, map[string]interface{}{
			"Now": time.Now().String(),
		})
	})

	// start server
	start(r)
}
