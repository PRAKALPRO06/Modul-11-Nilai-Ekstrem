package main

import "fmt"

// Definisi tipe array untuk data berat balita
type arrBalita [100]float64

// Fungsi untuk mencari berat minimum dan maksimum balita
func hitungMinMax(arrBerat arrBalita, bMin, bMax *float64) {
	*bMin = arrBerat[0]
	*bMax = arrBerat[0]

	for i := 1; i < 4; i++ {
		if arrBerat[i] < *bMin {
			*bMin = arrBerat[i]
		}
		if arrBerat[i] > *bMax {
			*bMax = arrBerat[i]
		}
	}
}

// Fungsi untuk menghitung rerata berat balita
func rerata(arrBerat arrBalita) float64 {
	var total float64
	for i := 0; i < 4; i++ {
		total += arrBerat[i]
	}
	return total / 4
}

func main() {
	var beratBalita arrBalita
	var bMin, bMax float64
	var n int

	// Input data berat balita
	fmt.Printf("Masukan banyak data berat balita : ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Printf("Masukan berat balita ke-%d: ", i+1)
		fmt.Scan(&beratBalita[i])
	}

	// Mencari berat minimum dan maksimum
	hitungMinMax(beratBalita, &bMin, &bMax)

	// Menghitung rerata
	avg := rerata(beratBalita)

	// Menampilkan hasil
	fmt.Printf("Berat balita minimum: %.2f kg\n", bMin)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", bMax)
	fmt.Printf("Rerata berat balita: %.2f kg\n", avg)
}
