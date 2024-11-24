package main

import "fmt"

func masukanBeratKelinci(n int) []float64 {
	beratKelinci := make([]float64, n)
	for i := 0; i < n; i++ {
		fmt.Printf("Berat Kelinci %d: ", i+1)
		fmt.Scan(&beratKelinci[i])
	}
	return beratKelinci
}

func cariBeratTerkecilTerbesar(beratKelinci []float64) (float64, float64) {
	beratMin := beratKelinci[0]
	beratMax := beratKelinci[0]

	for i := 1; i < len(beratKelinci); i++ {
		if beratKelinci[i] < beratMin {
			beratMin = beratKelinci[i]
		}
		if beratKelinci[i] > beratMax {
			beratMax = beratKelinci[i]
		}
	}

	return beratMin, beratMax
}

func main() {
	var n_226 int

	fmt.Print("Masukkan Jumlah Kelinci yang Ingin Dihitung Beratnya: ")
	fmt.Scan(&n_226)

	beratKelinci := masukanBeratKelinci(n_226)
	beratMin, beratMax := cariBeratTerkecilTerbesar(beratKelinci)
	fmt.Printf("\nBerat Kelinci Paling Ringan: %.2f\nBerat Kelinci Paling Berat: %.2f\n", beratMin, beratMax)
}