package main

import (
	"fmt"
	"math"
)

func main() {
	// Masukkan jumlah ikan (x) dan jumlah ikan per wadah (y)
	var x, y int
	fmt.Print("Masukkan jumlah ikan (x) dan jumlah ikan per wadah (y): ")
	fmt.Scan(&x, &y)

	// Validasi input
	if x <= 0 || y <= 0 || x > 1000 {
		fmt.Println("Input tidak valid. x harus > 0, y harus > 0, dan x <= 1000.")
		return
	}

	// Masukkan berat ikan
	beratIkan := make([]float64, x)
	fmt.Println("Masukkan berat ikan (pisahkan dengan spasi):")
	for i := 0; i < x; i++ {
		fmt.Scan(&beratIkan[i])
		if beratIkan[i] <= 0 {
			fmt.Println("Berat ikan harus bernilai positif.")
			return
		}
	}

	// Menghitung jumlah wadah
	jumlahWadah := int(math.Ceil(float64(x) / float64(y)))
	totalBeratPerWadah := make([]float64, jumlahWadah)

	// Distribusi ikan ke wadah
	for i := 0; i < x; i++ {
		indexWadah := i / y
		totalBeratPerWadah[indexWadah] += beratIkan[i]
	}

	// Menampilkan total berat di setiap wadah
	fmt.Println("Total berat di setiap wadah:")
	for i, total := range totalBeratPerWadah {
		fmt.Printf("Wadah %d: %.2f\n", i+1, total)
	}

	// Menghitung dan menampilkan rata-rata berat per wadah
	totalBeratKeseluruhan := 0.0
	for _, total := range totalBeratPerWadah {
		totalBeratKeseluruhan += total
	}
	rataRataBerat := totalBeratKeseluruhan / float64(jumlahWadah)
	fmt.Printf("Rata-rata berat per wadah: %.2f\n", rataRataBerat)
}
