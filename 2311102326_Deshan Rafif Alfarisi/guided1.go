package main

import "fmt"

// Mendeklarasikan tipe data array arrInt dengan panjang 2023
type arrInt [2023]int

// Fungsi untuk mencari elemen terkecil dalam array
func terkecil(tabInt arrInt, n int) int {
	var min int = tabInt[0]
	var j int = 1
	for j < n {
		if min > tabInt[j] {
			min = tabInt[j]
		}
		j = j + 1
	}
	return min
}

// Fungsi main
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

	// Memanggil fungsi terkecil untuk menemukan nilai terkecil
	minVal := terkecil(tab, n)

	// Menampilkan nilai terkecil
	fmt.Println("Nilai terkecil dalam array adalah:", minVal)
}
