package main

import "fmt"

type arrBerat [1000]float64

func beratTerkecil(T arrBerat, n int) int {
	var idx int = 0
	var j int = 1

	for j < n {
		if T[idx] > T[j] {
			idx = j
		}
		j = j + 1
	}
	return idx
}

func beratTerbesar(T arrBerat, n int) int {
	var idx int = 0
	var j int = 1

	for j < n {
		if T[idx] < T[j] {
			idx = j
		}
		j = j + 1
	}
	return idx
}

func main() {
	var dataKelinci arrBerat
	var n int

	fmt.Print("Jumlah kelinci: ")
	fmt.Scan(&n)

	if n <= 0 || n > 1000 {
		fmt.Println("Jumlah kelinci harus antara 1-1000")
		return
	}

	fmt.Println("Masukkan berat:")
	for i := 0; i < n; i++ {
		fmt.Printf("Berat kelinci ke-%d: ", i+1)
		fmt.Scan(&dataKelinci[i])
	}

	idxMin := beratTerkecil(dataKelinci, n)
	idxMax := beratTerbesar(dataKelinci, n)

	fmt.Printf("\nBerat kelinci terkecil: %.2f kg\n", dataKelinci[idxMin])
	fmt.Printf("Berat kelinci terbesar: %.2f kg\n", dataKelinci[idxMax])
}
