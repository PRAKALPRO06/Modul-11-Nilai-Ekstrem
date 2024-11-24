package main

import "fmt"

type arrBalita [100]float64

func hitungMinMax(arrBerat arrBalita, bMin, bMax *float64) {
	*bMin = arrBerat[0]
	*bMax = arrBerat[0]

	for i := 1; i < len(arrBerat); i++ {
		if arrBerat[i] != 0 {
			if arrBerat[i] < *bMin {
				*bMin = arrBerat[i]
			}
			if arrBerat[i] > *bMax {
				*bMax = arrBerat[i]
			}
		}
	}
}

func rerata(arrBerat arrBalita) float64 {
	total := 0.0
	count := 0

	for i := 0; i < len(arrBerat); i++ {
		if arrBerat[i] != 0 {
			total += arrBerat[i]
			count++
		}
	}

	if count > 0 {
		return total / float64(count)
	}
	return 0
}

func main() {
	var dataBalita arrBalita
	var n int
	var bMin, bMax float64

	fmt.Print("Masukan jumlah data balita : ")
	fmt.Scan(&n)

	if n <= 0 || n > 100 {
		fmt.Println("Jumlah data harus antara 1-100")
		return
	}

	for i := 0; i < n; i++ {
		fmt.Printf("Berat balita ke-%d: ", i+1)
		fmt.Scan(&dataBalita[i])
	}

	hitungMinMax(dataBalita, &bMin, &bMax)
	rataRata := rerata(dataBalita)

	fmt.Printf("Berat balita minimum: %.2f kg\n", bMin)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", bMax)
	fmt.Printf("Rerata berat balita: %.2f kg\n", rataRata)
}
