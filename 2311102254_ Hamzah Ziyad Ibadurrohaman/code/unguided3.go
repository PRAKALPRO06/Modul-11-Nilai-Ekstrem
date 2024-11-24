package main

import "fmt"

type arrbalita [100]float64

func rerata(arrberat arrbalita, n int) float64 {
	var total float64 = 0
	for i := 0; i < n; i++ {
		total += arrberat[i]
	}
	return total / float64(n)
}

func hitungmindanmaks(arrberat arrbalita, n int, bmin, bmaks *float64) {
	*bmin, *bmaks = arrberat[0], arrberat[0]
	for i := 1; i < n; i++ {
		if arrberat[i] < *bmin {
			*bmin = arrberat[i]
		}
		if arrberat[i] > *bmaks {
			*bmaks = arrberat[i]
		}
	}
}

func main() {
	var databerat arrbalita
	var n int
	var bmin, bmaks float64

	fmt.Print("Input data banyak berat balita: ")
	fmt.Scan(&n)

	if n < 1 || n > 100 {
		fmt.Println("data harus dari 1 hingga 100. kalau kelebihan kapasitas kepenuhan.")
		return
	}
	for i := 0; i < n; i++ {
		fmt.Printf("berat balita ke-%d: ", i+1)
		fmt.Scan(&databerat[i])
	}
	hitungmindanmaks(databerat, n, &bmin, &bmaks)
	rata := rerata(databerat, n)

	fmt.Printf("berat balita minimum: %.2f\n", bmin)
	fmt.Printf("berat balita maksimum: %.2f\n", bmaks)
	fmt.Printf("rata-rata berat: %.2f\n", rata)
}
