package main

import (
	"fmt"
)

const maxIkan = 1000

func main() {
	var x, y int
	var beratIkan [maxIkan]float64
	var totalBeratWadah []float64

	// Meminta input jumlah ikan dan kapasitas wadah
	fmt.Print("Masukkan jumlah ikan: ")
	fmt.Scan(&x)
	fmt.Print("Masukkan kapasitas wadah: ")
	fmt.Scan(&y)

	// Validasi input
	if x <= 0 || y <= 0 || x > maxIkan {
		fmt.Println("Input tidak valid.")
		return
	}

	// Memasukkan data berat ikan
	fmt.Println("Masukkan berat masing-masing ikan (dalam kg):")
	for i := 0; i < x; i++ {
		fmt.Scan(&beratIkan[i])
	}

	// Menghitung total berat ikan di setiap wadah
	var totalBerat float64
	for i := 0; i < x; i++ {
		totalBerat += beratIkan[i]
		if (i+1)%y == 0 || i == x-1 {
			totalBeratWadah = append(totalBeratWadah, totalBerat)
			totalBerat = 0
		}
	}

	// Menampilkan total berat ikan di setiap wadah
	fmt.Println("Total berat ikan di setiap wadah:")
	for _, berat := range totalBeratWadah {
		fmt.Printf("%.2f kg ", berat)
	}
	fmt.Println()

	// Menghitung berat rata-rata ikan di setiap wadah
	var totalSemuaWadah float64
	for _, berat := range totalBeratWadah {
		totalSemuaWadah += berat
	}
	rataRata := totalSemuaWadah / float64(len(totalBeratWadah))
	fmt.Printf("Berat rata-rata ikan di setiap wadah: %.2f kg\n", rataRata)
}
