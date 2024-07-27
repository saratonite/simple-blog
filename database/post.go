package database

import "fmt"

type Post struct {
	Id    int    `json:"id"`
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

	return p

}

// Update post

func (p *Post) Update(id int) {

}

// Delete

func (p *Post) Delete(id int) {

}

// Get
func (p *Post) Get(slug string) {

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
