package main

import "fmt"

type arrBalita [100]float64

func hitungMinMax(arrBerat arrBalita, bMin, bMax *float64, n int)  {
	*bMin = arrBerat[0]
	*bMax = arrBerat[0]
	
	for i := 1; i < n; i++ {
		if *bMin > arrBerat[i] {
			*bMin = arrBerat[i]
		}
		if *bMax < arrBerat[i] {
			*bMax = arrBerat[i]
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

func main()  {
	var bMin, bMax float64
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

	hitungMinMax(arrBerat, &bMin, &bMax, n)
	rataRata := rerata(arrBerat, n)

    fmt.Printf("Berat balita minimum : %.2f Kg\n", bMin)
    fmt.Printf("Berat balita maksimum : %.2f Kg\n", bMax)
    fmt.Printf("Rerata berat balita : %.2f Kg\n", rataRata)

}