//2311102024_FahriRamadhan
package main

import (
	"fmt"
)

type arrBalita [100]float64

func hitungMinMax(arrBerat arrBalita, n int, bMin *float64, bMax *float64) {
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

func rata(arrBerat arrBalita, n int) float64 {
	var total float64 = 0.0

	for i := 0; i < n; i++ {
		total += arrBerat[i]
	}

	return total / float64(n)
}

func main() {
	var n int
	var berat arrBalita
	var bMin, bMax float64

	fmt.Print("Masukkan banyak data berat balita: ")
	fmt.Scan(&n)

	if n < 1 || n > 100 {
		fmt.Println("Jumlah balita harus antara 1 dan 100.")
		return
	}

	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan berat balita ke-%d: ", i+1)
		fmt.Scan(&berat[i])
	}

	hitungMinMax(berat, n, &bMin, &bMax)
	rerata := rata(berat, n)

	fmt.Printf("\nBerat balita minimum: %.2f kg\n", bMin)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", bMax)
	fmt.Printf("Rata-rata berat balita: %.2f kg\n", rerata)
}