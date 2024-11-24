package main
import (
	"fmt"
)

func main() {
	var jumlahIkan, kapasitasWadah int
	var beratIkan [1000]float64

	fmt.Print("Masukkan jumlah ikan dan kapasitas wadah (x y): ")
	fmt.Scan(&jumlahIkan, &kapasitasWadah)

	if jumlahIkan < 1 || jumlahIkan > 1000 || kapasitasWadah < 1 {
		fmt.Println("Jumlah ikan harus antara 1 dan 1000, kapasitas wadah harus lebih dari 0.")
		return
	}

	fmt.Println("Masukkan berat ikan:")
	for i := 0; i < jumlahIkan; i++ {
		fmt.Printf("Berat ikan ke-%d: ", i+1)
		fmt.Scan(&beratIkan[i])
	}

	var jumlahWadah = (jumlahIkan + kapasitasWadah - 1) / kapasitasWadah
	var totalBeratPerWadah []float64
	var totalBerat float64

	for i := 0; i < jumlahWadah; i++ {
		var beratWadah float64 = 0
		for j := i * kapasitasWadah; j < (i+1)*kapasitasWadah && j < jumlahIkan; j++ {
			beratWadah += beratIkan[j]
		}
		totalBeratPerWadah = append(totalBeratPerWadah, beratWadah)
		totalBerat += beratWadah
	}

	fmt.Println("Total berat ikan di setiap wadah:")
	for i, berat := range totalBeratPerWadah {
		fmt.Printf("Wadah ke-%d: %.2f\n", i+1, berat)
	}

	var rataRataBerat float64 = totalBerat / float64(jumlahWadah)
	fmt.Printf("Berat rata-rata ikan di setiap wadah: %.2f\n", rataRataBerat)
}
