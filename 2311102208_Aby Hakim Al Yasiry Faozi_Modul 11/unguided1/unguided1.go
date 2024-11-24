package main

import (
	"fmt"
)

func main() {
	var jumlahKelinci int
	fmt.Print("Masukkan jumlah anak kelinci: ")
	fmt.Scan(&jumlahKelinci)

	if jumlahKelinci <= 0 || jumlahKelinci > 1000 {
		fmt.Println("Jumlah anak kelinci harus angka")
		return
	}

	beratKelinci := make([]float64, jumlahKelinci)

	fmt.Printf("Masukkan berat %d anak kelinci:\n", jumlahKelinci)
	for i := 0; i < jumlahKelinci; i++ {
		fmt.Printf("Berat kelinci ke-%d: ", i+1)
		fmt.Scan(&beratKelinci[i])
	}

	terkecil := beratKelinci[0]
	terbesar := beratKelinci[0]

	for i := 1; i < jumlahKelinci; i++ {
		if beratKelinci[i] < terkecil {
			terkecil = beratKelinci[i]
		}
		if beratKelinci[i] > terbesar {
			terbesar = beratKelinci[i]
		}
	}

	fmt.Printf("Berat kelinci paling ringan: %.2f kg\n", terkecil)
	fmt.Printf("Berat kelinci paling berat: %.2f kg\n", terbesar)
}
