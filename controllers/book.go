package controllers

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"sekolahbeta/miniproject3/config"
	"sekolahbeta/miniproject3/model"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/go-pdf/fpdf"
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

func HapusBuku() {
	var kodeBuku uint

	fmt.Println("===========================================")
	fmt.Println("Hapus Buku")
	fmt.Println("===========================================")
	LihatBuku()
	fmt.Println("===========================================")

	fmt.Print("Masukkan Kode Buku : ")
	_, err := fmt.Scanln(&kodeBuku)
	if err != nil {
		fmt.Println("Terjadi Error : ", err)
		return
	}

	buku := model.Book{
		ID: kodeBuku,
	}

	err = buku.DeleteByID(config.Mysql.DB)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Buku Berhasil Dihapus!")
}

func EditBuku() {
	inputanUser := bufio.NewReader(os.Stdin)

	var (
		kodeBuku    uint
		isbn        string
		tahunTerbit uint
		gambarBuku  string
		stokBuku    uint
	)

	fmt.Println("===========================================")
	fmt.Println("Edit Buku")
	fmt.Println("===========================================")
	LihatBuku()
	fmt.Println("===========================================")

	fmt.Print("Masukkan Kode Buku : ")
	_, err := fmt.Scanln(&kodeBuku)
	if err != nil {
		fmt.Println("Terjadi Error : ", err)
		return
	}

	fmt.Print("Silahkan Masukkan ISBN : ")
	_, err = fmt.Scanln(&isbn)
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

	buku := model.Book{
		ID:      kodeBuku,
		ISBN:    isbn,
		Penulis: penulisBuku,
		Tahun:   tahunTerbit,
		Judul:   judulBuku,
		Gambar:  gambarBuku,
		Stok:    stokBuku,
	}

	err = buku.UpdatedOneByID(config.Mysql.DB)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Data Buku Berhasil Diubah!")
}

func PrintPdfBuku() {
	listBuku := model.Book{}

	var pilihanMenu int

	fmt.Println("===========================================")
	fmt.Println("Print Buku")
	fmt.Println("===========================================")
	LihatBuku()
	fmt.Println("===========================================")
	fmt.Println("Silahkan Pilih :")
	fmt.Println("1. Print Salah Satu Buku")
	fmt.Println("2. Print Semua Buku")
	fmt.Println("===========================================")

	fmt.Print("Masukkan Pilihan : ")
	_, err := fmt.Scanln(&pilihanMenu)
	if err != nil {
		fmt.Println(err)
	}

	_ = os.Mkdir("pdf", 0777)

	pdf := fpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)
	pdf.SetLeftMargin(10)
	pdf.SetRightMargin(10)

	switch pilihanMenu {
	case 1:
		var kodeBuku uint

		fmt.Print("Masukkan Kode Buku : ")
		_, err := fmt.Scanln(&kodeBuku)
		if err != nil {
			fmt.Println("Terjadi Error : ", err)
			return
		}

		listBuku.ID = kodeBuku

		buku, err := listBuku.GetByID(config.Mysql.DB)
		if err != nil {
			fmt.Println(err)
		}

		bukuText := fmt.Sprintf(
			"Buku :\nKode Buku : %d\nISBN : %s\nPenulis Buku : %s\nTahun Terbit : %d\nJudul Buku : %s\nGambar Buku : %s\nStok Buku : %d\n",
			buku.ID,
			buku.ISBN,
			buku.Penulis,
			buku.Tahun,
			buku.Judul,
			buku.Gambar,
			buku.Stok,
		)

		pdf.MultiCell(0, 10, bukuText, "0", "L", false)

		err = pdf.OutputFileAndClose(
			fmt.Sprintf("pdf/book-%d.pdf", kodeBuku),
		)
		if err != nil {
			fmt.Println(err)
		}

	case 2:
		res, err := listBuku.GetAll(config.Mysql.DB)
		if err != nil {
			fmt.Println(err)
		}

		for i, buku := range res {
			bukuText := fmt.Sprintf(
				"Buku #%d:\nKode Buku : %d\nISBN : %s\nPenulis Buku : %s\nTahun Terbit : %d\nJudul Buku : %s\nGambar Buku : %s\nStok Buku : %d\n",
				i+1,
				buku.ID,
				buku.ISBN,
				buku.Penulis,
				buku.Tahun,
				buku.Judul,
				buku.Gambar,
				buku.Stok,
			)

			pdf.MultiCell(0, 10, bukuText, "0", "L", false)
			pdf.Ln(5)
		}

		err = pdf.OutputFileAndClose(
			fmt.Sprintf("pdf/daftar_buku_%s.pdf", time.Now().Format("2006-01-02-15-04-05")))

		if err != nil {
			fmt.Println(err)
		}
	}
}

func ImportData() {
	var direktori string

	fmt.Println("===========================================")
	fmt.Println("Import Data Buku dari File CSV")
	fmt.Println("===========================================")

	fmt.Print("Silahkan Masukkan Path atau Lokasi File CSV : ")
	_, err := fmt.Scanln(&direktori)
	if err != nil {
		fmt.Println("Terjadi Error : ", err)
	}

	file, err := openFile(direktori)
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	csvChan, err := loadFile(file)
	if err != nil {
		fmt.Println(err)
	}

	jmlGoroutine := 5

	var bookChanTemp []<-chan model.Book

	for i := 0; i < jmlGoroutine; i++ {
		bookChanTemp = append(bookChanTemp, processConverStruct(csvChan))
	}

	mergeCh := appendBooks(bookChanTemp...)

	var books []model.Book

	for ch := range mergeCh {
		books = append(books, ch)
	}

	ch := make(chan model.Book)
	wg := sync.WaitGroup{}

	for i := 0; i < jmlGoroutine; i++ {
		wg.Add(1)
		go simpanImportBuku(ch, &wg)
	}

	for _, book := range books {
		ch <- book
	}

	close(ch)
	wg.Wait()

	fmt.Println("Import Data Selesai!")

}

func openFile(path string) (*os.File, error) {
	return os.Open(path)
}

func loadFile(file *os.File) (<-chan []string, error) {
	bookChan := make(chan []string)
	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		return bookChan, err
	}

	go func() {
		for i, book := range records {
			if i == 0 {
				continue
			}
			bookChan <- book
		}

		close(bookChan)
	}()

	return bookChan, nil
}

func processConverStruct(csvChan <-chan []string) <-chan model.Book {
	booksChan := make(chan model.Book)

	go func() {
		for book := range csvChan {
			id, err := strconv.ParseUint(book[0],10,64)
			if err != nil {
				fmt.Println(err)
			}

			tahun, err := strconv.ParseUint(book[3],10,64)
			if err != nil {
				fmt.Println(err)
			}

			stok, err := strconv.ParseUint(book[6],10,64)
			if err != nil {
				fmt.Println(err)
			}

			booksChan <- model.Book{
				ID:   uint(id),
				ISBN: book[1],
				Penulis: book[2],
				Tahun: uint(tahun),
				Judul: book[4],
				Gambar: book[5],
				Stok: uint(stok),
			}
		}

		close(booksChan)
	}()

	return booksChan
}

func appendBooks(bookChanMany ...<-chan model.Book) <-chan model.Book {
	wg := sync.WaitGroup{}

	mergedChan := make(chan model.Book)

	wg.Add(len(bookChanMany))
	for _, ch := range bookChanMany {
		go func(ch <- chan model.Book)  {
			for books := range ch {
				mergedChan <- books
			}
			wg.Done()
		}(ch)
	}

	go func() {
		wg.Wait()
		close(mergedChan)
	}()

	return mergedChan
}

func simpanImportBuku(ch <-chan model.Book, wg *sync.WaitGroup)  {
	for book := range ch {
		err := book.SaveImport(config.Mysql.DB)
		if err != nil {
			fmt.Println(err)
		}
	}

	wg.Done()
}
