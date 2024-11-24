package main

import (
	"fmt"
)

func main() {
	var x, y int

	// Input jumlah ikan dan kapasitas setiap wadah
	fmt.Print("Masukkan jumlah ikan (x) dan kapasitas wadah (y): ")
	fmt.Scan(&x, &y)

	if x <= 0 || y <= 0 || x > 1000 || y > 1000 {
		fmt.Println("Nilai x dan y harus antara 1 dan 1000.")
		return
	}

	// Input berat masing-masing ikan
	weights := make([]float64, x)
	fmt.Println("Masukkan berat masing-masing ikan:")
	for i := 0; i < x; i++ {
		fmt.Printf("Berat ikan ke-%d: ", i+1)
		fmt.Scan(&weights[i])
	}

	// Inisialisasi wadah
	numberOfContainers := (x + y - 1) / y // Pembulatan ke atas
	containers := make([]float64, numberOfContainers)

	// Mengisi wadah dengan berat ikan
	for i := 0; i < x; i++ {
		containerIndex := i / y
		containers[containerIndex] += weights[i]
	}

	// Hitung rata-rata berat per wadah
	totalWeight := 0.0
	for _, weight := range containers {
		totalWeight += weight
	}
	averageWeight := totalWeight / float64(numberOfContainers)

	// Output hasil
	fmt.Println("\nTotal berat di setiap wadah:")
	for i, weight := range containers {
		fmt.Printf("Wadah %d: %.2f\n", i+1, weight)
	}
	fmt.Printf("\nBerat rata-rata ikan di setiap wadah: %.2f\n", averageWeight)
}