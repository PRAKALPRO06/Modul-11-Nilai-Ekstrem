package main

import (
	"fmt"
	"math"
)

type arrBalita [100]float64

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

func rata_rata(arrBerat arrBalita, n int) float64 {
	var total float64
	for i := 0; i < n; i++ {
		total += arrBerat[i]
	}
	return total / float64(n)
}

func main() {
	var n int
	var berat arrBalita

	fmt.Print("Masukkan banyak data berat balita: ")
	fmt.Scan(&n)

	if n <= 0 || n > 100 {
		fmt.Println("Jumlah data berat balita harus antara 1 hingga 100.")
		return
	}

	fmt.Println("Masukkan berat balita:")
	for i := 0; i < n; i++ {
		fmt.Printf("Masukan berat balita ke-%d: ", i+1)
		fmt.Scan(&berat[i])
	}

	var bMin, bMax float64
	hitungMinMax(berat, n, &bMin, &bMax)
	rata := rata_rata(berat, n)

	fmt.Printf("Berat balita minimum: %.2f kg\n", bMin)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", bMax)
	fmt.Printf("Berat balita rata-rata: %.2f kg\n", rata)
}
