package main

import "fmt"

const maxKelinci = 1000

func main() {
	var n int
	var berat [maxKelinci]float64
	var terkecil, terbesar float64

	// Meminta input jumlah kelinci
	fmt.Print("Masukkan jumlah anak kelinci: ")
	fmt.Scan(&n)

	// Validasi input
	if n <= 0 || n > maxKelinci {
		fmt.Println("Jumlah kelinci harus antara 1 dan", maxKelinci)
		return
	}

	// Memasukkan data berat kelinci
	fmt.Println("Masukkan berat masing-masing kelinci (dalam kg):")
	for i := 0; i < n; i++ {
		fmt.Scan(&berat[i])
	}

	// Inisialisasi nilai terkecil dan terbesar dengan berat kelinci pertama
	terkecil = berat[0]
	terbesar = berat[0]

	// Mencari nilai terkecil dan terbesar
	for i := 1; i < n; i++ {
		if berat[i] < terkecil {
			terkecil = berat[i]
		}
		if berat[i] > terbesar {
			terbesar = berat[i]
		}
	}

	// Menampilkan hasil
	fmt.Printf("Berat kelinci terkecil: %.2f kg\n", terkecil)
	fmt.Printf("Berat kelinci terbesar: %.2f kg\n", terbesar)
}
