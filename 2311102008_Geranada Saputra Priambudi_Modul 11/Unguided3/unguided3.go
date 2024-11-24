package main

import "fmt"

type arrBalita [100]float64

func hitungMinMax(arrBerat arrBalita, jumlahBalita int, bMin, bMax *float64) {
	*bMin = arrBerat[0]
	*bMax = arrBerat[0]
	for i := 0; i < jumlahBalita; i++ {
		if arrBerat[i] < *bMin {
			*bMin = arrBerat[i]
		}
		if arrBerat[i] > *bMax {
			*bMax = arrBerat[i]
		}
	}
}

func rerata(arrBerat arrBalita, jumlahBalita int) float64 {
	var total float64
	for n := 0; n < jumlahBalita; n++ {
		total += arrBerat[n]
	}
	return total / float64(jumlahBalita)
}

func main() {
	var jumlahBalita int
	fmt.Print("Masukan banyak data berat balita : ")
	fmt.Scan(&jumlahBalita)
	var arrBerat arrBalita
	for n := 0; n < jumlahBalita; n++ {
		fmt.Printf("Masukan berat balita ke-%d: ", n+1)
		fmt.Scan(&arrBerat[n])
	}

	var bMin, bMax float64
	hitungMinMax(arrBerat, jumlahBalita, &bMin, &bMax)
	rerataBerat := rerata(arrBerat, jumlahBalita)
	fmt.Printf("Berat balita minimum: %.2f kg\n", bMin)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", bMax)
	fmt.Printf("Rerata berat balita: %.2f kg\n", rerataBerat)
}
