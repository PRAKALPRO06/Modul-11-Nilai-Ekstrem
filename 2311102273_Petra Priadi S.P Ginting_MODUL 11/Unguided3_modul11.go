package main

import (
	"fmt"
)

type arrBalita [100]float64

func hitungMinMax(arrBerat arrBalita, x int, bMin, bMax *float64) {
	*bMin = arrBerat[0]
	*bMax = arrBerat[0]
	for i := 1; i < x; i++ {
		if arrBerat[i] < *bMin {
			*bMin = arrBerat[i]
		}
		if arrBerat[i] > *bMax {
			*bMax = arrBerat[i]
		}
	}
}

func rerata(arrBerat arrBalita, x int) float64 {
	var total float64
	for i := 0; i < x; i++ {
		total += arrBerat[i]
	}
	return total / float64(x)
}

func main() {
	var x int
	var berat arrBalita
	var bMin, bMax float64

	fmt.Print("Masukan banyak data berat balita: ")
	fmt.Scan(&x)

	for i := 0; i < x; i++ {
		fmt.Printf("Masukan berat balita ke-%d: ", i+1)
		fmt.Scan(&berat[i])
	}

	hitungMinMax(berat, x, &bMin, &bMax)
	rataRata := rerata(berat, x)

	fmt.Printf("\nBerat balita minimum: %.2f kg\n", bMin)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", bMax)
	fmt.Printf("Rerata berat balita: %.2f kg\n", rataRata)
}
