package controllers

import (
	"bufio"
	"fmt"
	"os"
	"sekolahbeta/miniproject3/config"
	"sekolahbeta/miniproject3/model"
	"strings"
	"sync"
)

func TambahBuku() {

	inputanUser := bufio.NewReader(os.Stdin)
	draftBuku := []model.Book{}

	var (
		isbn        string
		tahunTerbit uint
		gambarBuku  string
		stokBuku    uint
	)

	fmt.Println("===========================================")
	fmt.Println("Tambah Buku")
	fmt.Println("===========================================")

	for {
		fmt.Print("Silahkan Masukkan ISBN : ")
		_, err := fmt.Scanln(&isbn)
		if err != nil {
			fmt.Println("Terjadi Error : ", err)
			return
		}

		fmt.Print("Silahkan Masukkan Penulis Buku : ")
		penulisBuku, err := inputanUser.ReadString('\r')
		if err != nil {
			fmt.Println("Terjadi Error : ", err)
			return
		}
		penulisBuku = strings.Replace(penulisBuku, "\n", "", 1)
		penulisBuku = strings.Replace(penulisBuku, "\r", "", 1)

		fmt.Print("Silahkan Masukkan Tahun Terbit Buku : ")
		_, err = fmt.Scanln(&tahunTerbit)
		if err != nil {
			fmt.Println("Terjadi Error : ", err)
			return
		}

		fmt.Print("Silahkan Masukkan Judul Buku : ")
		judulBuku, err := inputanUser.ReadString('\r')
		if err != nil {
			fmt.Println("Terjadi Error : ", err)
			return
		}
		judulBuku = strings.Replace(judulBuku, "\n", "", 1)
		judulBuku = strings.Replace(judulBuku, "\r", "", 1)

		fmt.Print("Silahkan Masukkan Gambar Buku : ")
		_, err = fmt.Scanln(&gambarBuku)
		if err != nil {
			fmt.Println("Terjadi Error : ", err)
			return
		}

		fmt.Print("Silahkan Masukkan Stok Buku : ")
		_, err = fmt.Scanln(&stokBuku)
		if err != nil {
			fmt.Println("Terjadi Error : ", err)
			return
		}

		draftBuku = append(draftBuku, model.Book{
			ISBN:    isbn,
			Penulis: penulisBuku,
			Tahun:   tahunTerbit,
			Judul:   judulBuku,
			Gambar:  gambarBuku,
			Stok:    stokBuku,
		})

		var pilihanMenu = 0
		fmt.Println("Ketik 1 untuk tambah pesanan, ketik 0 untuk keluar")
		_, err = fmt.Scanln(&pilihanMenu)
		if err != nil {
			fmt.Println(err)
			return
		}

		if pilihanMenu == 0 {
			break
		}
	}

	fmt.Println("Menambah Pesanan")

	ch := make(chan model.Book)
	wg := sync.WaitGroup{}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go simpanBuku(ch, &wg)
	}

	for _, buku := range draftBuku {
		ch <- buku
	}

	close(ch)

	wg.Wait()

	fmt.Println("Buku Berhasil Ditambah!")
}

func simpanBuku(ch <-chan model.Book, wg *sync.WaitGroup) {
	for buku := range ch {
		err := buku.Create(config.Mysql.DB)
		if err != nil {
			fmt.Println(err)
		}
	}

	wg.Done()
}

func LihatBuku() {
	listBuku := model.Book{}

	fmt.Println("===========================================")
	fmt.Println("Lihat Buku")
	fmt.Println("===========================================")
	fmt.Println("Memuat data ...")

	res, err := listBuku.GetAll(config.Mysql.DB)
	if err != nil {
		fmt.Println(err)
	}

	for urutan, buku := range res {
		fmt.Printf("%d. Kode Buku : %d, ISBN : %s, Penulis Buku : %s, Tahun Terbit : %d, Judul Buku : %s, Gambar Buku : %s, Stok Buku : %d \n",
			urutan+1,
			buku.ID,
			buku.ISBN,
			buku.Penulis,
			buku.Tahun,
			buku.Judul,
			buku.Gambar,
			buku.Stok,
		)
	}
}
