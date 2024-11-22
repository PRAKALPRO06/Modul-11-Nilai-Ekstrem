package main 

import (
	"fmt"
)

const NMAX int = 1000
var array_berat_ikan = [NMAX] float64 {}

var jumlah_ikan_di_setiap_wadah int;

func isi_data_berat_ikan(array_berat_ikan *[NMAX]float64, jumlah_ikan_di_setiap_wadah *int) {
	var x, y int

	fmt.Print("Masukkan jumlah ikan dan jumlah ikan di setiap wadah: ")
	fmt.Scan(&x, &y)
	*jumlah_ikan_di_setiap_wadah = y
	
	fmt.Print("Masukkan berat setiap ikan: ")
	for i:=0; i<x; i++ {
		fmt.Scan(&array_berat_ikan[i])
	}
	fmt.Print("\n")
}

func cetak_total_berat_ikan_di_setiap_wadah(array_berat_ikan [NMAX]float64, jumlah_ikan_di_setiap_wadah int) {

	fmt.Println("Total berat ikan di setiap wadah: ")

	wadah_ke := 1
	i := 0

	total_berat_ikan_di_setiap_wadah := 0.0

	for array_berat_ikan[i] != 0.0 {
		total_berat_ikan_di_setiap_wadah += array_berat_ikan[i]
		if((i+1) % jumlah_ikan_di_setiap_wadah == 0) {
			fmt.Print(total_berat_ikan_di_setiap_wadah, " ")
			total_berat_ikan_di_setiap_wadah = 0
			wadah_ke++
		}
		i++
	}
	fmt.Print("\n\n")
}

func cetak_rata_rata_berat_ikan_di_setiap_wadah(array_berat_ikan [NMAX]float64, jumlah_ikan_di_setiap_wadah int) {

	fmt.Println("Rata-rata berat ikan di setiap wadah: ")

	wadah_ke := 1
	i := 0

	total_berat_ikan_di_setiap_wadah := 0.0
	for array_berat_ikan[i] != 0.0 {
		total_berat_ikan_di_setiap_wadah += array_berat_ikan[i]
		if((i+1) % jumlah_ikan_di_setiap_wadah == 0) {
			fmt.Print(total_berat_ikan_di_setiap_wadah/float64(jumlah_ikan_di_setiap_wadah), " ")
			total_berat_ikan_di_setiap_wadah = 0.0
			wadah_ke++
		}
		i++
	}
	fmt.Print("\n\n")
}


func main() {

	isi_data_berat_ikan(&array_berat_ikan, &jumlah_ikan_di_setiap_wadah)
	cetak_total_berat_ikan_di_setiap_wadah(array_berat_ikan, jumlah_ikan_di_setiap_wadah)
	cetak_rata_rata_berat_ikan_di_setiap_wadah(array_berat_ikan, jumlah_ikan_di_setiap_wadah)

}