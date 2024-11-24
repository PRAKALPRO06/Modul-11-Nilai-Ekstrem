package main

import "fmt"

type arrBalita [100]float64

func main() {
	var dataBalita arrBalita
	var jumlahBalita int
	var min, max, total float64

	fmt.Print("Masukkan jumlah balita: ")
	fmt.Scan(&jumlahBalita)

	for i := 0; i < jumlahBalita; i++ {
		fmt.Printf("Masukkan berat balita ke-%d: ", i+1)
		fmt.Scan(&dataBalita[i])
	}

	min = dataBalita[0]
	max = dataBalita[0]

	for i := 0; i < jumlahBalita; i++ {
		if dataBalita[i] < min {
			min = dataBalita[i]
		}

		if dataBalita[i] > max {
			max = dataBalita[i]
		}

		total = total + dataBalita[i]
	}

	rataRata := total / float64(jumlahBalita)

	fmt.Printf("\nBerat minimum: %.2f kg\n", min)
	fmt.Printf("Berat maximum: %.2f kg\n", max)
	fmt.Printf("Rata-rata berat: %.2f kg\n", rataRata)
}
