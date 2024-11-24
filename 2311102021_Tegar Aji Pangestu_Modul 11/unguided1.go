package main
import (
	"fmt"
)

//Tegar Aji Pangestu 2311102021
func main() {
	var N int
	var berat [1000]float64
	fmt.Println("masuka jumlah anak kelinci:")
	fmt.Scan(&N)
	fmt.Println("masukan berat anak kelinci:")
	for i := 0; i < N; i++ {
		fmt.Scan(&berat[i])
	}
	min := berat[0]
	max := berat[0]
	for i := 1; i < N; i++ {
		if berat[i] < min {
			min = berat[i]
		}
		if berat[i] > min {
			max = berat[i]
		}
	}
	fmt.Println("berat terkecil: ", min)
	fmt.Println("berat terbesar: ", max)
}