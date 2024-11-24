// 2311102266_Hanif Reyhan Zhafran Arytona

package main

import (
	"fmt"
	"math"
)

func main() {
	var N int
	var berat float64

	fmt.Print("Masukkan jumlah anak kelinci : ")
	fmt.Scan(&N)

	if N <= 0 || N > 1000 {
		fmt.Println("Jumlah anak kelinci harus antara 1 dan 1000.")
		return
	}
	beratKelinci := make([]float64, N)
	fmt.Println("berat anak kelinci:")
	for i := 0; i < N; i++ {
		fmt.Printf("%d: ", i+1)
		fmt.Scan(&berat)
		beratKelinci[i] = berat
	}

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
