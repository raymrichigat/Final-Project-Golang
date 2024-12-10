package main

import (
	"go-web-native/config"
	"go-web-native/controllers/brandcontroller"
	"go-web-native/controllers/carcontroller"
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

	// 2. Merek
	http.HandleFunc("/brands", brandcontroller.Index)
	http.HandleFunc("/brands/add", brandcontroller.Add)
	http.HandleFunc("/brands/edit", brandcontroller.Edit)
	http.HandleFunc("/brands/delete", brandcontroller.Delete)

	// 3. Products
	http.HandleFunc("/cars", carcontroller.Index)
	// http.HandleFunc("/books/add", carcontroller.Add)
	// http.HandleFunc("/books/detail", carcontroller.Detail)
	// http.HandleFunc("/books/edit", carcontroller.Edit)
	// http.HandleFunc("/books/delete", carcontroller.Delete)

	// Run server
	log.Println("Server running on port: 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
