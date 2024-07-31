package database

import (
	"fmt"
)

type Post struct {
	Id    int64  `json:"id"`
	Slug  string `json:"slug"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

// Create new post
func (p *Post) NewPost() *Post {

	p.Slug = Slugify(p.Title)

	result, err := DB.Exec("INSERT INTO posts (title, body, slug) VALUES (?, ?,?)", p.Title, p.Body, p.Slug)

	fmt.Println("Insert Result ", result)

	if err != nil {
		return nil
	}

	id, err := result.LastInsertId()

	if err != nil {
		return nil
	}

	p.Id = id

	return p

}

// Update post

func (p *Post) Update(id int) {

}

// Delete

func (p *Post) Delete(id int) {

}

// Get
func (p *Post) Get(slug string) Post {

	var post Post
	row := DB.QueryRow("SELECT id, title, body, slug from posts WHERE slug=?", slug)
	err := row.Scan(&post.Id, &post.Title, &post.Body, &post.Slug)

	if err != nil {
		fmt.Println("Get post error", err)
	}

	fmt.Println(post)

	return post
}

// List

func (p *Post) List() []Post {
	rows, err := DB.Query("SELECT id, title, body, slug FROM posts")

	if err != nil {

		return nil
	}

	defer rows.Close()

	var posts []Post

	for rows.Next() {

		var post Post

		err := rows.Scan(&post.Id, &post.Title, &post.Body, &post.Slug)

		if err != nil {

			return nil

		}

		posts = append(posts, post)

	}

	return posts
}
