// Meutya Azzahra Efendi
// 2311102166
// IF-11-06

package main

import "fmt"

func main() {
	// Kapasitas array
	const kapasitas = 1000
	var beratIkan [kapasitas]float64

	// Input jumlah ikan dan wadah
	var x, y int
	fmt.Print("Masukkan jumlah ikan (x) dan jumlah ikan per wadah (y): ")
	fmt.Scan(&x, &y)

	// Validasi jumlah ikan
	if x < 1 || x > kapasitas || y < 1 {
		fmt.Println("Jumlah ikan harus antara 1 dan 1000, dan jumlah ikan per wadah harus lebih dari 0.")
		return
	}

	// Input berat ikan
	fmt.Println("Masukkan berat ikan:")
	for i := 0; i < x; i++ {
		fmt.Printf("Berat ikan ke-%d: ", i+1)
		fmt.Scan(&beratIkan[i])
	}

	// Menghitung total berat ikan di setiap wadah
	var totalBeratPerWadah []float64
	var jumlahIkanDiWadah int
	var totalBerat float64

	for i := 0; i < x; i++ {
		totalBerat += beratIkan[i]
		jumlahIkanDiWadah++

		// Jika sudah mencapai jumlah ikan per wadah, simpan total berat dan reset
		if jumlahIkanDiWadah == y || i == x-1 {
			totalBeratPerWadah = append(totalBeratPerWadah, totalBerat)
			totalBerat = 0
			jumlahIkanDiWadah = 0
		}
	}

	// Output total berat ikan di setiap wadah
	fmt.Println("Total berat ikan di setiap wadah:")
	for _, total := range totalBeratPerWadah {
		fmt.Printf("%.2f ", total)
	}
	fmt.Println()

	// Menghitung dan menampilkan berat rata-rata ikan di setiap wadah
	var totalRataRata float64
	for _, total := range totalBeratPerWadah {
		totalRataRata += total
	}

	// Menghitung jumlah wadah yang digunakan
	jumlahWadah := len(totalBeratPerWadah)

	if jumlahWadah > 0 {
		rataRata := totalRataRata / float64(jumlahWadah)
		fmt.Printf("Berat rata-rata ikan di setiap wadah: %.2f\n", rataRata)
	} else {
		fmt.Println("Tidak ada wadah yang digunakan.")
	}
}