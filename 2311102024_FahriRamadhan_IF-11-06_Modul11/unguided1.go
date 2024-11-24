//2311102024_FahriRamadhan
package main

import (
	"fmt"
)

func cariBerat(berat []float64, n int) (float64, float64) {
	terkecil := berat[0]
	terbesar := berat[0]

	for i := 1; i < n; i++ {
		if berat[i] < terkecil {
			terkecil = berat[i]
		}
		if berat[i] > terbesar {
			terbesar = berat[i]
		}
	}

	return terkecil, terbesar
}

func main() {
	var n int

	fmt.Print("Masukkan jumlah anak kelinci yang akan ditimbang: ")
	fmt.Scan(&n)

	if n < 1 || n > 1000 {
		fmt.Println("Jumlah anak kelinci harus antara 1 dan 1000.")
		return
	}

	berat := make([]float64, n)

	fmt.Println("Masukkan berat anak kelinci:")
	for i := 0; i < n; i++ {
		fmt.Printf("Berat anak kelinci ke-%d: ", i+1)
		fmt.Scan(&berat[i])
	}

	terkecil, terbesar := cariBerat(berat, n)

	fmt.Printf("\nBerat kelinci terkecil adalah: %.2f\n", terkecil)
	fmt.Printf("Berat kelinci terbesar adalah: %.2f\n", terbesar)
}