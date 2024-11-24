package main

import "fmt"

func main() {
	var jumKelinci_2311102241 int
	var beratKelinci = make([]float64, 0, 1000)

	fmt.Print("Masukkan jumlah kelinci: ")
	fmt.Scan(&jumKelinci_2311102241)

	if jumKelinci_2311102241 > 1000 {
		fmt.Println("Maaf, jumlah kelinci tidak boleh lebih dari 1000")
		return
	}

	for i := 1; i <= jumKelinci_2311102241; i++ {
		var berat float64
		fmt.Printf("Masukkan berat kelinci ke-%d: ", i)
		fmt.Scan(&berat)
		beratKelinci = append(beratKelinci, berat)
	}

	var beratTerkecil = beratKelinci[0]
	var beratTerbesar = beratKelinci[0]

	for i := 0; i < jumKelinci_2311102241; i++ {
		if beratKelinci[i] < beratTerkecil {
			beratTerkecil = beratKelinci[i]
		}
		if beratKelinci[i] > beratTerbesar {
			beratTerbesar = beratKelinci[i]
		}
	}

	fmt.Println("\nHasil Pencarian:")
	fmt.Printf("Berat kelinci terkecil: %.2fkg\n", beratTerkecil)
	fmt.Printf("Berat kelinci terbesar: %.2fkg\n", beratTerbesar)
}
