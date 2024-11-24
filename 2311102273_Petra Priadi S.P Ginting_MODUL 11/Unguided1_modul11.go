package main

import (
	"fmt"
)

func main() {
	var berat [1000]float64

	var p int
	fmt.Print("Masukkan jumlah anak kelinci: ")
	fmt.Scan(&p)

	if p <= 0 || p > 1000 {
		fmt.Println("Jumlah anak kelinci harus antara 1 dan 1000.")
		return
	}

	fmt.Printf("Masukkan berat %d anak kelinci:\n", p)
	for i := 0; i < p; i++ {
		fmt.Scan(&berat[i])
	}

	minberat := berat[0]
	maxberat := berat[0]

	for i := 1; i < p; i++ {
		if berat[i] < minberat {
			minberat = berat[i]
		}
		if berat[i] > maxberat {
			maxberat = berat[i]
		}
	}

	fmt.Printf("Berat terkecil: %.2f\n", minberat)
	fmt.Printf("Berat terbesar: %.2f\n", maxberat)
}
