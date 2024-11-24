package main

//Erwin Rivaldo Silaban
//23111022248
import (
	"fmt"
	"sort"
)

func main() {
	var beratKelinci = make([]float64, 0, 1000)

	var N int
	fmt.Print("Masukkan Jumlah Kelinci: ")
	fmt.Scan(&N)

	if N > 1000 {
		fmt.Println("Jumlah kelinci melebihi kapasitas.")
		return
	}

	fmt.Println("Masukkan Berat Kelinci Yang Akan Ditimbang:")
	for i := 0; i < N; i++ {
		var berat float64
		fmt.Printf("Berat Kelinci ke-%d (kg): ", i+1)
		fmt.Scan(&berat)
		beratKelinci = append(beratKelinci, berat)
	}
	sort.Float64s(beratKelinci)

	fmt.Println("\nBerat Kelinci dari yang terkecil sampai terbesar:")
	for i := 0; i < N; i++ {
		fmt.Printf("%.1f kg\n", beratKelinci[i])
	}
}
