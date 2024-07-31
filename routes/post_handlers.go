package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/saratonite/simple-blog/database"
)

// List post handler
func ListPosts(w http.ResponseWriter, r *http.Request) {
	var post database.Post = database.Post{}
	posts := post.List()
	JsonResponse(w, posts)
}

// Get post handler

func GetPost(w http.ResponseWriter, r *http.Request) {

	slug := chi.URLParam(r, "slug")

	var p database.Post

	post := p.Get(slug)

	JsonResponse(w, post)

}

// Create new post

func CreatePost(w http.ResponseWriter, r *http.Request) {

	var post database.Post

	err := json.NewDecoder(r.Body).Decode(&post)

	if err != nil {
		fmt.Println("Json decode error.")
	}

	fmt.Println(">> Incoming post data ", post)
	post.NewPost()

	JsonResponse(w, post)

}

// Update post

func UpdatePost(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "PUT /{id} Update post")
}

// Delete post

func DeletePost(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "DELETE /{id} Delete post")

}
