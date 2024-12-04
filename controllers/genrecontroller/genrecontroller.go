package genrecontroller

import (
	"go-web-native/entities"
	"go-web-native/models/genremodel"
	"log"
	"net/http"
	"strconv"
	"text/template"
	"time"
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

// AddGenre menangani GET untuk menampilkan form dan POST untuk menambah genre
func Add(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// Tampilkan halaman form
		temp, err := template.ParseFiles("views/genre/create.html")
		if err != nil {
			http.Error(w, "Error loading create genre page", http.StatusInternalServerError)
			log.Println("Error:", err)
			return
		}

		temp.Execute(w, nil)
		return
	}

	// Proses form submission
	r.ParseForm()
	name := r.FormValue("name")

	// Cek apakah nama genre kosong
	if name == "" {
		http.Error(w, "Genre name is required", http.StatusBadRequest)
		return
	}

	// Log data yang diterima
	log.Println("Received genre name:", name)

	// Buat objek genre
	genre := entities.Genre{
		Name:      name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Log objek genre sebelum disimpan
	log.Printf("Adding genre: %+v\n", genre)

	// Tambahkan genre ke database
	err := genremodel.AddGenre(genre)
	if err != nil {
		// Jika error terjadi, tampilkan error dalam modal
		log.Println("Error adding genre:", err)
		data := map[string]interface{}{
			"error": err.Error(), // Kirimkan pesan error ke template
		}
		temp, _ := template.ParseFiles("views/genre/create.html")
		temp.Execute(w, data)
		return
	}

	// Setelah berhasil menambahkan, redirect ke halaman genres
	log.Println("Genre added successfully")
	http.Redirect(w, r, "/genres", http.StatusSeeOther)
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
