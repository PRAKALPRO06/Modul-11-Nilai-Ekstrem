package main

import "fmt"

func carinilaiminmaks(data []float64, n int) (float64, float64) {

	min := data[0]
	maks := data[0]

	for i := 1; i < n; i++ {
		if data[i] < min {
			min = data[i]
		}
		if data[i] > maks {
			maks = data[i]
		}
	}
	return min, maks
}

func main() {
	var n int
	var berat [1000]float64

	fmt.Print("Jumlah anak kelinci berapa?: ")
	fmt.Scan(&n)

	if n < 1 || n > 1000 {
		fmt.Println("jumlah harusnya 1 sampai 1000.")
		return
	}
	fmt.Println("Berat kelinci masukkan : ")
	for i := 0; i < n; i++ {
		fmt.Printf("berat ke-%d: ", i+1)
		fmt.Scan(&berat[i])
	}

	min, maks := carinilaiminmaks(berat[:n], n)

	fmt.Printf("\n berat terkecil %.2f\n", min)
	fmt.Printf("berat terbesar %.2f\n", maks)
}
