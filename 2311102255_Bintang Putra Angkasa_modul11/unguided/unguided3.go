package main

import (
	"fmt"
)

type BeratBalita [100]float64 // Array untuk menyimpan berat balita

// Fungsi untuk menghitung berat balita minimum dan maksimum
func hitungMinMax(berat BeratBalita, jumlahBalita int, beratMin *float64, beratMax *float64) {
	// Inisialisasi beratMin dan beratMax dengan berat balita pertama
	*beratMin = berat[0]
	*beratMax = berat[0]

	// Menemukan berat balita minimum dan maksimum
	for i := 1; i < jumlahBalita; i++ {
		if berat[i] < *beratMin {
			*beratMin = berat[i]
		}
		if berat[i] > *beratMax {
			*beratMax = berat[i]
		}
	}
}

// Fungsi untuk menghitung rata-rata berat balita
func hitungRataRata(berat BeratBalita, jumlahBalita int) float64 {
	var totalBerat float64 = 0.0
	for i := 0; i < jumlahBalita; i++ {
		totalBerat += berat[i]
	}
	return totalBerat / float64(jumlahBalita)
}

func main() {
	var jumlahBalita int
	var beratBalita BeratBalita
	var beratMin, beratMax float64

	// Input jumlah balita
	fmt.Print("Masukkan jumlah balita: ")
	fmt.Scan(&jumlahBalita)

	// Validasi jumlah balita
	if jumlahBalita < 1 || jumlahBalita > 100 {
		fmt.Println("Jumlah balita harus antara 1 dan 100.")
		return
	}

	// Input berat balita
	for i := 0; i < jumlahBalita; i++ {
		fmt.Printf("Masukkan berat balita ke-%d: ", i+1)
		fmt.Scan(&beratBalita[i])
	}

	// Menghitung berat balita minimum dan maksimum
	hitungMinMax(beratBalita, jumlahBalita, &beratMin, &beratMax)

	// Menghitung rata-rata berat balita
	rataRata := hitungRataRata(beratBalita, jumlahBalita)

	// Menampilkan hasil
	fmt.Printf("\nBerat balita minimum: %.2f kg\n", beratMin)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", beratMax)
	fmt.Printf("Rata-rata berat balita: %.2f kg\n", rataRata)
}
