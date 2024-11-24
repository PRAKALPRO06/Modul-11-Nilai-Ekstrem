package main
import (
	"fmt"
)

//Tegar Aji Pangestu 2311102021
func main() {
	var x, y int
	fmt.Print("masukkan jumlah ikan dan kapasitas wadah: ")
	fmt.Scan(&x, &y)
	berat := make([]float64, x)
	fmt.Println("masukkan berat tiap ikan: ")
	for i := 0; i < x; i++ {
		fmt.Scan(&berat[i])
	}
	jumlahWadah := (x + y - 1) / y
	totalBeratWadah := make([]float64, jumlahWadah)
	for i := 0; i < x; i++ {
		indekswadah := i / y
		totalBeratWadah[indekswadah] += berat[i]
	}
	fmt.Println("total berat tiap wadah: ")
	for _, total := range totalBeratWadah {
		fmt.Printf("%.2f", total)
	}
	fmt.Println()
	fmt.Println("rata rata berat tiap wadah: ")
	for _, total := range totalBeratWadah {
		ratarata := total / float64(y)
		fmt.Printf("%.2f", ratarata)
	}
	fmt.Println()
}