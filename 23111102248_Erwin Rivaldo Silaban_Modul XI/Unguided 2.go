package main

//2311102248
//Erwin Rivaldo Silaban
import (
	"fmt"
)

func main() {
	var x, y int

	fmt.Print("Masukkan jumlah ikan yang akan dijual dan banyaknya ikan per wadah (x y): ")
	fmt.Scan(&x, &y)

	if x <= 0 || y <= 0 {
		fmt.Println("Jumlah ikan dan jumlah ikan per wadah harus lebih besar dari 0.")
		return
	}

	ikan := make([]float64, x)

	fmt.Printf("Masukkan berat %d ikan yang akan dijual:\n", x)
	for i := 0; i < x; i++ {
		fmt.Printf("Berat ikan ke-%d: ", i+1)
		fmt.Scan(&ikan[i])
	}

	var totalBerat []float64
	var total float64
	for i := 0; i < x; i++ {
		total += ikan[i]
		if (i+1)%y == 0 || i == x-1 {
			totalBerat = append(totalBerat, total)
			total = 0
		}
	}

	var rataRataBerat []float64
	for _, berat := range totalBerat {
		rataRata := berat / float64(y)
		rataRataBerat = append(rataRataBerat, rataRata)
	}

	fmt.Println("\nTotal berat ikan di setiap wadah:")
	for _, berat := range totalBerat {
		fmt.Printf("%.2fkg\n ", berat)
	}
	fmt.Println()

	fmt.Println("Rata-rata berat ikan di setiap wadah:")
	for _, rata := range rataRataBerat {
		fmt.Printf("%.2fkg\n ", rata)
	}
	fmt.Println()
}
