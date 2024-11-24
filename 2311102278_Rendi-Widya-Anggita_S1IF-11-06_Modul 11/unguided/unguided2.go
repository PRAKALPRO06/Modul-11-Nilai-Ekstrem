package main

import (
	"fmt"
)

func main() {
	// Step 1: Penjelasan Program
	fmt.Println("Program ini mendistribusikan berat ikan ke dalam wadah.")
	fmt.Println("Masukkan jumlah ikan (x) dan jumlah wadah (y), diikuti berat ikan masing-masing.")

	// Step 2: Input
	var x, y int
	fmt.Print("Masukkan jumlah ikan (x) : ")
	fmt.Scan(&x)
	fmt.Print("Masukkan jumlah wadah (y): ")
	fmt.Scan(&y)

	// Membuat slice untuk menyimpan berat ikan
	fishWeights := make([]float64, x)
	fmt.Printf("Masukkan %d berat ikan (dipisahkan dengan spasi): ", x)
	for i := 0; i < x; i++ {
		fmt.Scan(&fishWeights[i])
	}

	// Step 3: Distribusi ikan ke dalam wadah
	buckets := make([][]float64, y)
	for i, weight := range fishWeights {
		bucketIndex := i % y // Distribusi secara bergilir (round-robin)
		buckets[bucketIndex] = append(buckets[bucketIndex], weight)
	}

	// Step 4: Perhitungan total dan rata-rata berat di setiap wadah
	totalWeights := make([]float64, y)
	averageWeights := make([]float64, y)

	for i := 0; i < y; i++ {
		var totalWeight float64
		for _, weight := range buckets[i] {
			totalWeight += weight
		}
		totalWeights[i] = totalWeight
		if len(buckets[i]) > 0 {
			averageWeights[i] = totalWeight / float64(len(buckets[i]))
		}
	}

	// Step 5: Output Hasil
	fmt.Println("\nHasil :")
	fmt.Println("Total berat di setiap wadah:")
	for i, total := range totalWeights {
		fmt.Printf("Wadah %d: %.2f\n", i+1, total)
	}

	fmt.Println("\nRata-rata berat ikan di setiap wadah:")
	for i, avg := range averageWeights {
		fmt.Printf("Wadah %d: %.2f\n", i+1, avg)
	}
}
