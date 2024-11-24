package main

import "fmt"

func main() {
	var x, y int

	fmt.Print("Masukkan jumlah ikan yang akan dijual dan banyaknya ikan dalam wadah (x y): ")
	fmt.Scan(&x, &y)

	if x < 1 || x > 1000 || y < 1 {
		fmt.Println("Jumlah ikan harus antara 1 dan 1000, dan banyaknya ikan dalam wadah harus lebih dari 0.")
		return
	}

	beratIkan := make([]float64, x)

	fmt.Println("Masukkan berat ikan:")
	for i := 0; i < x; i++ {
		fmt.Printf("Berat ikan ke-%d: ", i+1)
		fmt.Scan(&beratIkan[i])
	}

	var totalBerat []float64
	var totalIkanDiWadah []int

	for i := 0; i < x; i++ {
		if i%y == 0 {
			totalBerat = append(totalBerat, 0)
			totalIkanDiWadah = append(totalIkanDiWadah, 0)
		}
		totalBerat[len(totalBerat)-1] += beratIkan[i]
		totalIkanDiWadah[len(totalIkanDiWadah)-1]++
	}

	fmt.Println("Total berat ikan di setiap wadah:")
	for _, berat := range totalBerat {
		fmt.Printf("%.2f ", berat)
	}
	fmt.Println()

	fmt.Println("Berat rata-rata ikan di setiap wadah:")
	for i := 0; i < len(totalBerat); i++ {
		if totalIkanDiWadah[i] > 0 {
			rataRata := totalBerat[i] / float64(totalIkanDiWadah[i])
			fmt.Printf("%.2f ", rataRata)
		} else {
			fmt.Printf("0 ")
		}
	}
	fmt.Println()
}