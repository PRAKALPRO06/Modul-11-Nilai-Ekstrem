package main

import	"fmt"

type arrBalita [100]float64

func hitungMinMax(arrBerat arrBalita, n int, bMin, bMax *float64) {
	if n == 0 {
		return
	}
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

func rerata(arrBerat arrBalita, n int) float64 {
	if n == 0 {
		return 0.0
	}
	var total float64
	for i := 0; i < n; i++ {
		total += arrBerat[i]
	}
	return total / float64(n)
}

func main() {
	var n int
	fmt.Print("Masukkan jumlah balita: ")
	fmt.Scan(&n)

	if n > 100 {
		fmt.Println("Jumlah balita tidak boleh lebih dari 100.")
		return
	}

	var beratBalita arrBalita
	for i := 0; i < n; i++ {
	    fmt.Print("Masukkan berat balita ke-",i+1,": ")
		fmt.Scan(&beratBalita[i])
	}
	
	

	var minBerat, maxBerat float64
	hitungMinMax(beratBalita, n, &minBerat, &maxBerat)
	rataBerat := rerata(beratBalita, n)

	fmt.Printf("Berat balita terkecil: %.2f kg\n", minBerat)
	fmt.Printf("Berat balita terbesar: %.2f kg\n", maxBerat)
	fmt.Printf("Berat rata-rata balita: %.2f kg\n", rataBerat)
}
