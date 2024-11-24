//2311102018-Aryo Tegar Sukarno
package main
import (
	"fmt"
)
func main() {
	
	const maxCapacity = 1000
	var weights [maxCapacity]float64
	var count int

	for {
		fmt.Println("\nProgram Pencatatan Berat Badan Anak Kelinci")
		fmt.Println("1. Tambah berat badan")
		fmt.Println("2. Lihat semua berat badan")
		fmt.Println("3. Hapus data terakhir")
		fmt.Println("4. Keluar")
		fmt.Print("Pilih opsi: ")
		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			if count >= maxCapacity {
				fmt.Println("Kapasitas maksimum telah tercapai. Tidak bisa menambah data lagi.")
				continue
			}

			fmt.Print("Masukkan berat badan anak kelinci (kg): ")
			var weight float64
			fmt.Scan(&weight)
			if weight <= 0 {
				fmt.Println("Berat badan harus lebih dari 0.")
				continue
			}

			weights[count] = weight
			count++
			fmt.Println("Data berhasil ditambahkan.")

		case 2:
			if count == 0 {
				fmt.Println("Tidak ada data yang tercatat.")
			} else {
				fmt.Println("Berat badan anak kelinci yang tercatat:")
				for i := 0; i < count; i++ {
					fmt.Printf("%d. %.2f kg\n", i+1, weights[i])
				}
			}

		case 3:
			if count == 0 {
				fmt.Println("Tidak ada data untuk dihapus.")
			} else {
				count--
				fmt.Printf("Data terakhir dengan berat %.2f kg telah dihapus.\n", weights[count])
			}

		case 4:
			fmt.Println("Keluar dari program. Terima kasih!")
			return

		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}
	}
}
