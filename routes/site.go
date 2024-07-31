package routes

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/saratonite/simple-blog/database"
)

func SiteRoutes() *chi.Mux {

	r := chi.NewRouter()

	// Home route
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {

		t, err := template.ParseFiles("templates/index.html")

		if err != nil {
			fmt.Fprintf(w, "Template error %s", err)
		}

		// List of post

		var posts []database.Post

		var p database.Post

		posts = p.List()

		t.Execute(w, posts)
	})

	// Single post

	r.Get("/posts/{slug}", func(w http.ResponseWriter, r *http.Request) {

		slug := chi.URLParam(r, "slug")

		var p database.Post

		post := p.Get(slug)

		t, e := template.ParseFiles("templates/post.html")

		if e != nil {
			fmt.Fprintf(w, "Error parsing template")
			return
		}

		if t == nil {
			fmt.Fprintf(w, "Post not found!")
		}

		t.Execute(w, post)

	})

	return r

}
