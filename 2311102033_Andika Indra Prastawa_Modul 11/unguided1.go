package main

import (
	"fmt"
)

func cariberatdarikecilkebesar_2311102033(berat []float64, n int) (float64, float64) {
	min := berat[0]
	max := berat[0]

	for i := 1; i < n; i++ {
		if berat[i] < min {
			min = berat[i]
		}
		if berat[i] > max {
			max = berat[i]
		}
	}

	return min, max
}

func main() {
	const maxKapasitas = 1000
	var berat [maxKapasitas]float64
	var n int

	fmt.Print("Masukkan jumlah anak kelinci yang akan di timbang : ")
	fmt.Scan(&n)

	if n < 1 || n > maxKapasitas {
		fmt.Println("Jumlah anak kelinci harus antara 1 dan 1000.")
		return
	}

	fmt.Println("Masukkan berat anak kelinci:")
	for i := 0; i < n; i++ {
		fmt.Printf("Berat anak kelinci ke-%d: ", i+1)
		fmt.Scan(&berat[i])
	}

	terkecil, terbesar := cariberatdarikecilkebesar_2311102033(berat[:n], n)

	fmt.Printf("\nBerat terkecil: %.2f\n", terkecil)
	fmt.Printf("Berat terbesar: %.2f\n", terbesar)
}
