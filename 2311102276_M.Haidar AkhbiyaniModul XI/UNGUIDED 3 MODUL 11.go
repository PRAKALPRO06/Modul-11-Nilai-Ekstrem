package main
import "fmt"

type arrBalita [100]float64

func hitungMinMax(arrBerat arrBalita, n int, bMin, bMax *float64) {
	*bMin = arrBerat[0]
	*bMax = arrBerat[0]
	for i := 1; i < n; i++ {
		if arrBerat[i] < *bMin {
			*bMin = arrBerat[i]
		}
		if arrBerat[i] > *bMax {
			*bMax = arrBerat[i]
		}
	}
}

func rerata(arrBerat arrBalita, n int) float64 {
	total := 0.0
	for i := 0; i < n; i++ {
		total += arrBerat[i]
	}
	return total / float64(n)
}

func main() {
	var arrBerat arrBalita
	var n int
	var bMin, bMax float64

	fmt.Print("Masukkan jumlah balita (maks 100): ")
	fmt.Scan(&n)

	if n > 100 {
		fmt.Println("Jumlah balita tidak boleh lebih dari 100.")
		return
	}

	fmt.Println("Masukkan berat balita (kg):")
	for i := 0; i < n; i++ {
		fmt.Printf("Berat balita ke-%d: ", i+1)
		fmt.Scan(&arrBerat[i])
	}

	hitungMinMax(arrBerat, n, &bMin, &bMax)

	rerataBerat := rerata(arrBerat, n)

	fmt.Printf("Berat minimum: %.2f kg\n", bMin)
	fmt.Printf("Berat maksimum: %.2f kg\n", bMax)
	fmt.Printf("Rata-rata berat: %.2f kg\n", rerataBerat)
}
