package main

import (
	"fmt"
	"math"
)

func main() {
	var jumlahIkan, jumlahPerWadah int
	fmt.Print("Inputkan jumlah ikan yang mau dijual (x) dan jumlah ikan per wadah (y): ")
	fmt.Scan(&jumlahIkan, &jumlahPerWadah)

	fmt.Println("Inputkan berat ikan:")
	beratIkan := inputBerat(jumlahIkan)

	totalWadah := hitungWadah(jumlahIkan, jumlahPerWadah)
	totalPerWadah := hitungTotalPerWadah(beratIkan, jumlahPerWadah, totalWadah)
	rataRataPerWadah := hitungRataRata(totalPerWadah, jumlahIkan, jumlahPerWadah)

	tampilkanHasil("Total ikan di wadah:", totalPerWadah)
	tampilkanHasil("Rata-rata berat ikan di setiap wadah:", rataRataPerWadah)
}

func inputBerat(jumlah int) []float64 {
	berat := make([]float64, jumlah)
	for i := 0; i < jumlah; i++ {
		fmt.Scan(&berat[i])
	}
	return berat
}

func hitungWadah(jumlahIkan, jumlahPerWadah int) int {
	return int(math.Ceil(float64(jumlahIkan) / float64(jumlahPerWadah)))
}

func hitungTotalPerWadah(berat []float64, jumlahPerWadah, totalWadah int) []float64 {
	totals := make([]float64, totalWadah)
	for i, b := range berat {
		indexWadah := i / jumlahPerWadah
		totals[indexWadah] += b
	}
	return totals
}

func hitungRataRata(totalPerWadah []float64, jumlahIkan, jumlahPerWadah int) []float64 {
	rataRata := make([]float64, len(totalPerWadah))
	for i := range totalPerWadah {
		if (i+1)*jumlahPerWadah <= jumlahIkan {
			rataRata[i] = totalPerWadah[i] / float64(jumlahPerWadah)
		} else {
			rataRata[i] = totalPerWadah[i] / float64(jumlahIkan%jumlahPerWadah)
		}
	}
	return rataRata
}

func tampilkanHasil(label string, data []float64) {
	fmt.Println(label)
	for _, value := range data {
		fmt.Printf("%.2f ", value)
	}
	fmt.Println()
}
