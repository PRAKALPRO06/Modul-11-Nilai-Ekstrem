package main

import (
	"fmt"
)

type BeratBalita [100]float64

func cariEkstrim(data BeratBalita, jumlah int, terkecil, terbesar *float64) {
	*terkecil = data[0]
	*terbesar = data[0]

	for index := 1; index < jumlah; index++ {
		if data[index] < *terkecil {
			*terkecil = data[index]
		}
		if data[index] > *terbesar {
			*terbesar = data[index]
		}
	}
}

func hitungRataRata(data BeratBalita, jumlah int) float64 {
	totalBerat := 0.0
	for i := 0; i < jumlah; i++ {
		totalBerat += data[i]
	}
	return totalBerat / float64(jumlah)
}

func main() {
	var jumlahData int
	var dataBerat BeratBalita

	fmt.Print("Masukkan jumlah balita: ")
	fmt.Scan(&jumlahData)

	if jumlahData < 1 || jumlahData > 100 {
		fmt.Println("Jumlah balita harus angka")
		return
	}

	for i := 0; i < jumlahData; i++ {
		fmt.Printf("Masukkan berat balita ke-%d: ", i+1)
		fmt.Scan(&dataBerat[i])
	}

	var beratTerkecil, beratTerbesar float64

	cariEkstrim(dataBerat, jumlahData, &beratTerkecil, &beratTerbesar)

	rataRata := hitungRataRata(dataBerat, jumlahData)

	fmt.Printf("Berat balita terkecil: %.2f kg\n", beratTerkecil)
	fmt.Printf("Berat balita terbesar: %.2f kg\n", beratTerbesar)
	fmt.Printf("Rata-rata berat balita: %.2f kg\n", rataRata)
}
