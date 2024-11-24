package main

import (
	"fmt"
)

func main() {
	var x, y int
	fmt.Print("Masukkan jumlah wadah (x): ")
	fmt.Scanln(&x)
	fmt.Print("Masukkan jumlah ikan per wadah (y): ")
	fmt.Scanln(&y)

	var berat_2311102033 []float64
	totalberat := make([]float64, x)
	totalFish := x * y

	fmt.Println("Masukkan berat ikan satu per satu:")
	for i := 0; i < totalFish; i++ {
		var weight float64
		fmt.Printf("Berat ikan ke-%d: ", i+1)
		fmt.Scanln(&weight)
		berat_2311102033 = append(berat_2311102033, weight)

		currentWadah := i / y
		totalberat[currentWadah] += weight
	}

	fmt.Println("\nTotal berat ikan di setiap wadah:")
	for i := 0; i < x; i++ {
		fmt.Printf("Wadah %d: %.2f\n", i+1, totalberat[i])
	}

	fmt.Println("\nBerat rata-rata ikan di setiap wadah:")
	for i := 0; i < x; i++ {
		average := totalberat[i] / float64(y)
		fmt.Printf("Wadah %d: %.2f\n", i+1, average)
	}
}
