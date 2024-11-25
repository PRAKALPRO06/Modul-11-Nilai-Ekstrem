// Rakha Yudhistira_2311102010_IF-11-06

package main

import (
	"fmt"
)

type arrBalita [100]float64

func hitungMinMax_2311102010(arr arrBalita, n int, min, max *float64) {
	*min = arr[0]
	*max = arr[0]
	for i := 1; i < n; i++ {
		if arr[i] < *min {
			*min = arr[i]
		}
		if arr[i] > *max {
			*max = arr[i]
		}
	}
}
func rataRata(arr arrBalita, n int) float64 {
	var total float64 = 0
	for i := 0; i < n; i++ {
		total += arr[i]
	}
	return total / float64(n)
}
func main() {
	var n int
	var berat arrBalita
	var min, max float64
	fmt.Print("Masukkan banyak data berat balita: ")
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan berat balita ke-%d: ", i+1)
		fmt.Scan(&berat[i])
	}
	hitungMinMax_2311102010(berat, n, &min, &max)
	rata := rataRata(berat, n)
	// Output hasil
	fmt.Printf("Berat balita minimum: %.2f kg\n", min)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", max)
	fmt.Printf("Rata-rata berat balita: %.2f kg\n", rata)
}
