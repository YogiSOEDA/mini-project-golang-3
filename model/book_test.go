package model_test

import (
	"fmt"
	"sekolahbeta/miniproject3/config"
	"sekolahbeta/miniproject3/model"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func Init() {
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println("env not found, using global env")
	}
	config.OpenDB()
}

func TestCreateBook(t *testing.T) {
	Init()

	bookData := model.Book{
		ISBN:    "9786022202028",
		Penulis: "Faza Meonk",
		Tahun:   2016,
		Judul:   "Ngampus!!! Buka-bukaan Aib Mahasiswa",
		Gambar:  "https://cdn.gramedia.com/uploads/items/9786022202028_ngampus_buka-bukaan_aib_mahasiswa.jpg",
		Stok:    3,
	}

	err := bookData.Create(config.Mysql.DB)
	assert.Nil(t, err)
}

func TestGetAll(t *testing.T) {
	Init()

	bookData := model.Book{
		ISBN:    "9786022202028",
		Penulis: "Faza Meonk",
		Tahun:   2016,
		Judul:   "Ngampus!!! Buka-bukaan Aib Mahasiswa",
		Gambar:  "https://cdn.gramedia.com/uploads/items/9786022202028_ngampus_buka-bukaan_aib_mahasiswa.jpg",
		Stok:    3,
	}

	err := bookData.Create(config.Mysql.DB)
	assert.Nil(t, err)

	res, err := bookData.GetAll(config.Mysql.DB)
	assert.Nil(t, err)
	assert.GreaterOrEqual(t, len(res), 1)

	fmt.Println(res)
}

func TestUpdate(t *testing.T) {
	Init()

	bookData := model.Book{
		ISBN:    "9786022202028",
		Penulis: "Faza Meonk",
		Tahun:   2016,
		Judul:   "Ngampus!!! Buka-bukaan Aib Mahasiswa",
		Gambar:  "https://cdn.gramedia.com/uploads/items/9786022202028_ngampus_buka-bukaan_aib_mahasiswa.jpg",
		Stok:    3,
	}

	err := bookData.Create(config.Mysql.DB)
	assert.Nil(t, err)

	bookData.Stok = 1

	err = bookData.UpdatedOneByID(config.Mysql.DB)
	assert.Nil(t, err)
}

func TestDeleteByID(t *testing.T) {
	Init()

	bookData := model.Book{
		ISBN:    "9786022202028",
		Penulis: "Faza Meonk",
		Tahun:   2016,
		Judul:   "Ngampus!!! Buka-bukaan Aib Mahasiswa",
		Gambar:  "https://cdn.gramedia.com/uploads/items/9786022202028_ngampus_buka-bukaan_aib_mahasiswa.jpg",
		Stok:    3,
	}

	err := bookData.Create(config.Mysql.DB)
	assert.Nil(t, err)

	fmt.Println(bookData.ID)

	err = bookData.DeleteByID(config.Mysql.DB)
	assert.Nil(t, err)
}