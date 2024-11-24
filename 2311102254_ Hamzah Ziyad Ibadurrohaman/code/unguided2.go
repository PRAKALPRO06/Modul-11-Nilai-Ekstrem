package main

import "fmt"

func main() {
	var x, y int
	var beratikan [1000]float64

	fmt.Print("Jumlah ikan dijual berapa?: (x) dan jumlah ikan per wadah ?(y) : ")
	fmt.Scan(&x, &y)

	if x < 1 || x > 1000 || y < 1 {
		fmt.Println("jumlah nya harus dari 1 hingga 1000!. jumlah minimum wadah itu 1.")
	}

	fmt.Println("berat ikan: ")
	for i := 0; i < x; i++ {
		fmt.Printf("berat ikan %d: ", i+1)
		fmt.Scan(&beratikan[i])
	}

	var totalberat []float64
	var total float64
	for i := 0; i < x; i++ {
		total += beratikan[i]
		if (i+1)%y == 0 || i == x-1 {
			totalberat = append(totalberat, total)
			total = 0
		}
	}

	var rata float64
	for _, berat := range totalberat {
		rata += berat
	}
	rata /= float64(len(totalberat))

	fmt.Println("total berat setiap wadah: ")
	for i, berat := range totalberat {
		fmt.Printf("wadah %d: %.2f\n", i+1, berat)

	}
	fmt.Printf("berat rata-rata ikan setiap wadah: %.2f\n", rata)

}
