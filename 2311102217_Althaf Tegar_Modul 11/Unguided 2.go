package main

import "fmt"

func main() {
	var x22, y17 int
	fmt.Print("Masukkan jumlah ikan dan kapasitas wadah: ")
	fmt.Scan(&x22, &y17)

	berat_ikan := make([]float64, x22)
	fmt.Println("Masukkan berat tiap tiap ikan: ")
	for i := 0; i < x22; i++ {
		fmt.Scan(&berat_ikan[i])
	}

	jumlahWadah := x22 + y17 - 1/y17
	totalBeratWadah := make([]float64, jumlahWadah)

	for i := 0; i < x22; i++ {
		Wadah := i / y17
		totalBeratWadah[Wadah] += berat_ikan[i]
	}

	fmt.Print("Berat total wadah: ")
	for _, total := range totalBeratWadah {
		fmt.Printf("%2f", total)
	}
	fmt.Println()

	fmt.Print("Rata-rata berat tiap wadah: ")
	for _, total := range totalBeratWadah {
		rataRata := total / float64(y17)
		fmt.Printf("%2f", rataRata)
	}
	fmt.Println()

}
