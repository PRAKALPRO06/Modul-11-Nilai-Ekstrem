package main

import (
	"fmt"
)

// hitungStatistik menghitung berat minimum, maksimum, dan rata-rata
func hitungStatistik(arrBerat []float64) (float64, float64, float64) {
	if len(arrBerat) == 0 {
		return 0, 0, 0
	}

	bMin := arrBerat[0]
	bMax := arrBerat[0]
	total := 0.0

	for _, berat := range arrBerat {
		if berat < bMin {
			bMin = berat
		}
		if berat > bMax {
			bMax = berat
		}
		total += berat
	}

	rata := total / float64(len(arrBerat))
	return bMin, bMax, rata
}

func main() {
	var n int

	// Input jumlah data balita
	fmt.Print("Masukkan banyak data berat balita: ")
	fmt.Scan(&n)

	// Validasi jumlah balita
	if n <= 0 {
		fmt.Println("Jumlah data balita harus lebih dari 0.")
		return
	}

	// Input berat masing-masing balita
	arrBerat := make([]float64, n)
	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan berat balita ke-%d: ", i+1)
		fmt.Scan(&arrBerat[i])
	}

	// Hitung statistik
	bMin, bMax, rata := hitungStatistik(arrBerat)

	// Output hasil
	fmt.Printf("Berat balita minimum: %.2f kg\n", bMin)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", bMax)
	fmt.Printf("Rerata berat balita: %.2f kg\n", rata)
}
