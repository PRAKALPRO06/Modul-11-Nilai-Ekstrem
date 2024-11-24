package main

import (
	"fmt"
)

func main() {
	var jumlahWadah, jumlahIkanPerWadah int
	fmt.Print("Masukkan jumlah wadah (x): ")
	fmt.Scan(&jumlahWadah)
	fmt.Print("Masukkan jumlah ikan per wadah (y): ")
	fmt.Scan(&jumlahIkanPerWadah)

	// Validasi input
	if jumlahWadah < 1 || jumlahWadah > 1000 || jumlahIkanPerWadah < 1 {
		fmt.Println("Nilai jumlah wadah harus antara 1-1000 dan jumlah ikan per wadah harus lebih dari 0.")
		return
	}

	// Membuat array untuk menyimpan berat ikan
	beratIkan := make([]float64, jumlahWadah*jumlahIkanPerWadah)
	fmt.Println("Masukkan berat ikan satu per satu:")
	for i := 0; i < jumlahWadah*jumlahIkanPerWadah; i++ {
		fmt.Printf("Berat ikan ke-%d: ", i+1)
		fmt.Scan(&beratIkan[i])
	}

	// Membuat array untuk menyimpan total berat dan rata-rata berat
	totalBeratPerWadah := make([]float64, jumlahWadah)
	rataRataBeratPerWadah := make([]float64, jumlahWadah)

	// Menghitung total berat dan rata-rata berat ikan per wadah
	for i := 0; i < jumlahWadah; i++ {
		totalBerat := 0.0
		for j := 0; j < jumlahIkanPerWadah; j++ {
			totalBerat += beratIkan[i*jumlahIkanPerWadah+j]
		}
		totalBeratPerWadah[i] = totalBerat
		rataRataBeratPerWadah[i] = totalBerat / float64(jumlahIkanPerWadah)
	}

	// Menampilkan total berat ikan di setiap wadah
	fmt.Println("\nTotal berat ikan di setiap wadah:")
	for i := 0; i < jumlahWadah; i++ {
		fmt.Printf("Wadah %d: %.2f kg\n", i+1, totalBeratPerWadah[i])
	}

	// Menampilkan rata-rata berat ikan di setiap wadah
	fmt.Println("\nRata-rata berat ikan di setiap wadah:")
	for i := 0; i < jumlahWadah; i++ {
		fmt.Printf("Wadah %d: %.2f kg\n", i+1, rataRataBeratPerWadah[i])
	}
}
