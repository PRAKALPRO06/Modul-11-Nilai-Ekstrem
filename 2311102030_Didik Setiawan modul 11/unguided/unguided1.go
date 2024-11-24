package main

import (
	"fmt"
)

func main() {
	var N int

	// Membaca jumlah anak kelinci
	fmt.Print("Masukkan jumlah anak kelinci (N): ")
	fmt.Scan(&N)

	// Validasi jumlah anak kelinci
	if N <= 0 || N > 1000 {
		fmt.Println("Jumlah anak kelinci harus antara 1 dan 1000.")
		return
	}

	// Inisialisasi berat terkecil dan terbesar
	var berat, terkecil, terbesar float64

	fmt.Println("Masukkan berat anak kelinci:")
	for i := 0; i < N; i++ {
		fmt.Printf("Berat ke-%d: ", i+1)
		fmt.Scan(&berat)

		// Validasi berat positif
		if berat <= 0 {
			fmt.Println("Berat harus angka positif.")
			return
		}

		// Inisialisasi nilai pertama sebagai acuan
		if i == 0 {
			terkecil, terbesar = berat, berat
		} else {
			if berat < terkecil {
				terkecil = berat
			}
			if berat > terbesar {
				terbesar = berat
			}
		}
	}

	// Output hasil
	fmt.Printf("Berat terkecil: %.2f\n", terkecil)
	fmt.Printf("Berat terbesar: %.2f\n", terbesar)
}
