package main

import (
	"fmt"
	"os"
	"sekolahbeta/miniproject3/config"
	"sekolahbeta/miniproject3/controllers"

	"github.com/joho/godotenv"
)

func Init() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("env not found, using global env")
	}
	config.OpenDB()
}

func main() {
	Init()

	var pilihanMenu int

	fmt.Println("===========================================")
	fmt.Println("Aplikasi Manajemen Daftar Buku Perpustakaan")
	fmt.Println("===========================================")
	fmt.Println("Silahkan Pilih Menu : ")
	fmt.Println("1. Tambah Buku")
	fmt.Println("2. Lihat Buku")
	fmt.Println("3. Hapus Buku")
	fmt.Println("4. Edit Buku")
	fmt.Println("5. Print Buku")
	fmt.Println("6. Keluar")
	fmt.Println("===========================================")

	fmt.Print("Masukkan Pilihan : ")
	_, err := fmt.Scanln(&pilihanMenu)
	if err != nil {
		fmt.Println("Terjadi error :", err)
	}

	switch pilihanMenu {
	case 1:
		controllers.TambahBuku()
	case 2:
		controllers.LihatBuku()
	case 3:
		controllers.HapusBuku()
	case 4:
		controllers.EditBuku()
	case 6:
		os.Exit(0)
	}
	main()
}
