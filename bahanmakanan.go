package main

import (
	"fmt"
)

const NMAX int = 20

type BahanMakanan struct {
	Nama string
	Stok int
}

type stokrumah [NMAX]BahanMakanan

var dataBarang stokrumah 
var jumlahBahan int       

// simpanData 
func simpanData(nama string) int {
	//Is : terdeteksi nama sebagai string
	//Fs : Data nama disimpan
	for i := 0; i < jumlahBahan; i++ {
		if dataBarang[i].Nama == nama {
			return i
		}
	}
	return -1
}

// tambahBahan 
func tambahBahan(n int) {
	//Is : Terdeteksi n sebagai Integer
	//Fs : Stok bahan ditambahkan
	fmt.Println("Masukkan jumlah bahan yang akan ditambah")
	if n > NMAX {
		n = NMAX
	}

	for i := 0; i < n; i++ {
		if jumlahBahan >= NMAX {
			fmt.Println("Kapasitas maksimum tercapai. Tidak dapat menambah bahan lagi.")
			break
		}

		var nama string
		var jumlah int
		fmt.Printf("Masukkan Nama barang dan Jumlah Stoknya %d: ", i+1)
		fmt.Scan(&nama, &jumlah)

		idx := simpanData(nama)

		if idx != -1 {
			// Bahan sudah ada, tambahkan stoknya
			dataBarang[idx].Stok += jumlah
			fmt.Printf("Stok %s diperbarui. Stok baru: %d\n", nama, dataBarang[idx].Stok)
		} else {
			// Bahan belum ada, tambahkan sebagai bahan baru
			dataBarang[jumlahBahan] = BahanMakanan{Nama: nama, Stok: jumlah}
			jumlahBahan++
		}
	}
}

// kurangiBahan 
func kurangiBahan(nama string, jumlah int) {
	//Is : Terdeteksi jumlah sebagai integer
	//Fs : Bahan stok dikurangi
	idx := simpanData(nama)

	if idx != -1 {
		if dataBarang[idx].Stok >= jumlah {
			//jika stok barang yang dikurangi masih tersisa
			dataBarang[idx].Stok -= jumlah
			fmt.Printf("Stok %s dikurangi. Stok baru: %d\n", nama, dataBarang[idx].Stok)
			if dataBarang[idx].Stok <= 0 {
				//jika stok barang habis
				fmt.Printf("%s habis. Menghapus dari daftar.\n", nama)
				dataBarang[idx] = dataBarang[jumlahBahan-1]
				jumlahBahan--
			}
		} else {
			//jika stok barang yang ingin dikurangi tidak cukup
			fmt.Printf("Stok %s hanya %d, tidak cukup untuk mengurangi %d.\n", nama, dataBarang[idx].Stok, jumlah)
		}
	} else {
		//jika nama barang tidak ditemukan
		fmt.Printf("%s tidak ditemukan dalam daftar.\n", nama)
	}
}

// tampilStok 
func tampilStok() {
	//is : User Memilih case function 
	//fs : Stok ditampilkan
	if jumlahBahan == 0 {
		//jika belum ada bahan yang ditambah atau list kosong
		fmt.Println("Daftar bahan makanan kosong.")
		return
	}
	fmt.Printf("%-20s %s\n", "Nama", "Stok")
	fmt.Println("-------------------- ----")
	for i := 0; i < jumlahBahan; i++ {
		fmt.Printf("%-20s %d\n", dataBarang[i].Nama, dataBarang[i].Stok)
	}
}

// cekstok 
func cekstok(batasMinimum int) {
	//is : terdeteksi batasMinimum sebagai integer
	//fs : Terdeteksi stok yang harus di restok
	fmt.Println("\n--- Cek Restock ---")
	//true dan false sebagai penanda jika tidak ada maka fungsi dibawah if !foundRestock tidak bisa berjalan
	//atau jika ditambah else akan terprint menjadi banyak
	foundRestock := false
	for i := 0; i < jumlahBahan; i++ {
		if dataBarang[i].Stok < batasMinimum {
			fmt.Printf("Stok %s hanya %d - perlu restock!\n", dataBarang[i].Nama, dataBarang[i].Stok)
			foundRestock = true
		}
	}
	if !foundRestock {
		fmt.Println("Tidak ada bahan yang perlu restock.")
	}
}

// insertionSort DENGAN PARAMETER A *stokrumah DAN n int

func insertionSort(A *stokrumah, n int) { 
	//is : Terdeteksi Array n bilangan bulat
	//fs : Array A terurut secara ascending
	var i, pass int
	var temp BahanMakanan

	pass = 1
	for pass <= n-1 {
		i = pass
		temp = A[pass] // Mengakses melalui pointer A

		for i > 0 && temp.Stok < A[i-1].Stok { // Mengakses melalui pointer A
			A[i] = A[i-1] // Mengakses melalui pointer A
			i = i - 1
		}
		A[i] = temp // Mengakses melalui pointer A
		pass = pass + 1
	}
	fmt.Println("Daftar bahan telah diurutkan berdasarkan Nama.")
}

func main() {
	var pilihan int

	for {
		fmt.Println("\n====== Manajemen Stok Bahan Makanan ======")
		fmt.Println("1. Tambah Bahan")
		fmt.Println("2. Kurangi Bahan")
		fmt.Println("3. Tampilkan Stok (Tersortir)")
		fmt.Println("4. Cek Restock")
		fmt.Println("5. Keluar")
		fmt.Print("Pilih menu (1-5): ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			var n int
			fmt.Print("Masukkan berapa bahan yang ingin ditambah: ")
			fmt.Scan(&n)
			tambahBahan(n)

		case 2:
			var nama string
			var jumlah int
			fmt.Print("Masukkan nama bahan yang akan dikurangi: ")
			fmt.Scan(&nama)
			fmt.Print("Jumlah yang diambil: ")
			fmt.Scan(&jumlah)
			kurangiBahan(nama, jumlah)

		case 3:
			// Panggil insertionSort dengan meneruskan alamat dari dataBarang
			// dan nilai dari jumlahBahan
			insertionSort(&dataBarang, jumlahBahan)
			tampilStok() // tampilStok tetap menggunakan global dataBarang

		case 4:
			var batas int
			fmt.Print("Masukkan batas minimum stok: ")
			fmt.Scan(&batas)
			cekstok(batas)

		case 5:
			fmt.Println("Terima kasih!")
			return

		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}