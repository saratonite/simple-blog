package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/saratonite/simple-blog/database"
	"github.com/saratonite/simple-blog/routes"
)

func main() {

	database.Connect("data.db")

	fmt.Println("Database connected")

	database.InitTables()

	r := chi.NewRouter()

	// Home route
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {

		t, err := template.ParseFiles("templates/index.html")

		if err != nil {
			fmt.Fprintf(w, "Template error %s", err)
		}

		t.Execute(w, struct{}{})
	})

	// Mount API Endpoints
	r.Mount("/api", routes.APIRoutes())

	fmt.Println("Server starting")

	err := http.ListenAndServe(":3000", r)

	if err != nil {

		log.Fatal("Error staring server ", err)
	}

}
