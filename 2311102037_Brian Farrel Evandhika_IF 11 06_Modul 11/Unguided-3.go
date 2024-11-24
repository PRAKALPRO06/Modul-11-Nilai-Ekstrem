// 2311102037_BRIAN FARREL EVANDHIKA_IF 11 06
package main

import (
	"fmt"
)

type arrBalita [100]float64

// hitungMinMax menghitung berat minimum dan maksimum dalam array
func hitungMinMax(arrBerat arrBalita, n int, bMin, bMax *float64) {
	*bMin = arrBerat[0]
	*bMax = arrBerat[0]

	for i := 1; i < n; i++ {
		if arrBerat[i] < *bMin {
			*bMin = arrBerat[i]
		}
		if arrBerat[i] > *bMax {
			*bMax = arrBerat[i]
		}
	}
}

// rerata menghitung dan mengembalikan rata-rata berat balita dalam array
func rerata(arrBerat arrBalita, n int) float64 {
	total := 0.0
	for i := 0; i < n; i++ {
		total += arrBerat[i]
	}
	return total / float64(n)
}

func main() {
	var n int
	var arrBerat arrBalita
	var bMin, bMax float64

	// Input jumlah data balita
	fmt.Print("Masukan banyak data berat balita: ")
	fmt.Scan(&n)

	// Input berat masing-masing balita
	for i := 0; i < n; i++ {
		fmt.Printf("Masukan berat balita ke-%d: ", i+1)
		fmt.Scan(&arrBerat[i])
	}

	// Hitung minimum dan maksimum
	hitungMinMax(arrBerat, n, &bMin, &bMax)

	// Hitung rata-rata
	rata := rerata(arrBerat, n)

	// Output hasil
	fmt.Printf("Berat balita minimum: %.2f kg\n", bMin)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", bMax)
	fmt.Printf("Rerata berat balita: %.2f kg\n", rata)
}
