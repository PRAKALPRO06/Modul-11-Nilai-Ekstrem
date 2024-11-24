package main

import (
	"fmt"
)

func main() {
	var x, y int

	fmt.Print("Masukkan jumlah wadah (x): ")
	fmt.Scan(&x)
	fmt.Print("Masukkan jumlah ikan per wadah (y): ")
	fmt.Scan(&y)

	if x < 1 || x > 1000 || y < 1 {
		fmt.Println("Nilai x harus antara 1-1000 dan y harus lebih dari 0.")
		return
	}

	ikan := make([]float64, x*y)
	fmt.Println("Masukkan berat ikan satu per satu:")
	for i := 0; i < x*y; i++ {
		fmt.Printf("Berat ikan ke-%d: ", i+1)
		fmt.Scan(&ikan[i])
	}

	totalBerat := make([]float64, x)
	rataRata := make([]float64, x)

	for i := 0; i < x; i++ {
		sum := 0.0
		for j := 0; j < y; j++ {
			sum += ikan[i*y+j]
		}
		totalBerat[i] = sum
		rataRata[i] = sum / float64(y)
	}

	fmt.Println("\nTotal berat ikan di setiap wadah:")
	for i := 0; i < x; i++ {
		fmt.Printf("Wadah %d: %.2f\n", i+1, totalBerat[i])
	}

	fmt.Println("\nBerat rata-rata ikan di setiap wadah:")
	for i := 0; i < x; i++ {
		fmt.Printf("Wadah %d: %.2f\n", i+1, rataRata[i])
	}
}
