package main

import "fmt"

const maxIkan = 1000
var beratIkan [maxIkan]float64

func hitungTotalBerat(beratIkan []float64, x, y int) ([]float64, []int) {
	var totalBerat []float64
	var totalIkan []int

	for i := 0; i < x; i++ {
		if i%y == 0 {
			totalBerat = append(totalBerat, 0)
			totalIkan = append(totalIkan, 0)
		}
		totalBerat[len(totalBerat)-1] += beratIkan[i]
		totalIkan[len(totalIkan)-1]++
	}
	return totalBerat, totalIkan
}

func hitungRataRata(totalBerat []float64, totalIkan []int) []float64 {
	var rataRata []float64
	for i := 0; i < len(totalBerat); i++ {
		if totalIkan[i] > 0 {
			rataRata = append(rataRata, totalBerat[i]/float64(totalIkan[i]))
		} else {
			rataRata = append(rataRata, 0)
		}
	}
	return rataRata
}

func main() {
	var x_226, y_226 int
	
    fmt.Print("Masukkan jumlah ikan dan ikan per wadah (x y): ")
	fmt.Scan(&x_226, &y_226)
	fmt.Printf("Masukkan berat %d ikan: ", x_226)
	for i := 0; i < x_226; i++ {
		fmt.Scan(&beratIkan[i])
	}

	totalBerat, totalIkan := hitungTotalBerat(beratIkan[:x_226], x_226, y_226)

    fmt.Println("Total berat ikan di setiap wadah: ")
	for _, berat := range totalBerat {
		fmt.Printf("%.2f ", berat)
	}
	fmt.Println()

    fmt.Println("Rata-rata berat ikan di setiap wadah: ")
	rataRata := hitungRataRata(totalBerat, totalIkan)
	for _, rata := range rataRata {
		fmt.Printf("%.2f ", rata)
	}
	fmt.Println()
}