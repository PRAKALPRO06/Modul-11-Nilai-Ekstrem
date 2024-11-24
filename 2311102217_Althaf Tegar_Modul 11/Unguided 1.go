package main

import (
	"fmt"
	"math"
)

func main() {
	const maxKapasitas = 1000
	var berat2217 [maxKapasitas]float64
	var n217 int

	fmt.Print("Masukkan jumlah anak kelinci: ")
	_, err := fmt.Scan(&n217)
	if err != nil || n217 < 1 || n217 > maxKapasitas {
		fmt.Println("Input tidak valid. Pastikan jumlah anak kelinci antara 1 hingga 1000.")
		return
	}

	fmt.Print("Masukkan berat anak kelinci:")
	for i := 0; i < n217; i++ {
		_, err := fmt.Scan(&berat2217[i])
		if err != nil {
			fmt.Println("Input tidak valid. Pastikan memasukkan bilangan riil untuk berat.")
			return
		}
	}

	// Inisialisasi nilai terendah dan tertinggi
	minWeight := math.Inf(1)
	maxWeight := math.Inf(-1)

	// Mencari berat terkecil dan terbesar
	for i := 0; i < n217; i++ {
		if berat2217[i] < minWeight {
			minWeight = berat2217[i]
		}
		if berat2217[i] > maxWeight {
			maxWeight = berat2217[i]
		}
	}
	fmt.Printf("Berat anak kelinci terkecil: %.2f\n", minWeight)
	fmt.Printf("Berat anak kelinci terbesar: %.2f\n", maxWeight)
}
