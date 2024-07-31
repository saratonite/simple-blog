package routes

import (
	"github.com/go-chi/chi/v5"
)

func PostRoutes() *chi.Mux {
	r := chi.NewRouter()

	// List all posts
	r.Get("/", ListPosts)

	// Get single post
	r.Get("/{slug}", GetPost)

	// Create post
	r.Post("/", CreatePost)

	// Update post
	r.Put("/{id}", UpdatePost)

	// Delete post
	r.Delete("/{id}", DeletePost)

	return r
}
