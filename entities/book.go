package entities

import "time"

// Book adalah entitas untuk buku
type Book struct {
	Id          uint      `json:"id"`
	Title       string    `json:"title"`
	GenreID  uint      `json:"genre_id"` // Relasi dengan Category menggunakan ID
	PublishedAt time.Time `json:"published_at"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Description string    `json:"description"`

	// Category dan Author sebagai objek jika diperlukan untuk pengambilan data terkait
	Genre Genre  `json:"category,omitempty"`
}
