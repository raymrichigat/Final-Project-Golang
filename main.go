package main

import (
	"go-web-native/config"
	"go-web-native/controllers/brandcontroller"
	"go-web-native/controllers/carcontroller"
	"go-web-native/controllers/homecontroller"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func init() {
	// Membuka file log
	file, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
}

func main() {
	// Database connection
	config.ConnectDB()
	defer config.DB.Close()

	// Initialize Gin engine
	router := gin.Default()

	// Load templates
	router.LoadHTMLGlob("views/*")

	// Routes
	router.GET("/", homecontroller.Welcome)

	// Brand routes
	brandRoutes := router.Group("/brands")
	{
		brandRoutes.GET("/", brandcontroller.Index)
		brandRoutes.GET("/add", brandcontroller.AddForm)
		brandRoutes.POST("/add", brandcontroller.Add)
		brandRoutes.GET("/edit/:id", brandcontroller.EditForm)
		brandRoutes.POST("/edit/:id", brandcontroller.Edit)
		brandRoutes.POST("/delete/:id", brandcontroller.Delete)
	}

	// Car routes
	carRoutes := router.Group("/cars")
	{
		carRoutes.GET("/", carcontroller.Index)
	}

	log.Println("Server running on port 8080")
	router.Run(":8080")
}
