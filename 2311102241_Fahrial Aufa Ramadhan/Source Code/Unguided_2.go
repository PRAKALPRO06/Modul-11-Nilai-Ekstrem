package main

import "fmt"

func main() {
	var jumlahIkan int
	var jumlahWadah int

	fmt.Print("Berapa jumlah ikan(x): ")
	fmt.Scan(&jumlahIkan)

	if jumlahIkan > 1000 {
		fmt.Println("Maaf, jumlah ikan tidak boleh lebih dari 1000")
		return
	}

	fmt.Print("Berapa jumlah wadah(y): ")
	fmt.Scan(&jumlahWadah)

	var beratIkan = make([]float64, 0, 1000)

	for i := 0; i < jumlahIkan; i++ {
		var berat float64
		fmt.Printf("Masukkan berat ikan ke-%d: ", i+1)
		fmt.Scan(&berat)
		beratIkan = append(beratIkan, berat)
	}

	var totalPerWadah = make([]float64, jumlahWadah)

	var ikanPerWadah = jumlahIkan / jumlahWadah
	var nomorIkan = 0

	for wadah := 0; wadah < jumlahWadah; wadah++ {
		var totalBerat float64 = 0

		for ikan := 0; ikan < ikanPerWadah; ikan++ {
			if nomorIkan < jumlahIkan {
				totalBerat = totalBerat + beratIkan[nomorIkan]
				nomorIkan = nomorIkan + 1
			}
		}

		totalPerWadah[wadah] = totalBerat
	}

	fmt.Println("\nTotal berat tiap wadah:")
	for i := 0; i < jumlahWadah; i++ {
		fmt.Printf("Wadah %d: %.1f kg\n", i+1, totalPerWadah[i])
	}

	var totalSemuaWadah float64 = 0
	for i := 0; i < jumlahWadah; i++ {
		totalSemuaWadah = totalSemuaWadah + totalPerWadah[i]
	}
	var rataRata = totalSemuaWadah / float64(jumlahWadah)

	fmt.Printf("\nRata-rata berat ikan per wadah: %.1f kg\n", rataRata)
}
