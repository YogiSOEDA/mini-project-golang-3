package model

import (
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Book struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	ISBN      string         `json:"isbn"`
	Penulis   string         `json:"penulis"`
	Tahun     uint           `json:"tahun"`
	Judul     string         `json:"judul"`
	Gambar    string         `json:"gambar"`
	Stok      uint           `json:"stok"`
}

func (bk *Book) Create(db *gorm.DB) error {
	err := db.
		Model(Book{}).
		Create(&bk).
		Error
	if err != nil {
		return err
	}

	return nil
}

func (bk *Book) GetByID(db *gorm.DB) (Book, error) {
	res := Book{}

	err := db.
		Model(Book{}).
		Where("id = ?", bk.ID).
		Take(&res).
		Error
	if err != nil {
		return Book{}, err
	}

	return res, nil
}

func (bk *Book) GetAll(db *gorm.DB) ([]Book, error) {
	res := []Book{}

	err := db.
		Model(Book{}).
		Find(&res).
		Error

	if err != nil {
		return []Book{}, err
	}

	return res, nil
}

func (bk *Book) UpdatedOneByID(db *gorm.DB) error {
	err := db.
		Model(Book{}).
		Select("isbn", "penulis", "tahun", "judul", "gambar", "stok").
		Where("id = ?", bk.ID).
		Updates(map[string]interface{}{
			"isbn":    bk.ISBN,
			"penulis": bk.Penulis,
			"tahun":   bk.Tahun,
			"judul":   bk.Judul,
			"gambar":  bk.Gambar,
			"stok":    bk.Stok,
		}).
		Error

	if err != nil {
		return err
	}

	return nil
}

func (bk *Book) DeleteByID(db *gorm.DB) error {
	err := db.
		Model(Book{}).
		Where("id = ?", bk.ID).
		Delete(&bk).
		Error

	if err != nil {
		return err
	}

	return nil
}

func (bk *Book) SaveImport(db *gorm.DB) error {
	err := db.
		Model(Book{}).
		Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "id"}},
			DoUpdates: clause.AssignmentColumns([]string{"isbn", "penulis", "tahun", "judul","gambar", "stok"}),
		}).Create(&bk).
		Error
	if err != nil {
		return err
	}

	return nil
}
