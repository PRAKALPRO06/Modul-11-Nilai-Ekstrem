package main

import "fmt"

type arrayBerat [1000]float64

func terkecil(tabel arrayBerat, n int) float64 {
	var kecil float64 = tabel[0]
	var i int = 1
	for i < n {
		if tabel[i] < kecil {
			kecil = tabel[i]
		}
		i = i + 1
	}
	return kecil
}

func terbesar(tabel arrayBerat, n int) float64 {
	var besar float64 = tabel[0]
	var i int = 1
	for i < n {
		if tabel[i] > besar {
			besar = tabel[i]
		}
		i = i + 1
	}
	return besar
}

func main() {
	var jumlah int
	var beratKelinci arrayBerat

	fmt.Print("Masukkan jumlah kelinci (maksimal 1000): ")
	fmt.Scan(&jumlah)

	// Validasi input
	if jumlah < 1 || jumlah > 1000 {
		fmt.Println("Jumlah kelinci harus antara 1 sampai 1000")
		return
	}

	fmt.Println("Masukkan berat kelinci (dalam kg):")
	for i := 0; i < jumlah; i++ {
		fmt.Printf("Berat kelinci ke-%d: ", i+1)
		fmt.Scan(&beratKelinci[i])
	}

	beratTerkecil := terkecil(beratKelinci, jumlah)
	beratTerbesar := terbesar(beratKelinci, jumlah)

	fmt.Printf("\nBerat kelinci terkecil: %.2f kg\n", beratTerkecil)
	fmt.Printf("Berat kelinci terbesar: %.2f kg\n", beratTerbesar)
}
