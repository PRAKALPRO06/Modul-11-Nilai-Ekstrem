package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Print("Masukkan jumlah anak kelinci: ")
	fmt.Scan(&n)

	if n <= 0 || n > 1000 {
		fmt.Println("Jumlah anak kelinci harus antara 1 hingga 1000.")
		return
	}

	weights := make([]float64, n)
	fmt.Println("Masukkan berat masing-masing anak kelinci:")

	minWeight := math.MaxFloat64
	maxWeight := -math.MaxFloat64

	for i := 0; i < n; i++ {
		fmt.Printf("Berat kelinci ke-%d: ", i+1)
		fmt.Scan(&weights[i])
		if weights[i] < minWeight {
			minWeight = weights[i]
		}
		if weights[i] > maxWeight {
			maxWeight = weights[i]
		}
	}

	fmt.Printf("Berat terkecil: %.2f\n", minWeight)
	fmt.Printf("Berat terbesar: %.2f\n", maxWeight)
}
