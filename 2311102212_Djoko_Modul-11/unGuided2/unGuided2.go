package main

import (
	"fmt"
)

func hitungTotalBerat(weights []float64, x, y int) []float64 {
	totalWeights := make([]float64, (x+y-1)/y)
	var idx int
	for i := 0; i < len(totalWeights); i++ {
		for j := 0; j < y && idx < x; j++ {
			totalWeights[i] += weights[idx]
			idx++
		}
	}
	return totalWeights
}

func hitungRataRata(totalWeights []float64) float64 {
	var totalWeight float64
	for _, w := range totalWeights {
		totalWeight += w
	}
	return totalWeight / float64(len(totalWeights))
}

func main() {
	var x2311102212, y2311102212 int
	fmt.Print("Masukkan jumlah ikan (x) dan jumlah ikan per wadah (y): ")
	fmt.Scan(&x2311102212, &y2311102212)

	weights := make([]float64, x2311102212)
	fmt.Println("Masukkan berat ikan (pisahkan dengan spasi):")
	for i := 0; i < x2311102212; i++ {
		fmt.Scan(&weights[i])
	}

	totalWeights := hitungTotalBerat(weights, x2311102212, y2311102212)

	fmt.Println("Total berat ikan di setiap wadah:")
	for i, weight := range totalWeights {
		fmt.Printf("Wadah %d: %.2f\n", i+1, weight)
	}

	averageWeight := hitungRataRata(totalWeights)
	fmt.Printf("Rata-rata berat ikan di setiap wadah: %.2f\n", averageWeight)
}
