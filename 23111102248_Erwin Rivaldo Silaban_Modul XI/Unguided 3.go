package main

//Erwin Rivaldo Silaban
//2311102248

import "fmt"

type arrBalita [100]float64

func hitungMinMax(arrBerat arrBalita, beratMin, beratMax *float64, n int) {
	*beratMin = arrBerat[0]
	*beratMax = arrBerat[0]

	for i := 1; i < n; i++ {
		if *beratMin > arrBerat[i] {
			*beratMin = arrBerat[i]
		}
		if *beratMax < arrBerat[i] {
			*beratMax = arrBerat[i]
		}
	}
}

func rerata(arrBerat arrBalita, n int) float64 {
	var rata float64 = 0
	bagi := 0
	for i := 0; i < n; i++ {
		rata += arrBerat[i]
		bagi++
	}

	return rata / float64(bagi)
}

func main() {
	var beratMin, beratMax float64
	var arrBerat arrBalita
	var n int

	fmt.Print("Masukkan banyak data berat balita : ")
	fmt.Scan(&n)

	if n < 1 || n > 100 {
		fmt.Println("Jumlah data berat balita harus antara 1 dan 100.")
		return
	}

	for i := 0; i < n; i++ {
		fmt.Print("Masukkan berat balita ke-", i+1, ": ")
		fmt.Scan(&arrBerat[i])
	}

	hitungMinMax(arrBerat, &beratMin, &beratMax, n)
	rataRata := rerata(arrBerat, n)

	fmt.Printf("Berat balita minimum : %.2f Kg\n", beratMin)
	fmt.Printf("Berat balita maksimum : %.2f Kg\n", beratMax)
	fmt.Printf("Rerata berat balita : %.2f Kg\n", rataRata)

}
