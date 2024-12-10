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

// AddBrand menangani GET untuk menampilkan form dan POST untuk menambah brand
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
	log.Printf("Adding brand: %+v\n", brand)

	// Tambahkan genre ke database
	err := brandmodel.AddBrand(brand)
	if err != nil {
		// Jika error terjadi, tampilkan error dalam modal
		log.Println("Error adding brand:", err)
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

func Edit(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/brand/edit.html")
		if err != nil {
			http.Error(w, "Error loading edit brand page", http.StatusInternalServerError)
			log.Println("Error:", err)
			return
		}

		idString := r.URL.Query().Get("id")
		if idString == "" {
			http.Error(w, "ID is required", http.StatusBadRequest)
			return
		}

		brand := brandmodel.Detail(idString)
		data := map[string]any{
			"brand": brand,
		}

		temp.Execute(w, data)
		return
	}

	if r.Method == "POST" {
		var category entities.Brand

		idString := r.FormValue("id")
		if idString == "" {
			http.Error(w, "ID is required", http.StatusBadRequest)
			return
		}

		category.Name = r.FormValue("name")
		category.UpdatedAt = time.Now()

		if ok := brandmodel.Update(idString, category); !ok {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusTemporaryRedirect)
			return
		}
		
		log.Println("Brand Edited successfully")
		http.Redirect(w, r, "/brands", http.StatusSeeOther)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	// Validasi apakah ID kosong
	if id == "" {
			http.Error(w, "ID is required", http.StatusBadRequest)
			return
	}

	// Hapus brand berdasarkan ID
	if err := brandmodel.Delete(id); err != nil {
			http.Error(w, "Failed to delete brand", http.StatusInternalServerError)
			log.Println("Error deleting brand:", err)
			return
	}

	// Redirect setelah sukses menghapus
	http.Redirect(w, r, "/brands", http.StatusSeeOther)
}
