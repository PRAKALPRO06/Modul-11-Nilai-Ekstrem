package main 

import "fmt"

const NMAX int = 1000
type arrKelinci [NMAX]float64

func isi_data_kelinci(arrBerat *arrKelinci) {
	var n int
	fmt.Print("Masukkan banyak data berat kelinci: ")
	fmt.Scan(&n)
	if(n > NMAX) {
		fmt.Print("Maksimal jumlah data adalah ", NMAX, "\n")
		return
	}
	for i:=0; i<n; i++ {
		fmt.Print("Masukkan berat kelinci ke-", i+1, ": ")
		fmt.Scan(&arrBerat[i])
	}
}

func cetak_berat_min(arrBerat arrKelinci) {
	min := arrBerat[0]
	i := 1

	for arrBerat[i] != 0 && i < 100 {
		if(min > arrBerat[i]) {
			min = arrBerat[i]
		}
		i++
	}
	fmt.Print("Berat kelinci minimum: ", min, " kg\n")
}
func cetak_berat_max(arrBerat arrKelinci) {
	max := arrBerat[0]
	i := 1

	for arrBerat[i] != 0 && i < 100 {
		if(max < arrBerat[i]) {
			max = arrBerat[i]
		}
		i++
	}
	fmt.Print("Berat kelinci maximum: ", max, " kg\n")
}

func main() {
	
	var arrBerat arrKelinci
	isi_data_kelinci(&arrBerat)
	cetak_berat_min(arrBerat)
	cetak_berat_max(arrBerat)
	
}