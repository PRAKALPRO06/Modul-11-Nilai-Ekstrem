package main

import (
	"fmt"
)

type dataIkan struct {
	berat          []float64
	jumlahIkan     int
	kapasitasWadah int
}

type wadahAkhir struct {
	beratWadah     []float64
	beratRataWadah []float64
}

func getInput() (dataIkan, error) {
	var data dataIkan
	data.berat = make([]float64, 1000)

	fmt.Print("Masukkan jumlah ikan (x) dan kapasitas wadah (y): ")
	fmt.Scan(&data.jumlahIkan, &data.kapasitasWadah)

	if err := validateInput(data.jumlahIkan, data.kapasitasWadah); err != nil {
		return data, err
	}

	fmt.Printf("Masukkan berat %d ikan:\n", data.jumlahIkan)
	for i := 0; i < data.jumlahIkan; i++ {
		fmt.Scan(&data.berat[i])
	}

	return data, nil
}

func validateInput(jumlahIkan, kapasitasWadah int) error {
	if jumlahIkan <= 0 || jumlahIkan > 1000 {
		return fmt.Errorf("Jumlah ikan (x) harus antara 1 dan 1000")
	}
	if kapasitasWadah <= 0 {
		return fmt.Errorf("Kapasitas wadah (y) harus lebih besar dari 0")
	}
	return nil
}

func calculateDistribution(data dataIkan) wadahAkhir {
	jumlahWadah := (data.jumlahIkan + data.kapasitasWadah - 1) / data.kapasitasWadah
	var results wadahAkhir
	results.beratWadah = make([]float64, jumlahWadah)
	results.beratRataWadah = make([]float64, jumlahWadah)

	ikanPerWadah := make([]int, jumlahWadah)

	for i := 0; i < data.jumlahIkan; i++ {
		indexWadah := i / data.kapasitasWadah
		results.beratWadah[indexWadah] += data.berat[i]
		ikanPerWadah[indexWadah]++
	}

	for i := 0; i < jumlahWadah; i++ {
		if ikanPerWadah[i] > 0 {
			results.beratRataWadah[i] = results.beratWadah[i] / float64(ikanPerWadah[i])
		}
	}

	return results
}

func displayResults(results wadahAkhir) {
	fmt.Println("\nTotal berat ikan di setiap wadah:")
	for i, berat := range results.beratWadah {
		fmt.Printf("Wadah %d: %.2f\n", i+1, berat)
	}

	fmt.Println("\nRata-rata berat ikan di setiap wadah:")
	for i, rata := range results.beratRataWadah {
		fmt.Printf("Wadah %d: %.2f\n", i+1, rata)
	}
}

func main() {
	data, err := getInput()
	if err != nil {
		fmt.Println(err)
		return
	}

	results := calculateDistribution(data)
	displayResults(results)
}
