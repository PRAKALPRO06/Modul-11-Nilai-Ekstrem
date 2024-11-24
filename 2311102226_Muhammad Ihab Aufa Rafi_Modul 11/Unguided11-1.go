package main

import "fmt"

func main() {
    
    var n_226 int
    var beratKelinci [1000]float64

	fmt.Print("Masukkan Jumlah Kelinci yang Ingin Dihitung Beratnya: ")
    fmt.Scan(&n_226)

    for i := 0; i < n_226; i++ {
		fmt.Printf("Berat Kelinci %d: ", i+1)
        fmt.Scan(&beratKelinci[i])
    }
    
    beratMin := beratKelinci[0]
    beratMax := beratKelinci[0]

    for i := 1; i < n_226; i++ {
        
        if beratKelinci[i] < beratMin {
            beratMin = beratKelinci[i]
        }
		
        if beratKelinci[i] > beratMax {
            beratMax = beratKelinci[i]
        }
    }
    
    fmt.Printf("\nBerat Kelinci Paling Ringan: %.2f\nBerat kelinci Paling Berat: %.2f\n", beratMin, beratMax)
}