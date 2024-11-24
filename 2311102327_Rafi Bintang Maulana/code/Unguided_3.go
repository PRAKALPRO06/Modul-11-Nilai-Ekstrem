package main
import (
	"fmt"
)

type arrBalita [100]float64
func hitungMinMax(arrBerat arrBalita, jumlah int, bMin, bMax *float64) {
	*bMin = arrBerat[0]
	*bMax = arrBerat[0]

	for i := 1; i < jumlah; i++ {
		if arrBerat[i] < *bMin {
			*bMin = arrBerat[i]
		}
		if arrBerat[i] > *bMax {
			*bMax = arrBerat[i]
		}
	}
}

func rerata(arrBerat arrBalita, jumlah int) float64 {
	var total float64
	for i := 0; i < jumlah; i++ {
		total += arrBerat[i]
	}
	return total / float64(jumlah)
}

func main() {
	var jumlahBalita int
	var beratBalita arrBalita
	var beratTerkecil, beratTerbesar, rataRata float64

	fmt.Print("Masukkan banyak data berat balita: ")
	fmt.Scan(&jumlahBalita)

	if jumlahBalita < 1 || jumlahBalita > 100 {
		fmt.Println("Jumlah balita harus antara 1 dan 100.")
		return
	}

	fmt.Println("Masukkan berat balita:")
	for i := 0; i < jumlahBalita; i++ {
		fmt.Printf("Masukkan berat balita ke-%d: ", i+1)
		fmt.Scan(&beratBalita[i])
	}

	hitungMinMax(beratBalita, jumlahBalita, &beratTerkecil, &beratTerbesar)
	rataRata = rerata(beratBalita, jumlahBalita)

	fmt.Printf("\nBerat balita minimum: %.2f kg\n", beratTerkecil)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", beratTerbesar)
	fmt.Printf("Rerata berat balita: %.2f kg\n", rataRata)
}