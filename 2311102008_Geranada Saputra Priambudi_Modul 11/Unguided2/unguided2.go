package main

import "fmt"

type Ikan struct {
	Berat float64
}

type Wadah struct {
	TotalBerat float64
	RataRata   float64
}

type arrIkan [1000]Ikan

func main() {
	var x, y int
	fmt.Print("Masukkan jumlah ikan yang akan dijual (x) dan jumlah ikan per wadah (y): ")
	fmt.Scan(&x, &y)

	if x > 1000 {
		fmt.Println("Jumlah ikan melebihi kapasitas array.")
		return
	}
	var ikanArray arrIkan

	fmt.Println("Masukkan berat ikan:")
	for i := 0; i < x; i++ {
		fmt.Scan(&ikanArray[i].Berat)
	}

	jumlahWadah := x / y
	if x%y != 0 {
		jumlahWadah++
	}

	wadahArray := make([]Wadah, jumlahWadah)

	for i := 0; i < x; i++ {
		indexWadah := i / y
		wadahArray[indexWadah].TotalBerat += ikanArray[i].Berat
	}

	fmt.Println("Total berat ikan di setiap wadah:")
	for _, wadah := range wadahArray {
		fmt.Printf("%.2f ", wadah.TotalBerat)
	}
	fmt.Println()

	fmt.Println("Rata-rata berat ikan di setiap wadah:")
	for i := 0; i < jumlahWadah; i++ {
		var jumlahIkanPerWadah int
		if i == jumlahWadah-1 {
			jumlahIkanPerWadah = x % y
			if jumlahIkanPerWadah == 0 && x > 0 {
				jumlahIkanPerWadah = y
			}
		} else {
			jumlahIkanPerWadah = y
		}

		wadahArray[i].RataRata = wadahArray[i].TotalBerat / float64(jumlahIkanPerWadah)
		fmt.Printf("%.2f ", wadahArray[i].RataRata)
	}
	fmt.Println()
}
