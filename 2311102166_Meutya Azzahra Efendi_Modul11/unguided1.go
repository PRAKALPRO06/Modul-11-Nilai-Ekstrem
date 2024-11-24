// Meutya Azzahra Efendi
// 2311102166
// IF-11-06

package main

import "fmt"

func main() {
	// Kapasitas array
	const kapasitas = 1000
	var beratKelinci [kapasitas]float64

	// Input jumlah kelinci
	var N int
	fmt.Print("Masukkan jumlah anak kelinci: ")
	fmt.Scan(&N)

	// Validasi jumlah kelinci
	if N < 1 || N > kapasitas {
		fmt.Println("Jumlah kelinci harus antara 1 dan 1000.")
		return
	}

	// Input berat kelinci
	for i := 0; i < N; i++ {
		fmt.Printf("Masukkan berat kelinci ke-%d: ", i+1)
		fmt.Scan(&beratKelinci[i])
	}

	// Mencari berat terkecil dan terbesar
	beratTerkecil := beratKelinci[0]
	beratTerbesar := beratKelinci[0]

	for i := 1; i < N; i++ {
		if beratKelinci[i] < beratTerkecil {
			beratTerkecil = beratKelinci[i]
		}
		if beratKelinci[i] > beratTerbesar {
			beratTerbesar = beratKelinci[i]
		}
	}

	fmt.Printf("Berat kelinci terkecil: %.2f\n", beratTerkecil)
	fmt.Printf("Berat kelinci terbesar: %.2f\n", beratTerbesar)
}