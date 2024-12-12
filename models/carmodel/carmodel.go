package carmodel

import (
	"database/sql"
	"fmt"
	"go-web-native/config"
	"go-web-native/entities"
)

func GetAllCars() ([]entities.Car, error) {
	query := `
		SELECT 
			c.id, c.tipe, c.brand_id, c.license_plate, c.color, 
			c.price, c.image, c.description, c.created_at, c.updated_at, c.deleted_at,
			b.id, b.name, b.created_at, b.updated_at
		FROM cars c
		LEFT JOIN brands b ON c.brand_id = b.id
		WHERE c.deleted_at IS NULL
	`
	rows, err := config.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cars []entities.Car
	for rows.Next() {
		var car entities.Car
		var brand entities.Brand
		var deletedAt sql.NullTime

		err := rows.Scan(
			&car.Id, &car.Tipe, &car.BrandID, &car.LicensePlate, &car.Color, 
			&car.Price, &car.Image, &car.Description, &car.CreatedAt, &car.UpdatedAt, &deletedAt,
			&brand.Id, &brand.Name, &brand.CreatedAt, &brand.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		
		car.BrandName = brand
		if deletedAt.Valid {
			car.DeletedAt = deletedAt.Time
		}
		cars = append(cars, car)
	}
	return cars, nil
}

func AddCar(car entities.Car) error {
	// Check if brand exists
	var brandExists bool
	err := config.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM brands WHERE id = $1)", car.BrandID).Scan(&brandExists)
	if err != nil {
		return fmt.Errorf("error checking brand: %v", err)
	}
	if !brandExists {
		return fmt.Errorf("brand with ID %d does not exist", car.BrandID)
	}

	// Check if license plate already exists
	var count int
	err = config.DB.QueryRow("SELECT COUNT(*) FROM cars WHERE license_plate = $1", car.LicensePlate).Scan(&count)
	if err != nil {
		return fmt.Errorf("error checking license plate: %v", err)
	}
	if count > 0 {
		return fmt.Errorf("car with license plate '%s' already exists", car.LicensePlate)
	}

	// Insert new car
	query := `
		INSERT INTO cars (
			brand_id, tipe, license_plate, color, price, 
			image, description, created_at, updated_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7, NOW(), NOW())
	`
	_, err = config.DB.Exec(
		query, 
		car.BrandID, car.Tipe, car.LicensePlate, car.Color, 
		car.Price, car.Image, car.Description,
	)
	return err
}

func DeleteCar(id string) error {
	// Hard delete: Menghapus data secara permanen
	query := "DELETE FROM cars WHERE id = $1"

	// Eksekusi query
	_, err := config.DB.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error deleting car with ID %s: %v", id, err)
	}

	return nil
}

func DetailCar(id string) (entities.Car, error) {
	query := `
		SELECT 
			c.id, c.tipe, c.brand_id, c.license_plate, c.color, 
			c.price, c.image, c.description, c.created_at, c.updated_at,
			b.id, b.name, b.created_at, b.updated_at
		FROM cars c
		LEFT JOIN brands b ON c.brand_id = b.id
		WHERE c.id = $1
	`
	var car entities.Car
	var brand entities.Brand

	err := config.DB.QueryRow(query, id).Scan(
		&car.Id, &car.Tipe, &car.BrandID, &car.LicensePlate, &car.Color, 
		&car.Price, &car.Image, &car.Description, &car.CreatedAt, &car.UpdatedAt,
		&brand.Id, &brand.Name, &brand.CreatedAt, &brand.UpdatedAt,
	)
	if err != nil {
		return car, err
	}

	car.BrandName = brand
	return car, nil
}

func UpdateCar(id string, car entities.Car) error {
	// Check if car exists
	var exists bool
	err := config.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM cars WHERE id = $1)", id).Scan(&exists)
	if err != nil {
		return fmt.Errorf("error checking car existence: %v", err)
	}
	if !exists {
		return fmt.Errorf("car with ID %s does not exist", id)
	}

	// Update car without modifying the image column
	query := `
		UPDATE cars SET 
			brand_id = $1, tipe = $2, license_plate = $3, 
			color = $4, price = $5, 
			description = $6, updated_at = NOW()
		WHERE id = $7
	`
	_, err = config.DB.Exec(
		query,
		car.BrandID, car.Tipe, car.LicensePlate,
		car.Color, car.Price,
		car.Description, id,
	)
	return err
}