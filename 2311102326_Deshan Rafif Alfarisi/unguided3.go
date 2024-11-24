package main

import "fmt"

const maxBalita = 100

// Tipe data array untuk menyimpan berat balita
type arrBalita [maxBalita]float64

// Fungsi untuk menghitung berat minimum dan maksimum
func hitungMinMax(arrBerat arrBalita, n int) (float64, float64) {
	var min, max float64

	// Inisialisasi nilai awal min dan max dengan elemen pertama
	min = arrBerat[0]
	max = arrBerat[0]

	// Iterasi untuk mencari nilai min dan max
	for i := 1; i < n; i++ {
		if arrBerat[i] < min {
			min = arrBerat[i]
		}
		if arrBerat[i] > max {
			max = arrBerat[i]
		}
	}

	return min, max
}

// Fungsi untuk menghitung rata-rata berat
func rerata(arrBerat arrBalita, n int) float64 {
	var total float64

	// Hitung total berat semua balita
	for i := 0; i < n; i++ {
		total += arrBerat[i]
	}

	// Hitung rata-rata
	return total / float64(n)
}

func main() {
	var n int
	var beratBalita arrBalita

	// Input jumlah balita
	fmt.Print("Masukkan jumlah balita: ")
	fmt.Scan(&n)

	// Validasi input
	if n <= 0 || n > maxBalita {
		fmt.Println("Jumlah balita harus antara 1 dan", maxBalita)
		return
	}

	// Input berat balita
	fmt.Println("Masukkan berat balita (dalam kg):")
	for i := 0; i < n; i++ {
		fmt.Scan(&beratBalita[i])
	}

	// Panggil fungsi untuk menghitung min, max, dan rata-rata
	min, max := hitungMinMax(beratBalita, n)
	rata := rerata(beratBalita, n)

	// Tampilkan hasil
	fmt.Printf("Berat balita terkecil: %.2f kg\n", min)
	fmt.Printf("Berat balita terbesar: %.2f kg\n", max)
	fmt.Printf("Rata-rata berat balita: %.2f kg\n", rata)
}
