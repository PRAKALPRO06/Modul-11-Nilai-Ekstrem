package main

import "fmt"

type arrInt [1000]float64

func nilaiEkstrim(berat arrInt, jumlahKelinci int) (float64, float64) {
	nilaiMin := berat[0]
	nilaiMax := berat[0]

	for i := 1; i < jumlahKelinci; i++ {
		if berat[i] < nilaiMin {
			nilaiMin = berat[i]
		}
		if berat[i] > nilaiMax {
			nilaiMax = berat[i]
		}
	}
	return nilaiMin, nilaiMax
}

func main() {
	var jumlahKelinci int
	var beratKelinci arrInt

	fmt.Print("Masukkan jumlah anak kelinci: ")
	fmt.Scan(&jumlahKelinci)

	if jumlahKelinci < 1 || jumlahKelinci > 1000 {
		fmt.Println("Masukan harus angka 1-1000!")
		return
	}

	for i := 0; i < jumlahKelinci; i++ {
		fmt.Printf("Masukkan berat kelinci ke-%d: ", i+1)
		fmt.Scan(&beratKelinci[i])
	}

	nilaiMin, nilaiMax := nilaiEkstrim(beratKelinci, jumlahKelinci)

	fmt.Printf("\nBerat anak kelinci terkecil: %.2f\n", nilaiMin)
	fmt.Printf("Berat anak kelinci terbesar: %.2f\n", nilaiMax)
}
