package main

import (
	"fmt"
)

// Mendefinisikan tipe data arrBalita sebagai array float64 dengan ukuran maksimum 100
type arrBalita [100]float64

// Fungsi untuk menghitung berat minimum dan maksimum
func hitungMinMax(arrBerat []float64, bMin, bMax *float64) {
	*bMin = arrBerat[0]
	*bMax = arrBerat[0]

	for _, berat := range arrBerat {
		if berat < *bMin {
			*bMin = berat
		}
		if berat > *bMax {
			*bMax = berat
		}
	}
}

// Fungsi untuk menghitung rata-rata berat
func hitungRerata(arrBerat []float64) float64 {
	var total float64
	for _, berat := range arrBerat {
		total += berat
	}
	return total / float64(len(arrBerat))
}

func main() {
	var n int
	var berat arrBalita

	// Input: jumlah data berat balita
	fmt.Print("Masukan banyak data berat balita: ")
	fmt.Scanln(&n)

	// Input: berat masing-masing balita
	for i := 0; i < n; i++ {
		fmt.Printf("Masukan berat balita ke-%d: ", i+1)
		fmt.Scanln(&berat[i])
	}

	arrBerat := berat[:n]

	// Menghitung berat minimum, maksimum, dan rata-rata
	var bMin, bMax float64
	hitungMinMax(arrBerat, &bMin, &bMax)
	rerata := hitungRerata(arrBerat)

	// Output hasil
	fmt.Printf("Berat balita minimum: %.2f kg\n", bMin)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", bMax)
	fmt.Printf("Rata-rata berat balita: %.2f kg\n", rerata)
}
