package main

import (
	"fmt"
)

func main() {
	var jumlahIkan, jumlahWadah int
	fmt.Print("Masukkan jumlah ikan dan wadah: ")
	fmt.Scan(&jumlahIkan, &jumlahWadah)

	if jumlahIkan <= 0 || jumlahIkan > 1000 || jumlahWadah <= 0 || jumlahWadah > 1000 {
		fmt.Println("Input tidak valid")
		return
	}

	beratIkan := make([]float64, jumlahIkan)
	fmt.Printf("Masukkan berat dari %d ikan:\n", jumlahIkan)
	for i := 0; i < jumlahIkan; i++ {
		fmt.Scan(&beratIkan[i])
	}

	beratTotalWadah := make([]float64, jumlahWadah)
	jumlahIkanPerWadah := make([]int, jumlahWadah)

	for i := 0; i < jumlahIkan; i++ {
		indeksWadah := i % jumlahWadah
		beratTotalWadah[indeksWadah] += beratIkan[i]
		jumlahIkanPerWadah[indeksWadah]++
	}

	rataRataWadah := make([]float64, jumlahWadah)
	for i := 0; i < jumlahWadah; i++ {
		if jumlahIkanPerWadah[i] > 0 {
			rataRataWadah[i] = beratTotalWadah[i] / float64(jumlahIkanPerWadah[i])
		}
	}

	fmt.Println("Berat total pada masing-masing wadah:")
	for i := 0; i < jumlahWadah; i++ {
		fmt.Printf("Wadah %d: %.2f kg\n", i+1, beratTotalWadah[i])
	}

	fmt.Println("Rata-rata berat ikan di masing-masing wadah:")
	for i := 0; i < jumlahWadah; i++ {
		fmt.Printf("Wadah %d: %.2f kg\n", i+1, rataRataWadah[i])
	}
}
