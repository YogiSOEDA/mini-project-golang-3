package model

import "time"

type Book struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	ISBN      string
	Penulis   string
	Tahun     uint
	Judul     string
	Gambar    string
	Stok      uint
}
