package brandmodel

import (
	"database/sql"
	"fmt"
	"go-web-native/config"
	"go-web-native/entities"
	"log"
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
		var deletedAt sql.NullTime
		if err := rows.Scan(&brand.Id, &brand.Name, &brand.CreatedAt, &brand.UpdatedAt, &deletedAt); err != nil {
			panic(err)
		}
		if deletedAt.Valid {
			brand.DeletedAt = deletedAt.Time
		}
		brands = append(brands, brand)
	}
	return brands
}

func AddBrand(brand entities.Brand) error {
	var count int
	err := config.DB.QueryRow("SELECT COUNT(*) FROM brands WHERE name = $1", brand.Name).Scan(&count)
	if err != nil {
		return fmt.Errorf("error checking if brand exists: %v", err)
	}
	if count > 0 {
		return fmt.Errorf("brand with name '%s' already exists", brand.Name)
	}

	_, err = config.DB.Exec(
		"INSERT INTO brands (name, created_at, updated_at) VALUES ($1, NOW(), NOW())",
		brand.Name,
	)
	return err
}

func Detail(id string) entities.Brand {
	row := config.DB.QueryRow(`SELECT id, name FROM brands WHERE id = $1`, id)
	var brand entities.Brand
	if err := row.Scan(&brand.Id, &brand.Name); err != nil {
		panic(err)
	}
	return brand
}

func Update(id string, brand entities.Brand) bool {
	_, err := config.DB.Exec(
		"UPDATE brands SET name = $1, updated_at = $2 WHERE id = $3",
		brand.Name, brand.UpdatedAt, id,
	)
	if err != nil {
		log.Println("Error updating brand:", err)
		return false
	}
	return true
}

func Delete(id string) error {
	_, err := config.DB.Exec("DELETE FROM brands WHERE id = $1", id)
	return err
}
