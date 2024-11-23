package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int
	fmt.Print("Inputkan jumlah ikan yang  mau dijual (x) dan jumlah ikan per wadah (y): ")
	fmt.Scan(&x, &y)
	fmt.Println("Inputkan berat ikan:")
	berat := make([]float64, x)
	for i := 0; i < x; i++ {
		fmt.Scan(&berat[i])
	}


	wadahHitung := int(math.Ceil(float64(x) / float64(y)))
	wadahTotals := make([]float64, wadahHitung)

	for i := 0; i < x; i++ {
		wadahjumlah := i / y
		wadahTotals[wadahjumlah] += berat[i]
	}

	wadahRatarata := make([]float64, wadahHitung)
	for i := 0; i < wadahHitung; i++ {
		if (i+1)*y <= x {
			wadahRatarata[i] = wadahTotals[i] / float64(y)
		} else {
			wadahRatarata[i] = wadahTotals[i] / float64(x%y)
		}
	}


	fmt.Println("Total ikan di wadah:")
	for _, total := range wadahTotals {
		fmt.Printf("%.2f ", total)
	}
	fmt.Println()

	fmt.Println("Rata-rata berat ikan di setiap wadah:")
	for _, avg := range wadahRatarata {
		fmt.Printf("%.2f ", avg)
	}
	fmt.Println()
}