package main

import "fmt"

func main() {
	var N int

	// Input jumlah anak kelinci
	fmt.Print("Masukkan jumlah anak kelinci (N): ")
	fmt.Scan(&N)

	// Validasi jumlah N
	if N > 1000 {
		fmt.Println("Jumlah anak kelinci tidak boleh lebih dari 1000.")
		return
	}

	// Deklarasi array untuk menyimpan berat
	berat := make([]float64, N)

	// Input berat anak kelinci
	fmt.Printf("Masukkan %d berat anak kelinci (pisahkan dengan spasi):\n", N)
	for i := 0; i < N; i++ {
		fmt.Scan(&berat[i])
	}

	// Inisialisasi berat terkecil dan terbesar
	beratTerkecil := berat[0]
	beratTerbesar := berat[0]

	// Pencarian nilai terkecil dan terbesar
	for i := 1; i < N; i++ {
		if berat[i] < beratTerkecil {
			beratTerkecil = berat[i]
		}
		if berat[i] > beratTerbesar {
			beratTerbesar = berat[i]
		}
	}

	// Cetak hasil
	fmt.Printf("Berat terkecil: %.2f\n", beratTerkecil)
	fmt.Printf("Berat terbesar: %.2f\n", beratTerbesar)
}
