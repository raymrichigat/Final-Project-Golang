package main

import (
	"go-web-native/config"
	"go-web-native/controllers/bookcontroller"
	"go-web-native/controllers/genrecontroller"
	"go-web-native/controllers/homecontroller"
	"log"
	"net/http"
	"os"
)

func init() {
	// Membuka file log
	file, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	// Menetapkan output log ke file
	log.SetOutput(file)
}

func main() {
	// 	Database connection
	config.ConnectDB()
	defer config.DB.Close() // Pastikan koneksi database ditutup saat aplikasi berhenti

	// Routes
	// 1.Homepage
	http.HandleFunc("/", homecontroller.Welcome)

	// 2. Category
	http.HandleFunc("/genres", genrecontroller.Index)
	http.HandleFunc("/genres/add", genrecontroller.Add)
	http.HandleFunc("/genres/edit", genrecontroller.Edit)
	http.HandleFunc("/genres/delete", genrecontroller.Delete)

	// 3. Products
	http.HandleFunc("/books", bookcontroller.Index)
	http.HandleFunc("/books/add", bookcontroller.Add)
	http.HandleFunc("/books/detail", bookcontroller.Detail)
	http.HandleFunc("/books/edit", bookcontroller.Edit)
	http.HandleFunc("/books/delete", bookcontroller.Delete)

	// Run server
	log.Println("Server running on port: 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
