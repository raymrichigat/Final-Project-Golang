package genrecontroller

import (
	"go-web-native/entities"
	"go-web-native/models/genremodel"
	"net/http"
	"strconv"
	"text/template"
	"time"
	"log"
)

func Index(w http.ResponseWriter, r *http.Request) {
	genres := genremodel.GetAll()
	data := map[string]any{
		"genres": genres,
	}

	temp, err := template.ParseFiles("views/genre/index.html")
	if err != nil {
		http.Error(w, "Error loading genres page", http.StatusInternalServerError)
    log.Println("Error:", err)
    return
	}

	temp.Execute(w, data)
}

func Add(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/genre/create.html")
		if err != nil {
			panic(err)
		}

		temp.Execute(w, nil)
	}

	if r.Method == "POST" {
		var category entities.Genre

		category.Name = r.FormValue("name")
		category.CreatedAt = time.Now()
		category.UpdatedAt = time.Now()

		ok := genremodel.Create(category)
		if !ok {
			temp, _ := template.ParseFiles("views/genre/create.html")
			temp.Execute(w, nil)
		}

		http.Redirect(w, r, "/genres", http.StatusSeeOther)
	}
}

func Edit(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/genre/edit.html")
		if err != nil {
			panic(err)
		}

		idString := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}

		genre := genremodel.Detail(id)
		data := map[string]any{
			"genre": genre,
		}

		temp.Execute(w, data)
	}

	if r.Method == "POST" {
		var category entities.Genre

		idString := r.FormValue("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}

		category.Name = r.FormValue("name")
		category.UpdatedAt = time.Now()

		if ok := genremodel.Update(id, category); !ok {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusTemporaryRedirect)
			return
		}

		http.Redirect(w, r, "/genres", http.StatusSeeOther)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(err)
	}

	if err := genremodel.Delete(id); err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/genres", http.StatusSeeOther)
}
