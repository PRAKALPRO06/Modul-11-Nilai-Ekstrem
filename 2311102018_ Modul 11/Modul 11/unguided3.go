//2311102018
package main

import "fmt"


func findMin(weights []float64, n int) float64 {
	min := weights[0]
	for i := 1; i < n; i++ {
		if weights[i] < min {
			min = weights[i]
		}
	}
	return min
}
func findMax(weights []float64, n int) float64 {
	max := weights[0]
	for i := 1; i < n; i++ {
		if weights[i] > max {
			max = weights[i]
		}
	}
	return max
}
func calculateAverage(weights []float64, n int) float64 {
	var total float64
	for i := 0; i < n; i++ {
		total += weights[i]
	}
	return total / float64(n)
}
func main() {
	var n int
	fmt.Print("Masukkan jumlah balita: ")
	fmt.Scan(&n)
	if n < 1 {
		fmt.Println("Jumlah balita harus lebih dari 0.")
		return
	}
	weights := make([]float64, n)
	fmt.Println("Masukkan berat balita (dalam kg):")
	for i := 0; i < n; i++ {
		fmt.Printf("Berat balita ke-%d: ", i+1)
		fmt.Scan(&weights[i])
		if weights[i] <= 0 {
			fmt.Println("Berat balita harus lebih dari 0.")
			i--
		}
	}
	minWeight := findMin(weights, n)
	maxWeight := findMax(weights, n)
	averageWeight := calculateAverage(weights, n)
	fmt.Println("\nHasil Analisis Berat Balita:")
	fmt.Printf("Berat terkecil: %.2f kg\n", minWeight)
	fmt.Printf("Berat terbesar: %.2f kg\n", maxWeight)
	fmt.Printf("Rata-rata berat: %.2f kg\n", averageWeight)
}
