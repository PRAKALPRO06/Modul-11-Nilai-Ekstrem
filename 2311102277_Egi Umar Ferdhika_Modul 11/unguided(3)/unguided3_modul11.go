package main

import (
	"fmt"
)

type arrBalita [100]float64

func hitungMinMax(arrBerat arrBalita, n int) (min, max float64) {
	min = arrBerat[0]
	max = arrBerat[0]

	for i := 1; i < n; i++ {
		if arrBerat[i] < min {
			min = arrBerat[i]
		}
		if arrBerat[i] > max {
			max = arrBerat[i]
		}
	}
	return
}

func rerata(arrBerat arrBalita, n int) float64 {
	var total float64
	for i := 0; i < n; i++ {
		total += arrBerat[i]
	}
	return total / float64(n)
}

func main() {
	var dataBalita arrBalita
	var n int

	fmt.Print("Masukkan banyak data berat balita: ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan berat balita ke-%d: ", i+1)
		fmt.Scan(&dataBalita[i])
	}

	min, max := hitungMinMax(dataBalita, n)
	rata := rerata(dataBalita, n)

	fmt.Printf("Berat balita minimum: %.2f kg\n", min)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", max)
	fmt.Printf("Rerata berat balita: %.2f kg\n", rata)
}
