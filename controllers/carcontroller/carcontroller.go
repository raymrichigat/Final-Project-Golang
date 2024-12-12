package carcontroller

import (
	"go-web-native/entities"
	"go-web-native/models/brandmodel"
	"go-web-native/models/carmodel"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	cars, err := carmodel.GetAllCars()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "home-car.html", gin.H{
			"error": "Failed to load cars",
		})
		return
	}
	c.HTML(http.StatusOK, "home-car.html", gin.H{
		"cars": cars,
	})
}

func AddForm(c *gin.Context) {
	brands := brandmodel.GetAll()
	c.HTML(http.StatusOK, "create-car.html", gin.H{
		"brands": brands,
	})
}

func Add(c *gin.Context) {
	// Get form values
	brandID, _ := strconv.ParseUint(c.PostForm("brand_id"), 10, 32)
	price, _ := strconv.ParseFloat(c.PostForm("price"), 64)

	// Handle file upload
	file, err := c.FormFile("image")
	if err != nil {
		c.HTML(http.StatusBadRequest, "create-car.html", gin.H{
			"error": "Image upload is required",
		})
		return
	}

	// Save the file to the "public" folder
	filePath := "./public/" + file.Filename
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.HTML(http.StatusInternalServerError, "create-car.html", gin.H{
			"error": "Failed to save the image",
		})
		return
	}

	// Create car entity
	car := entities.Car{
		BrandID:      uint(brandID),
		Tipe:         c.PostForm("tipe"),
		LicensePlate: c.PostForm("license_plate"),
		Color:        c.PostForm("color"),
		Price:        price,
		Image:        "/static/" + file.Filename, // Save relative path to the image
		Description:  c.PostForm("description"),
	}

	// Validate required fields
	if car.Tipe == "" || car.LicensePlate == "" || car.BrandID == 0 {
		brands := brandmodel.GetAll()
		c.HTML(http.StatusBadRequest, "create-car.html", gin.H{
			"error":  "Tipe, License Plate, and Brand are required",
			"brands": brands,
			"car":    car,
		})
		return
	}

	// Add car
	if err := carmodel.AddCar(car); err != nil {
		brands := brandmodel.GetAll()
		c.HTML(http.StatusInternalServerError, "create-car.html", gin.H{
			"error":  err.Error(),
			"brands": brands,
			"car":    car,
		})
		return
	}

	c.Redirect(http.StatusSeeOther, "/cars")
}

func Delete(c *gin.Context) {
    id := c.Param("id")

    if err := carmodel.DeleteCar(id); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.Redirect(http.StatusSeeOther, "/cars")
}
