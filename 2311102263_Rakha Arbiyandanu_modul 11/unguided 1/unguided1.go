package main
import "fmt"

func main() {
	var N int
	var berat [1000]float64
	var terkecil ,terbesar float64

	fmt.Print("masukkan banyaknya anak kelinci: ")
	fmt.Scanln(&N)

	fmt.Println("Masukkan berat anak kecil satu per satu: ")
	for i := 0; i < N; i++ {
		fmt.Printf("Kelinci ke-%d: ", i+1)
		fmt.Scanln(&berat[i])
	}

	terkecil = berat[0]
	terbesar = berat[0]

	for i := 1; i < N; i++ {
		if berat[i] < terkecil {
			terkecil = berat[i]
		}
		if berat[i] > terbesar{
			terbesar = berat[i]
		}
	}
	fmt.Printf("Berat terkecil: %.2f\n", terkecil)
    fmt.Printf("Berat terbesar: %.2f\n", terbesar)
}