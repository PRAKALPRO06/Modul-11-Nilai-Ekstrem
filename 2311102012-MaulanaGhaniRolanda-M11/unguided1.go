//2311102012
package main

import (
	"fmt"
	"math"
)

func main() {
	// Deklarasi variabel
	var n int
	fmt.Print("Masukkan jumlah anak kelinci: ")
	fmt.Scan(&n)

	if n <= 0 || n > 1000 {
		fmt.Println("Jumlah anak kelinci harus antara 1 dan 1000.")
		return
	}

	berat := make([]float64, n) // Array untuk menampung berat anak kelinci
	fmt.Println("Masukkan berat anak kelinci satu per satu:")

	for i := 0; i < n; i++ {
		fmt.Printf("Berat anak kelinci ke-%d: ", i+1)
		fmt.Scan(&berat[i])
	}

	// Inisialisasi nilai min dan max
	minBerat := math.MaxFloat64
	maxBerat := -math.MaxFloat64

	// Cari berat terkecil dan terbesar
	for _, b := range berat {
		if b < minBerat {
			minBerat = b
		}
		if b > maxBerat {
			maxBerat = b
		}
	}

	// Output hasil
	fmt.Printf("Berat terkecil: %.2f\n", minBerat)
	fmt.Printf("Berat terbesar: %.2f\n", maxBerat)
}
