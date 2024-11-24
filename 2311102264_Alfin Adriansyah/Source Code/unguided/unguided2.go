package main

import "fmt"

type arrayBerat [1000]float64

func hitungRataRata(wadah []float64) float64 {
	var total float64
	for _, berat := range wadah {
		total += berat
	}
	return total / float64(len(wadah))
}

func main() {
	var jumlahIkan, ikanPerWadah int
	var beratIkan arrayBerat

	fmt.Print("Masukkan jumlah ikan dan ikan per wadah: ")
	fmt.Scan(&jumlahIkan, &ikanPerWadah)

	fmt.Println("Masukkan berat ikan (dalam kg):")
	for i := 0; i < jumlahIkan; i++ {
		fmt.Printf("Berat ikan ke-%d: ", i+1)
		fmt.Scan(&beratIkan[i])
	}

	jumlahWadah := (jumlahIkan + ikanPerWadah - 1) / ikanPerWadah
	wadah := make([]float64, jumlahWadah)

	for i := 0; i < jumlahIkan; i++ {
		wadahKe := i / ikanPerWadah
		wadah[wadahKe] += beratIkan[i]
	}

	fmt.Print("\nBerat ikan di setiap wadah: ")
	for _, berat := range wadah {
		fmt.Printf("%.2f ", berat)
	}

	rataRata := hitungRataRata(wadah)
	fmt.Printf("\nRata-rata berat ikan per wadah: %.2f kg\n", rataRata)
}
