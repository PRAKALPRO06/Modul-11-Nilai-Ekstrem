// 2311102266_Hanif Reyhan Zhafran Arytona
package main

import "fmt"

type arrInt [2023]int

func terkecil(tabInt arrInt, n int) int {
	var idx int = 0
	for j := 1; j < n; j++ {
		if tabInt[idx] > tabInt[j] {
			idx = j
		}
	}
	return idx
}

func main() {
	var n int
	var tab arrInt

	fmt.Print("Masukkan jumlah elemen (maks 2023): ")
	fmt.Scan(&n)

	if n < 1 || n > 2023 {
		fmt.Println("Jumlah elemen harus antara 1 dan 2023.")
		return
	}

	fmt.Println("Masukkan elemen-elemen array:")
	for i := 0; i < n; i++ {
		fmt.Printf("Elemen ke-%d: ", i+1)
		fmt.Scan(&tab[i])
	}

	idxMin := terkecil(tab, n)

	fmt.Println("Nilai terkecil dalam array adalah:", tab[idxMin], "pada indeks:", idxMin)
}
