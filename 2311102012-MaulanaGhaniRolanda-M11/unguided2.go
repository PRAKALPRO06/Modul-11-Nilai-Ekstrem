//2311102012
package main

import (
	"fmt"
)

func main() {
	// Deklarasi variabel
	var x, y int
	fmt.Print("Masukkan jumlah ikan (x) dan kapasitas per wadah (y): ")
	fmt.Scan(&x, &y)

	// Validasi input
	if x <= 0 || y <= 0 || x > 1000 {
		fmt.Println("Jumlah ikan atau kapasitas wadah tidak valid.")
		return
	}

	// Deklarasi array untuk berat ikan
	beratIkan := make([]float64, x)
	fmt.Println("Masukkan berat ikan satu per satu:")

	for i := 0; i < x; i++ {
		fmt.Printf("Berat ikan ke-%d: ", i+1)
		fmt.Scan(&beratIkan[i])
	}

	// Hitung jumlah wadah yang dibutuhkan
	jumlahWadah := (x + y - 1) / y // Pembulatan ke atas
	totalBeratWadah := make([]float64, jumlahWadah)

	// Distribusi ikan ke wadah
	for i := 0; i < x; i++ {
		wadahIndex := i / y
		totalBeratWadah[wadahIndex] += beratIkan[i]
	}

	// Hitung rata-rata berat per wadah
	var totalBeratSemua float64
	for _, total := range totalBeratWadah {
		totalBeratSemua += total
	}
	rataRata := totalBeratSemua / float64(jumlahWadah)

	// Output hasil
	fmt.Println("Total berat di setiap wadah:")
	for i, berat := range totalBeratWadah {
		fmt.Printf("Wadah %d: %.2f\n", i+1, berat)
	}
	fmt.Printf("Berat rata-rata per wadah: %.2f\n", rataRata)
}
