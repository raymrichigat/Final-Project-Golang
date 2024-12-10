package brandmodel

import (
	"fmt"
	"database/sql"
	"go-web-native/config"
	"go-web-native/entities"
	"time"
	"log"
	"strconv"
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
	} else {
			brand.DeletedAt = time.Time{} // Waktu default
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

func Detail(id string) entities.Brand {
	row := config.DB.QueryRow(`SELECT id, name FROM brands WHERE id = $1 `, id)

	var brand entities.Brand

	if err := row.Scan(&brand.Id, &brand.Name); err != nil {
		panic(err.Error())
	}

	return brand
}

func Update(id string, brand entities.Brand) bool {
	if id == "" {
		log.Println("Error: ID is empty")
		return false
	}

	// Konversi ID ke integer jika database memerlukan integer
	brandID, err := strconv.Atoi(id)
	if err != nil {
		log.Println("Error converting ID to integer:", err)
		return false
	}

	// Query untuk update brand
	query := `UPDATE brands SET name = $1, updated_at = $2 WHERE id = $3`
	_, err = config.DB.Exec(query, brand.Name, brand.UpdatedAt, brandID)
	if err != nil {
		log.Println("Error executing update query:", err)
		return false
	}

	return true
}


func Delete(id string) error {
	// Konversi ID dari string ke integer
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return fmt.Errorf("invalid ID format: %v", err)
	}

	// Eksekusi query DELETE
	_, err = config.DB.Exec("DELETE FROM brands WHERE id = $1", idInt)
	if err != nil {
		return fmt.Errorf("failed to delete brands: %v", err)
	}

	return nil
}