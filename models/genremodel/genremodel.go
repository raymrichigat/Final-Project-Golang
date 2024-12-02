package genremodel

import (
	"go-web-native/config"
	"go-web-native/entities"
)

func GetAll() []entities.Genre {
	rows, err := config.DB.Query(`SELECT * FROM genre`)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var categories []entities.Genre

	for rows.Next() {
		var category entities.Genre
		if err := rows.Scan(&category.Id, &category.Name, &category.CreatedAt, &category.UpdatedAt); err != nil {
			panic(err)
		}

		categories = append(categories, category)
	}

	return categories
}

func Create(category entities.Genre) bool {
	result, err := config.DB.Exec(`
		INSERT INTO genre (name, created_at, updated_at) 
		VALUE (?, ?, ?)`,
		category.Name,
		category.CreatedAt,
		category.UpdatedAt,
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
