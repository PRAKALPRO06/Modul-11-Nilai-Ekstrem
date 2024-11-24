	package main

	import "fmt"

	// Mendeklarasikan tipe data array arrInt dengan panjang 2023
	type arrInt [2023]int

	// Fungsi untuk mencari indeks elemen terkecil dalam array
	func terkecil(tabInt arrInt, n int) int {
		var idx int = 0  // idx menyimpan indeks elemen terkecil
		var j int = 1
		for j < n {
			if tabInt[idx] > tabInt[j] {
				idx = j  // Simpan indeks j jika elemen di indeks j lebih kecil
			}
			j = j + 1
		}
		return idx
	}

	// Fungsi main untuk menguji fungsi terkecil
	func main() {
		var n int
		var tab arrInt

		// Meminta input jumlah elemen array
		fmt.Print("Masukkan jumlah elemen (maks 2023): ")
		fmt.Scan(&n)

		// Validasi input jumlah elemen
		if n < 1 || n > 2023 {
			fmt.Println("Jumlah elemen harus antara 1 dan 2023.")
			return
		}

		// Memasukkan elemen-elemen array
		fmt.Println("Masukkan elemen-elemen array:")
		for i := 0; i < n; i++ {
			fmt.Print("Elemen ke-", i+1, ": ")
			fmt.Scan(&tab[i])
		}

		// Memanggil fungsi terkecil untuk menemukan indeks elemen terkecil
		idxMin := terkecil(tab, n)

		// Menampilkan nilai dan indeks terkecil
		fmt.Println("Nilai terkecil dalam array adalah:", tab[idxMin], "pada indeks:", idxMin)
	}
