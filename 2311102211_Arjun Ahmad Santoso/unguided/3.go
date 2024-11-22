package main 

import "fmt"

type arrBalita [100]float64

func isi_data_balita(arrBerat *arrBalita) {
	var n int
	fmt.Print("Masukkan banyak data berat balita: ")
	fmt.Scan(&n)
	if(n > 100) {
		fmt.Print("Maksimal jumlah data adalah 100\n")
		return
	}
	for i:=0; i<n; i++ {
		fmt.Print("Masukkan berat balita ke-", i+1, ": ")
		fmt.Scan(&arrBerat[i])
	}
}

func cetak_berat_min(arrBerat arrBalita) {
	min := arrBerat[0]
	i := 1

	for arrBerat[i] != 0 && i < 100 {
		if(min > arrBerat[i]) {
			min = arrBerat[i]
		}
		i++
	}
	fmt.Print("Berat balita minimum: ", min, " kg\n")
}
func cetak_berat_max(arrBerat arrBalita) {
	max := arrBerat[0]
	i := 1

	for arrBerat[i] != 0 && i < 100 {
		if(max < arrBerat[i]) {
			max = arrBerat[i]
		}
		i++
	}
	fmt.Print("Berat balita maximum: ", max, " kg\n")
}
func cetak_rerata_berat(arrBerat arrBalita) {
	sum := arrBerat[0]
	i := 1

	for arrBerat[i] != 0 && i < 100 {
		sum += arrBerat[i]
		i++
	}
	fmt.Print("Rerata berat balita: ", sum/float64(i), " kg\n")
}

func main() {
	
	var arrBerat arrBalita
	isi_data_balita(&arrBerat)
	cetak_berat_min(arrBerat)
	cetak_berat_max(arrBerat)
	cetak_rerata_berat(arrBerat)
}