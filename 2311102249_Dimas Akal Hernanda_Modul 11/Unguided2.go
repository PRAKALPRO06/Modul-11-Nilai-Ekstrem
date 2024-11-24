package main

import (
	"fmt"
)

func main() {
	var x, y int
	fmt.Print("Masukkan jumlah ikan (x) dan jumlah wadah (y): ")
	fmt.Scan(&x, &y)

	if x <= 0 || x > 1000 || y <= 0 || y > 1000 {
		fmt.Println("Jumlah ikan dan wadah harus antara 1 hingga 1000.")
		return
	}

	weights := make([]float64, x)
	fmt.Println("Masukkan berat masing-masing ikan:")
	for i := 0; i < x; i++ {
		fmt.Printf("Berat ikan ke-%d: ", i+1)
		fmt.Scan(&weights[i])
	}

	capacities := make([]int, y)
	fmt.Println("Masukkan kapasitas masing-masing wadah:")
	for i := 0; i < y; i++ {
		fmt.Printf("Kapasitas wadah ke-%d: ", i+1)
		fmt.Scan(&capacities[i])
	}

	totalWeights := make([]float64, y)
	index := 0

	for i := 0; i < y; i++ {
		for j := 0; j < capacities[i]; j++ {
			if index >= x {
				break
			}
			totalWeights[i] += weights[index]
			index++
		}
	}

	var sumWeights float64
	for _, weight := range totalWeights {
		sumWeights += weight
	}

	fmt.Println("Total berat di setiap wadah:")
	for i, weight := range totalWeights {
		fmt.Printf("Wadah ke-%d: %.2f\n", i+1, weight)
	}

	averageWeight := sumWeights / float64(y)
	fmt.Printf("Berat rata-rata di setiap wadah: %.2f\n", averageWeight)
}
