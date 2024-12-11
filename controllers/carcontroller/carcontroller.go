package carcontroller

import (
	"go-web-native/models/carmodel"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	cars, err := carmodel.GetAllCars()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "car/index.html", gin.H{
			"error": "Failed to load cars",
		})
		return
	}
	c.HTML(http.StatusOK, "car/index.html", gin.H{
		"cars": cars,
	})
}


// func Add(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == "GET" {
// 		// Render add book form
// 		temp, err := template.ParseFiles("views/book/create.html")
// 		if err != nil {
// 			panic(err)
// 		}

// 		// Get all genres to show in the form
// 		genres := genremodel.GetAll()
// 		data := map[string]any{
// 			"genres": genres,
// 		}

// 		// Execute the template with genre data
// 		temp.Execute(w, data)
// 	}

// 	if r.Method == "POST" {
// 		// Handle form submission
// 		var book entities.Book

// 		// Parse category (genre) ID
// 		genreId, err := strconv.Atoi(r.FormValue("genre_id"))
// 		if err != nil {
// 			panic(err)
// 		}

// 		// Parse other form fields
// 		book.Title = r.FormValue("title")
// 		book.GenreID = uint(genreId)
// 		book.Description = r.FormValue("description")
// 		book.PublishedAt = time.Now() // You may want to parse this from the form as well
// 		book.CreatedAt = time.Now()
// 		book.UpdatedAt = time.Now()

// 		// Create the book in the database
// 		if ok := bookmodel.Create(book); !ok {
// 			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusTemporaryRedirect)
// 			return
// 		}

// 		http.Redirect(w, r, "/books", http.StatusSeeOther)
// 	}
// }

// func Detail(w http.ResponseWriter, r *http.Request) {
// 	// Get book ID from URL
// 	idString := r.URL.Query().Get("id")

// 	// Convert ID to integer
// 	id, err := strconv.Atoi(idString)
// 	if err != nil {
// 		panic(err)
// 	}

// 	// Get book details from the model
// 	book := bookmodel.Detail(id)
// 	data := map[string]any{
// 		"book": book,
// 	}

// 	// Render the detail page
// 	temp, err := template.ParseFiles("views/book/detail.html")
// 	if err != nil {
// 		panic(err)
// 	}

// 	// Execute template with book data
// 	temp.Execute(w, data)
// }

// func Edit(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == "GET" {
// 		// Get book ID from URL
// 		idString := r.URL.Query().Get("id")
// 		id, err := strconv.Atoi(idString)
// 		if err != nil {
// 			panic(err)
// 		}

// 		// Get book details and genres
// 		book := bookmodel.Detail(id)
// 		genres := genremodel.GetAll()

// 		// Prepare data for the template
// 		data := map[string]any{
// 			"book":   book,
// 			"genres": genres,
// 		}

// 		// Render edit form with data
// 		temp, err := template.ParseFiles("views/book/edit.html")
// 		if err != nil {
// 			panic(err)
// 		}

// 		temp.Execute(w, data)
// 	}

// 	if r.Method == "POST" {
// 		// Handle form submission
// 		var book entities.Book

// 		// Get ID from form
// 		idString := r.FormValue("id")
// 		id, err := strconv.Atoi(idString)
// 		if err != nil {
// 			panic(err)
// 		}

// 		// Get category (genre) ID from form
// 		genreId, err := strconv.Atoi(r.FormValue("genre_id"))
// 		if err != nil {
// 			panic(err)
// 		}

// 		// Parse other form fields
// 		book.Title = r.FormValue("title")
// 		book.GenreID = uint(genreId)
// 		book.Description = r.FormValue("description")
// 		book.PublishedAt = time.Now() // Again, may want to parse this as well
// 		book.UpdatedAt = time.Now()

// 		// Update the book in the database
// 		if ok := bookmodel.Update(id, book); !ok {
// 			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusTemporaryRedirect)
// 			return
// 		}

// 		http.Redirect(w, r, "/books", http.StatusSeeOther)
// 	}
// }

// func Delete(w http.ResponseWriter, r *http.Request) {
// 	// Get book ID from URL
// 	idString := r.URL.Query().Get("id")

// 	// Convert ID to integer
// 	id, err := strconv.Atoi(idString)
// 	if err != nil {
// 		panic(err)
// 	}

// 	// Delete the book from the database
// 	if err := bookmodel.Delete(id); err != nil {
// 		panic(err)
// 	}

// 	// Redirect to the book list page
// 	http.Redirect(w, r, "/books", http.StatusSeeOther)
// }
