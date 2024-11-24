// 2311102037_BRIAN FARREL EVANDHIKA_IF 11 06
package main

import (
	"fmt"
	"math"
)

func main() {
	var N int
	var berat float64

	// Membaca jumlah anak kelinci
	fmt.Print("Masukkan jumlah anak kelinci (N): ")
	fmt.Scan(&N)

	// Validasi kapasitas array
	if N <= 0 || N > 1000 {
		fmt.Println("Jumlah anak kelinci harus antara 1 dan 1000.")
		return
	}

	beratKelinci := make([]float64, N)

	fmt.Println("Masukkan berat anak kelinci:")
	for i := 0; i < N; i++ {
		fmt.Printf("Berat ke-%d: ", i+1)
		fmt.Scan(&berat)
		beratKelinci[i] = berat
	}

	// Inisialisasi nilai terkecil dan terbesar
	terkecil := math.MaxFloat64
	terbesar := -math.MaxFloat64

	for _, b := range beratKelinci {
		if b < terkecil {
			terkecil = b
		}
		if b > terbesar {
			terbesar = b
		}
	}

	fmt.Printf("Berat terkecil: %.2f\n", terkecil)
	fmt.Printf("Berat terbesar: %.2f\n", terbesar)
}
