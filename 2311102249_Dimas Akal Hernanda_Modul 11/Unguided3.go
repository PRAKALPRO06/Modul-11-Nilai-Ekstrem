package main

import "fmt"

func main() {
	var jumlahBalita int
	var min, max, total float64

	fmt.Print("Masukkan jumlah balita: ")
	fmt.Scan(&jumlahBalita)

	if jumlahBalita <= 0 {
		fmt.Println("Jumlah balita harus lebih dari 0.")
		return
	}

	dataBalita := make([]float64, jumlahBalita)

	for i := 0; i < jumlahBalita; i++ {
		fmt.Printf("Masukkan berat balita ke-%d: ", i+1)
		fmt.Scan(&dataBalita[i])
	}

	min = dataBalita[0]
	max = dataBalita[0]

	for _, berat := range dataBalita {
		if berat < min {
			min = berat
		}
		if berat > max {
			max = berat
		}
		total += berat
	}

	rataRata := total / float64(jumlahBalita)

	fmt.Printf("\nBerat minimum: %.2f kg\n", min)
	fmt.Printf("Berat maksimum: %.2f kg\n", max)
	fmt.Printf("Rata-rata berat: %.2f kg\n", rataRata)
}
