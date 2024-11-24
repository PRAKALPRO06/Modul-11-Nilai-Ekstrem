// Meutya Azzahra Efendi
// 2311102166
// IF-11-06

package main

import "fmt"

type arrBalita [100]float64

// Fungsi untuk menghitung berat minimum dan maksimum
func hitungMinMax(arrBerat arrBalita, bMin, bMax *float64, n int) {
	*bMin = arrBerat[0]
	*bMax = arrBerat[0]

	for i := 1; i < n; i++ {
		if arrBerat[i] < *bMin {
			*bMin = arrBerat[i]
		}
		if arrBerat[i] > *bMax {
			*bMax = arrBerat[i]
		}
	}
}

// Fungsi untuk menghitung rata-rata berat balita
func rerata(arrBerat arrBalita, n int) float64 {
	total := 0.0
	for i := 0; i < n; i++ {
		total += arrBerat[i]
	}
	return total / float64(n)
}

func main() {
	var beratBalita arrBalita
	var n int

	// Input banyak data berat balita
	fmt.Print("Masukan banyak data berat balita: ")
	fmt.Scan(&n)

	// Validasi jumlah data
	if n < 1 || n > 100 {
		fmt.Println("Jumlah data harus antara 1 dan 100.")
		return
	}

	// Input berat balita
	for i := 0; i < n; i++ {
		fmt.Printf("Masukan berat balita ke-%d: ", i+1)
		fmt.Scan(&beratBalita[i])
	}

	// Variabel untuk menyimpan berat minimum dan maksimum
	var beratMin, beratMax float64

	// Hitung berat minimum dan maksimum
	hitungMinMax(beratBalita, &beratMin, &beratMax, n)

	// Hitung rata-rata berat balita
	rataRata := rerata(beratBalita, n)

	// Output hasil
	fmt.Printf("Berat balita minimum: %.2f kg\n", beratMin)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", beratMax)
	fmt.Printf("Rata-rata berat balita: %.2f kg\n", rataRata)
}