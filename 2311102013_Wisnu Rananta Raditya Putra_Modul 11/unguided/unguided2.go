//WISNU RANANTA RADITYA PUTRA (2311102013) IF-11-06
package main

import "fmt"

const maxCapacity = 1000

func calculateWeight(x int, y int, weights_2311102013 [maxCapacity]float64) ([]float64, float64) {
	numContainers := (x + y - 1) / y
	containerWeights := make([]float64, numContainers)

	for i := 0; i < x; i++ {
		containerIndex := i / y
		containerWeights[containerIndex] += weights_2311102013[i]
	}

	totalWeight := 0.0
	for _, weight := range containerWeights {
		totalWeight += weight
	}
	averageWeight := totalWeight / float64(numContainers)

	return containerWeights, averageWeight
}

func main() {
	var x, y int
	var weights [maxCapacity]float64

	fmt.Print("Masukkan jumlah ikan (x) dan kapasitas wadah (y): ")
	fmt.Scan(&x, &y)

	if x > maxCapacity {
		fmt.Printf("Jumlah ikan melebihi kapasitas maksimum %d\n", maxCapacity)
		return
	}

	fmt.Printf("Masukkan berat ikan sebanyak %d:\n", x)
	for i := 0; i < x; i++ {
		fmt.Printf("Berat ikan ke-%d: ", i+1)
		fmt.Scan(&weights[i])
	}

	containerWeights, averageWeight := calculateWeight(x, y, weights)

	fmt.Println("\nHASIL PERHITUNGAN IKAN")
	fmt.Println("Total berat ikan di setiap wadah adalah sebagai berikut:")
	for i, weight := range containerWeights {
		fmt.Printf("Wadah %d: %.2f kg\n", i+1, weight)
	}

	fmt.Printf("\nBerat rata-rata ikan di setiap wadah adalah %.2f kg\n", averageWeight)
}