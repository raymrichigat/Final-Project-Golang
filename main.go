package main

import (
	"go-web-native/config"
	"go-web-native/controllers/genrecontroller"
	"go-web-native/controllers/homecontroller"
	"go-web-native/controllers/bookcontroller"
	"log"
	"net/http"
)

func main() {
	// 	Database connection
	config.ConnectDB()

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
