package main

import "fmt"

type arrFloat [1000]float64

func hitungTotalPerWadah(berat []float64, x, y int) []float64 {
	totalPerWadah := make([]float64, (x+y-1)/y)
	index := 0

	for i := 0; i < len(totalPerWadah); i++ {
		total := 0.0
		for j := 0; j < y && index < x; j++ {
			total += berat[index]
			index++
		}
		totalPerWadah[i] = total
	}

	return totalPerWadah
}

func hitungRataRata(totalPerWadah []float64) float64 {
	total := 0.0
	for _, nilai := range totalPerWadah {
		total += nilai
	}
	return total / float64(len(totalPerWadah))
}

func main() {
	var x, y int
	var beratIkan arrFloat

	fmt.Print("Masukkan jumlah ikan dan kapasitas wadah: ")
	fmt.Scan(&x, &y)

	if x <= 0 || y <= 0 || x > 1000 {
		fmt.Println("Input tidak valid")
		return
	}

	fmt.Println("Masukkan berat masing-masing ikan:")
	for i := 0; i < x; i++ {
		fmt.Printf("Berat ikan ke-%d (kg): ", i+1)
		fmt.Scan(&beratIkan[i])
	}

	totalPerWadah := hitungTotalPerWadah(beratIkan[:x], x, y)
	rataRata := hitungRataRata(totalPerWadah)

	fmt.Println("\nTotal berat ikan per wadah:")
	for i, total := range totalPerWadah {
		fmt.Printf("Wadah %d: %.2f kg\n", i+1, total)
	}
	fmt.Printf("Rata-rata berat ikan per wadah: %.2f kg\n", rataRata)
}
