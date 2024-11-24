package main

import "fmt"

func main() {
	var N int

	fmt.Print("Masukkan jumlah anak kelinci: ")
	fmt.Scan(&N)


	if N < 1 || N > 1000 {
		fmt.Println("Jumlah kelinci tidak valid.")
		return
	}

	beratKelinci := make([]float64, N)

	for i := 0; i < N; i++ {
		fmt.Printf("Masukkan berat anak kelinci ke-%d: ", i+1)
		fmt.Scan(&beratKelinci[i])
	}


	beratTerkecil := beratKelinci[0]
	beratTerbesar := beratKelinci[0]

	for _, berat := range beratKelinci {
		if berat < beratTerkecil {
			beratTerkecil = berat
		}
		if berat > beratTerbesar {
			beratTerbesar = berat
		}
	}


	fmt.Printf("Berat kelinci terkecil: %.2f\n", beratTerkecil)
	fmt.Printf("Berat kelinci terbesar: %.2f\n", beratTerbesar)
}