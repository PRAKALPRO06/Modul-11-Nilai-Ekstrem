// 2311102012
package main

import (
	"fmt"
	"math"
)

// Definisi array balita
type arrBalita [100]float64

// Fungsi untuk menghitung minimum dan maksimum
func hitungMinMax(arrBerat arrBalita, n int, bMin, bMax *float64) {
	*bMin = math.MaxFloat64
	*bMax = -math.MaxFloat64
	for i := 0; i < n; i++ {
		if arrBerat[i] < *bMin {
			*bMin = arrBerat[i]
		}
		if arrBerat[i] > *bMax {
			*bMax = arrBerat[i]
		}
	}
}

// Fungsi untuk menghitung rata-rata
func rata(arrBerat arrBalita, n int) float64 {
	var total float64
	for i := 0; i < n; i++ {
		total += arrBerat[i]
	}
	return total / float64(n)
}

func main() {
	var n int
	var arr arrBalita
	var bMin, bMax float64

	// Input jumlah data balita
	fmt.Print("Masukkan banyak data berat balita: ")
	fmt.Scan(&n)

	// Validasi input
	if n <= 0 || n > 100 {
		fmt.Println("Jumlah data harus antara 1 dan 100.")
		return
	}

	// Input data berat balita
	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan berat balita ke-%d: ", i+1)
		fmt.Scan(&arr[i])
	}

	// Hitung nilai minimum, maksimum, dan rata-rata
	hitungMinMax(arr, n, &bMin, &bMax)
	rataRata := rata(arr, n)

	// Output hasil
	fmt.Printf("Berat balita minimum: %.2f kg\n", bMin)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", bMax)
	fmt.Printf("Rata-rata berat balita: %.2f kg\n", rataRata)
}
