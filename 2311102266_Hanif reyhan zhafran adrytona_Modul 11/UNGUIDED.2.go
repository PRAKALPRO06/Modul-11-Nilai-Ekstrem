// 2311102266_Hanif Reyhan Zhafran Arytona

package main

import "fmt"

func main() {

	fmt.Println("masukan jumlah ikan dan kapasitas wadah pada (x)&(y) :")

	var x, y int
	fmt.Print("Masukkan jumlah ikan (x) dan jumlah wadah (y): ")
	fmt.Scan(&x, &y)

	fishWeights := make([]float64, x)
	fmt.Printf("Masukkan %d berat ikan (pisahkan dengan spasi): ", x)
	for i := 0; i < x; i++ {
		fmt.Scan(&fishWeights[i])
	}

	buckets := make([][]float64, y)
	for i, weight := range fishWeights {
		bucketIndex := i % y
		buckets[bucketIndex] = append(buckets[bucketIndex], weight)
	}

	totalWeights := make([]float64, y)
	averageWeights := make([]float64, y)
	for i := 0; i < y; i++ {
		var totalWeight float64
		for _, weight := range buckets[i] {
			totalWeight += weight
		}
		totalWeights[i] = totalWeight
		if len(buckets[i]) > 0 {
			averageWeights[i] = totalWeight / float64(len(buckets[i]))
		}
	}

	fmt.Println("\nHasil Distribusi:")
	fmt.Println("Total berat di setiap wadah:")
	for i, total := range totalWeights {
		fmt.Printf("Wadah %d: %.2f\n", i+1, total)
	}
	fmt.Println("\nRata-rata berat ikan di setiap wadah:")
	for i, avg := range averageWeights {
		fmt.Printf("Wadah %d: %.2f\n", i+1, avg)
	}
}
