package main

import (
	"fmt"
)

func main() {
	var beratKelinci [1000]float64
	var N int

	fmt.Print("Masukkan jumlah anak kelinci: ")
	fmt.Scan(&N)

	for i := 0; i < N; i++ {
		fmt.Printf("Masukkan berat kelinci ke-%d: ", i+1)
		fmt.Scan(&beratKelinci[i])
	}

	min := beratKelinci[0]
	max := beratKelinci[0]

	for i := 1; i < N; i++ {
		if beratKelinci[i] < min {
			min = beratKelinci[i]
		}
		if beratKelinci[i] > max {
			max = beratKelinci[i]
		}
	}

	fmt.Printf("Berat kelinci terkecil: %.2f\n", min)
	fmt.Printf("Berat kelinci terbesar: %.2f\n", max)
}
