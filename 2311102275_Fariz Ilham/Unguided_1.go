package main

import "fmt"

func main() {
	var n int
	var minweight, maxweight float64
	var weights [1000]float64

	fmt.Print("Masukkan jumlah anak kelinci : ")
	fmt.Scan(&n)

	if n < 1 || n > 1000 {
		fmt.Println("Jumlah anak kelinci harus antara 1 dan 1000.")
		return
	}

	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan berat kelinci ke-%d: ", i+1)
		fmt.Scan(&weights[i])

		if i == 0 {
			minweight = weights[i]
			maxweight = weights[i]
		} else {
			if weights[i] < minweight {
				minweight = weights[i]
			}
			if weights[i] > maxweight {
				maxweight = weights[i]
			}
		}
	}
	fmt.Printf("Berat Kelinci Terkecil: %.2f kg\n", minweight)
	fmt.Printf("Berat Kelinci Terbesar: %.2f kg\n", maxweight)
}
