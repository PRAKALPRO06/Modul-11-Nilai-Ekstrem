package main

import "fmt"

func main() {
	var x, y int
	var weights [1000]float64
	var totalWeights []float64

	fmt.Printf("Masukkan jumlah ikan: ")
	fmt.Scan(&x)

	fmt.Printf("Masukkan kapasitas wadah: ")
	fmt.Scan(&y)

	if x < 1 || x > 1000 || y < 1 {
		fmt.Println("Jumlah ikan harus antara 1 dan 1000, kapasitas wadah minimal harus 1.")
		return
	}

	for i := 0; i < x; i++ {
		fmt.Printf("Masukkan berat ikan ke-%d: ", i+1)
		fmt.Scan(&weights[i])
	}

	var sum, count, totalWeightsum float64
	for i := 0; i < x; i++ {
		sum += weights[i]
		count++
		if int(count) == y || i == x-1 {
			totalWeights = append(totalWeights, sum)
			totalWeightsum += sum
			sum = 0
			count = 0
		}
	}

	fmt.Println("\nTotal berat di setiap wadah: ")
	for i, weigth := range totalWeights {
		fmt.Printf("Wadah ke-%d: %.2f kg\n", i+1, weigth)
	}

	averangeweight := totalWeightsum / float64(len(totalWeights))
	fmt.Printf("\nBerat rata-rata per wadah: %.2f kg\n", averangeweight)
}
