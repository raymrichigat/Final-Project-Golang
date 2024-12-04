package genremodel

import (
	"fmt"
	"go-web-native/config"
	"go-web-native/entities"
)

func GetAll() []entities.Genre {
	rows, err := config.DB.Query(`SELECT * FROM genres`)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var genres []entities.Genre

	for rows.Next() {
		var genre entities.Genre
		if err := rows.Scan(&genre.Id, &genre.Name, &genre.CreatedAt, &genre.UpdatedAt); err != nil {
			panic(err)
		}

		genres = append(genres, genre)
	}

	return genres
}

func AddGenre(genre entities.Genre) error {
	// Cek apakah genre sudah ada
	var count int
	err := config.DB.QueryRow("SELECT COUNT(*) FROM genres WHERE name = $1", genre.Name).Scan(&count)
	if err != nil {
		return fmt.Errorf("error checking if genre exists: %v", err)
	}

	if count > 0 {
		return fmt.Errorf("genre with name '%s' already exists", genre.Name)
	}

	// Menyimpan genre baru
	_, err = config.DB.Exec(
		"INSERT INTO genres (name, created_at, updated_at) VALUES ($1, NOW(), NOW())",
		genre.Name,
	)
	if err != nil {
		return fmt.Errorf("error inserting genre: %v", err)
	}

	return nil
}

func Detail(id int) entities.Genre {
	row := config.DB.QueryRow(`SELECT id, name FROM genre WHERE id = ? `, id)

	var category entities.Genre

	if err := row.Scan(&category.Id, &category.Name); err != nil {
		panic(err.Error())
	}

	return category
}

func Update(id int, category entities.Genre) bool {
	query, err := config.DB.Exec(`UPDATE genre SET name = ?, updated_at = ? where id = ?`, category.Name, category.UpdatedAt, id)
	if err != nil {
		panic(err)
	}

	result, err := query.RowsAffected()
	if err != nil {
		panic(err)
	}

	return result > 0
}

func Delete(id int) error {
	_, err := config.DB.Exec("DELETE FROM genre WHERE id = ?", id)
	return err
}
