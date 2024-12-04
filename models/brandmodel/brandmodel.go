package brandmodel

import (
	"fmt"
	"go-web-native/config"
	"go-web-native/entities"
)

func GetAll() []entities.Brand {
	rows, err := config.DB.Query(`SELECT * FROM brands`)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var brands []entities.Brand

	for rows.Next() {
		var brand entities.Brand
		if err := rows.Scan(&brand.Id, &brand.Name, &brand.CreatedAt, &brand.UpdatedAt, &brand.UpdatedAt); err != nil {
			panic(err)
		}

		brands = append(brands, brand)
	}

	return brands
}

func AddBrand(brand entities.Brand) error {
	// Cek apakah genre sudah ada
	var count int
	err := config.DB.QueryRow("SELECT COUNT(*) FROM brands WHERE name = $1", brand.Name).Scan(&count)
	if err != nil {
		return fmt.Errorf("error checking if brand exists: %v", err)
	}

	if count > 0 {
		return fmt.Errorf("brand with name '%s' already exists", brand.Name)
	}

	// Menyimpan genre baru
	_, err = config.DB.Exec(
		"INSERT INTO brands (name, created_at, updated_at) VALUES ($1, NOW(), NOW())",
		brand.Name,
	)
	if err != nil {
		return fmt.Errorf("error inserting brand: %v", err)
	}

	return nil
}

// func Detail(id int) entities.Merek {
// 	row := config.DB.QueryRow(`SELECT id, name FROM genre WHERE id = ? `, id)

// 	var category entities.Merek

// 	if err := row.Scan(&category.Id, &category.Name); err != nil {
// 		panic(err.Error())
// 	}

// 	return category
// }

// func Update(id int, category entities.Merek) bool {
// 	query, err := config.DB.Exec(`UPDATE genre SET name = ?, updated_at = ? where id = ?`, category.Name, category.UpdatedAt, id)
// 	if err != nil {
// 		panic(err)
// 	}

// 	result, err := query.RowsAffected()
// 	if err != nil {
// 		panic(err)
// 	}

// 	return result > 0
// }

// func Delete(id int) error {
// 	_, err := config.DB.Exec("DELETE FROM genre WHERE id = ?", id)
// 	return err
// }
