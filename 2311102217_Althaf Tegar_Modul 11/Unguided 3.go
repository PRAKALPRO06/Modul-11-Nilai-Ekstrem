package main

import "fmt"

func main() {
	var n217 int
	fmt.Print("Masukkan jumlah balita: ")
	fmt.Scan(&n217)

	if n217 <= 0 {
		return
	}

	beratBalita := make([]float64, n217)
	for i := 0; i < n217; i++ {
		fmt.Printf("Masukkan berat balita ke-%d (dalam kg): ", i+1)
		fmt.Scan(&beratBalita[i])
	}

	beratTerkecil := beratBalita[0]
	beratTerbesar := beratBalita[0]
	totalBerat := 0.0

	for _, berat := range beratBalita {
		if berat < beratTerkecil {
			beratTerkecil = berat
		}
		if berat > beratTerbesar {
			beratTerbesar = berat
		}
		totalBerat += berat
	}

	rataRata := totalBerat / float64(n217)

	fmt.Printf("Berat terkecil: %.2f kg\n", beratTerkecil)
	fmt.Printf("Berat terbesar: %.2f kg\n", beratTerbesar)
	fmt.Printf("Rata-rata berat: %.2f kg\n", rataRata)
}
