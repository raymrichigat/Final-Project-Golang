package entities

import "time"

// Car adalah entitas untuk mobil
type Car struct {
	Id            uint      `json:"id"`            // ID mobil
	BrandID       uint      `json:"brand_id"`      // ID brand, sebagai foreign key
	Tipe         string    `json:"title"`         // Judul mobil
	LicensePlate  string    `json:"license_plate"` // Nomor plat mobil
	Color         string    `json:"color"`         // Warna mobil
	Price         float64   `json:"price"`         // Harga mobil
	Image         string    `json:"image,omitempty"` // Tidak wajib dikirimkan
	Description   string    `json:"description"`   // Deskripsi mobil
	CreatedAt     time.Time `json:"created_at"`     // Tanggal pembuatan
	UpdatedAt     time.Time `json:"updated_at"`     // Tanggal update
	DeletedAt     time.Time `json:"deleted_at"`     // Tanggal penghapusan

	BrandName Brand `json:"brand,omitempty"` // Relasi dengan Brand, nama brand akan ditampilkan di sini
}
