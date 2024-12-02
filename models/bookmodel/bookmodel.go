package bookmodel

import (
	"go-web-native/config"
	"go-web-native/entities"
)

// func Getall() []entities.Book {
// 	rows, err := config.DB.Query(`
// 		SELECT 
// 			products.id, 
// 			products.name, 
// 			categories.name as category_name,
// 			products.stock, 
// 			products.description, 
// 			products.created_at, 
// 			products.updated_at FROM products
// 		JOIN categories ON products.category_id = categories.id
// 	`)

// 	if err != nil {
// 		panic(err)
// 	}

// 	defer rows.Close()

// 	var products []entities.Product

// 	for rows.Next() {
// 		var product entities.Product
// 		if err := rows.Scan(
// 			&product.Id,
// 			&product.Name,
// 			&product.Category.Name,
// 			&product.Stock,
// 			&product.Description,
// 			&product.CreatedAt,
// 			&product.UpdatedAt,
// 		); err != nil {
// 			panic(err)
// 		}

// 		products = append(products, product)
// 	}

// 	return products
// }

func GetAll() ([]entities.Book, error) {
	rows, err := config.DB.Query(`
		SELECT 
			books.id, 
			books.title, 
			books.category_id,
			genres.name as genre_name,
			books.published_at, 
			books.description, 
			books.created_at, 
			books.updated_at
		FROM books
		JOIN genres ON books.category_id = genres.id
	`)
	if err != nil {
		return nil, err // kembalikan error jika ada masalah dengan query
	}
	defer rows.Close()

	var books []entities.Book
	for rows.Next() {
		var book entities.Book
		if err := rows.Scan(
			&book.Id,
			&book.Title,
			&book.GenreID, // Assign category_id
			&book.Genre.Name, // Assign genre name
			&book.PublishedAt,
			&book.Description,
			&book.CreatedAt,
			&book.UpdatedAt,
		); err != nil {
			return nil, err // kembalikan error jika ada masalah saat scan
		}
		books = append(books, book)
	}

	return books, nil
}

// func Create(product entities.Product) bool {
// 	result, err := config.DB.Exec(`
// 		INSERT INTO products(
// 			name, category_id, stock, description, created_at, updated_at
// 		) VALUES (?, ?, ?, ?, ?, ?)`,
// 		product.Name,
// 		product.Category.Id,
// 		product.Stock,
// 		product.Description,
// 		product.CreatedAt,
// 		product.UpdatedAt,
// 	)

// 	if err != nil {
// 		panic(err)
// 	}

// 	lastInsertId, err := result.LastInsertId()
// 	if err != nil {
// 		panic(err)
// 	}

// 	return lastInsertId > 0
// }

func Create(book entities.Book) bool {
	result, err := config.DB.Exec(`
		INSERT INTO books(
			title, category_id, published_at, description, created_at, updated_at
		) VALUES (?, ?, ?, ?, ?, ?)`,
		book.Title,
		book.GenreID,
		book.PublishedAt,
		book.Description,
		book.CreatedAt,
		book.UpdatedAt,
	)

	if err != nil {
		panic(err)
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	return lastInsertId > 0
}


// func Detail(id int) entities.Product {
// 	row := config.DB.QueryRow(`
// 		SELECT 
// 			products.id, 
// 			products.name, 
// 			categories.name as category_name,
// 			products.stock, 
// 			products.description, 
// 			products.created_at, 
// 			products.updated_at FROM products
// 		JOIN categories ON products.category_id = categories.id
// 		WHERE products.id = ?
// 	`, id)

// 	var product entities.Product

// 	err := row.Scan(
// 		&product.Id,
// 		&product.Name,
// 		&product.Category.Name,
// 		&product.Stock,
// 		&product.Description,
// 		&product.CreatedAt,
// 		&product.UpdatedAt,
// 	)

// 	if err != nil {
// 		panic(err)
// 	}

// 	return product
// }

func Detail(id int) entities.Book {
	row := config.DB.QueryRow(`
		SELECT 
			books.id, 
			books.title, 
			books.category_id,
			genres.name as genre_name,
			books.published_at, 
			books.description, 
			books.created_at, 
			books.updated_at 
		FROM books
		JOIN genres ON books.category_id = genres.id
		WHERE books.id = ?`,
		id)

	var book entities.Book

	err := row.Scan(
		&book.Id,
		&book.Title,
		&book.GenreID, // Relasi dengan Genre menggunakan CategoryID
		&book.Genre.Name, // Nama genre
		&book.PublishedAt,
		&book.Description,
		&book.CreatedAt,
		&book.UpdatedAt,
	)

	if err != nil {
		panic(err)
	}

	return book
}


// func Update(id int, product entities.Product) bool {
// 	query, err := config.DB.Exec(`
// 		UPDATE products SET 
// 			name = ?, 
// 			category_id = ?,
// 			stock = ?,
// 			description = ?,
// 			updated_at = ?
// 		WHERE id = ?`,
// 		product.Name,
// 		product.Category.Id,
// 		product.Stock,
// 		product.Description,
// 		product.UpdatedAt,
// 		id,
// 	)

// 	if err != nil {
// 		panic(err)
// 	}

// 	result, err := query.RowsAffected()
// 	if err != nil {
// 		panic(err)
// 	}

// 	return result > 0
// }

func Update(id int, book entities.Book) bool {
	query, err := config.DB.Exec(`
		UPDATE books SET 
			title = ?, 
			category_id = ?,
			published_at = ?,
			description = ?,
			updated_at = ?
		WHERE id = ?`,
		book.Title,
		book.GenreID,
		book.PublishedAt,
		book.Description,
		book.UpdatedAt,
		id,
	)

	if err != nil {
		panic(err)
	}

	result, err := query.RowsAffected()
	if err != nil {
		panic(err)
	}

	return result > 0
}


// func Delete(id int) error {
// 	_, err := config.DB.Exec("DELETE FROM products WHERE id = ?", id)
// 	return err
// }

func Delete(id int) error {
	_, err := config.DB.Exec("DELETE FROM books WHERE id = ?", id)
	return err
}
