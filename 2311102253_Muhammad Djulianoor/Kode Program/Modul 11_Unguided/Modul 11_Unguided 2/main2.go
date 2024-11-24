package main

import (
	"fmt"
)

func bacaBeratIkan(x int) []float64 {
	berat := make([]float64, x)
	fmt.Println("Masukkan berat ikan (dipisahkan dengan spasi):")
	for i := 0; i < x; i++ {
		fmt.Scan(&berat[i])
	}
	return berat
}

func kelompokkanIkan(berat []float64, x, y int) ([][]float64, []float64) {
	var containers [][]float64
	var totalBerat []float64

	for i := 0; i < x; i++ {
		if i%y == 0 {
			containers = append(containers, []float64{})
		}
		lastIndex := len(containers) - 1
		containers[lastIndex] = append(containers[lastIndex], berat[i])
	}

	var sum_berat float64
	for _, container := range containers {
		var BeratTotal float64
		for _, weight := range container {
			BeratTotal += weight
		}
		totalBerat = append(totalBerat, BeratTotal)
		sum_berat += BeratTotal
	}

	return containers, totalBerat
}

func hitungRataRata(totalBerat []float64) float64 {
	var sum float64
	for _, weight := range totalBerat {
		sum += weight
	}
	return sum / float64(len(totalBerat))
}

func main() {
	var x, y int
	fmt.Print("Masukkan jumlah ikan (x) dan jumlah ikan per wadah (y): ")
	fmt.Scan(&x, &y)

	if x <= 0 || y <= 0 || x > 1000 {
		fmt.Println("Nilai x harus antara 1 hingga 1000, dan y harus lebih dari 0.")
		return
	}

	berat := bacaBeratIkan(x)

	_, totalBerat := kelompokkanIkan(berat, x, y)

	fmt.Println("Total berat ikan di setiap wadah:")
	for _, BeratTotal := range totalBerat {
		fmt.Printf("%.2f ", BeratTotal)
	}
	fmt.Println()

	averageWeight := hitungRataRata(totalBerat)
	fmt.Printf("Berat rata-rata ikan di setiap wadah: %.2f\n", averageWeight)
}
