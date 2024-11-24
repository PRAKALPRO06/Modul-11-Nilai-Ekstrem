package main

import (
	"fmt"
)

type arrBalita [100]float64

func hitungMinimalMaksimum_2311102033(arrBerat []float64, bMin, bMax *float64) {
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

func hitungratarata_2311102033(arrBerat []float64) float64 {
	var total float64
	for _, berat := range arrBerat {
		total += berat
	}
	return total / float64(len(arrBerat))
}

func main() {
	var n int
	fmt.Print("Masukkan banyak data berat balita: ")
	fmt.Scanln(&n)

	arrBerat := make([]float64, n)

	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan berat balita ke-%d: ", i+1)
		fmt.Scanln(&arrBerat[i])
	}

	var bMin, bMax float64
	hitungMinimalMaksimum_2311102033(arrBerat, &bMin, &bMax)
	rerata := hitungratarata_2311102033(arrBerat)

	fmt.Printf("\nBerat balita minimum: %.2f kg\n", bMin)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", bMax)
	fmt.Printf("Rerata berat balita: %.2f kg\n", rerata)
}
