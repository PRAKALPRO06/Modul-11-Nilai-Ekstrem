package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Print("Masukkan jumlah anak kelinci (N): ")
	fmt.Scan(&n)

	if n <= 0 || n > 1000 {
		fmt.Println("Jumlah anak kelinci harus di antara 1 hingga 1000.")
		return
	}

	beratKelinci := make([]float64, n)
	fmt.Println("Masukkan berat masing-masing anak kelinci:")

	for i := 0; i < n; i++ {
		fmt.Printf("Berat ke-%d: ", i+1)
		fmt.Scan(&beratKelinci[i])
	}

	beratTerkecil := math.MaxFloat64
	BeratMaksimal := -math.MaxFloat64

	for _, weight := range beratKelinci {
		if weight < beratTerkecil {
			beratTerkecil = weight
		}
		if weight > BeratMaksimal {
			BeratMaksimal = weight
		}
	}

	fmt.Printf("Berat terkecil: %.2f\n", beratTerkecil)
	fmt.Printf("Berat terbesar: %.2f\n", BeratMaksimal)
}
