package brandcontroller

import (
	"go-web-native/entities"
	"go-web-native/models/brandmodel"
	"log"
	"net/http"
	"text/template"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request) {
	brands := brandmodel.GetAll()
	data := map[string]any{
		"brands": brands,
	}

	temp, err := template.ParseFiles("views/brand/index.html")
	if err != nil {
		http.Error(w, "Error loading genres page", http.StatusInternalServerError)
    log.Println("Error:", err)
    return
	}

	temp.Execute(w, data)
}

// AddBrand menangani GET untuk menampilkan form dan POST untuk menambah genre
func Add(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// Tampilkan halaman form
		temp, err := template.ParseFiles("views/brand/create.html")
		if err != nil {
			http.Error(w, "Error loading create brand page", http.StatusInternalServerError)
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
		http.Error(w, "Brand name is required", http.StatusBadRequest)
		return
	}

	// Log data yang diterima
	log.Println("Received brand name:", name)

	// Buat objek genre
	brand := entities.Brand{
		Name:      name,
		CreatedAt: time.Now(),
	}

	// Log objek genre sebelum disimpan
	log.Printf("Adding genre: %+v\n", brand)

	// Tambahkan genre ke database
	err := brandmodel.AddBrand(brand)
	if err != nil {
		// Jika error terjadi, tampilkan error dalam modal
		log.Println("Error adding genre:", err)
		data := map[string]interface{}{
			"error": err.Error(), // Kirimkan pesan error ke template
		}
		temp, _ := template.ParseFiles("views/brand/create.html")
		temp.Execute(w, data)
		return
	}

	// Setelah berhasil menambahkan, redirect ke halaman genres
	log.Println("Brand New added successfully")
	http.Redirect(w, r, "/brands", http.StatusSeeOther)
}

// func Edit(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == "GET" {
// 		temp, err := template.ParseFiles("views/genre/edit.html")
// 		if err != nil {
// 			panic(err)
// 		}

// 		idString := r.URL.Query().Get("id")
// 		id, err := strconv.Atoi(idString)
// 		if err != nil {
// 			panic(err)
// 		}

// 		genre := genremodel.Detail(id)
// 		data := map[string]any{
// 			"genre": genre,
// 		}

// 		temp.Execute(w, data)
// 	}

// 	if r.Method == "POST" {
// 		var category entities.Genre

// 		idString := r.FormValue("id")
// 		id, err := strconv.Atoi(idString)
// 		if err != nil {
// 			panic(err)
// 		}

// 		category.Name = r.FormValue("name")
// 		category.UpdatedAt = time.Now()

// 		if ok := genremodel.Update(id, category); !ok {
// 			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusTemporaryRedirect)
// 			return
// 		}

// 		http.Redirect(w, r, "/genres", http.StatusSeeOther)
// 	}
// }

// func Delete(w http.ResponseWriter, r *http.Request) {
// 	idString := r.URL.Query().Get("id")

// 	id, err := strconv.Atoi(idString)
// 	if err != nil {
// 		panic(err)
// 	}

// 	if err := genremodel.Delete(id); err != nil {
// 		panic(err)
// 	}

// 	http.Redirect(w, r, "/genres", http.StatusSeeOther)
// }
