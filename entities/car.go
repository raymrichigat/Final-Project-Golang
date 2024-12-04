package entities

import "time"

// Car adalah entitas untuk mobil
type Car struct {
	Id          	uint       `json:"id"`
	BrandID  		uint       `json:"brand_id"` // Relasi dengan Brand menggunakan ID
	Type       		string     `json:"type"`
	LicensePlate	string	   `json:"license_plate"` // Sesuai dengan kolom di tabel
	Color			string	   `json:"color"`
	PublishedAt 	*time.Time `json:"published_at"` // Gunakan pointer untuk kolom nullable
	Description 	string     `json:"description"`
	CreatedAt   	time.Time  `json:"created_at"`
	UpdatedAt   	*time.Time `json:"updated_at"` // Gunakan pointer untuk kolom nullable
	DeletedAt 		*time.Time `json:"deleted_at"` // Gunakan pointer untuk kolom nullable

	Brand Brand `json:"brand,omitempty"` // Objek Brand terkait
}
