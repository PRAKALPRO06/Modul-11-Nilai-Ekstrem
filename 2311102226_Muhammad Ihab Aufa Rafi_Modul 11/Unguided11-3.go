package main

import "fmt"

type arrBalita [100]float64

func hitungMinMax(arrBerat arrBalita, bmin, bmax *float64) {
    *bmin = arrBerat[0]
    *bmax = arrBerat[0]
    
    for i := 1; i < len(arrBerat); i++ {
        if arrBerat[i] != 0 { 
            if arrBerat[i] < *bmin {
                *bmin = arrBerat[i]
            }
            if arrBerat[i] > *bmax {
                *bmax = arrBerat[i]
            }
        }
    }
}

func ratarata(arrBerat arrBalita) float64 {
    var total float64
    var count int
    
    for i := 0; i < len(arrBerat); i++ {
        if arrBerat[i] != 0 { 
            total += arrBerat[i]
            count++
        }
    }
    
    if count > 0 {
        return total / float64(count)
    }
    return 0
}

func main() {
    var n int
    var beratBalita arrBalita
    var min_226, max_226 float64
    
    fmt.Print("Masukkan banyak data berat balita : ")
    fmt.Scan(&n)
    
    for i := 0; i < n; i++ {
        fmt.Printf("Masukkan berat balita ke-%d: ", i+1)
        fmt.Scan(&beratBalita[i])
    }
    
    hitungMinMax(beratBalita, &min_226, &max_226)
    
    rerata_226 := ratarata(beratBalita)
    
    fmt.Printf("Berat balita minimum: %.2f kg\n", min_226)
    fmt.Printf("Berat balita maksimum: %.2f kg\n", max_226)
    fmt.Printf("Rata-rata berat balita: %.2f kg\n", rerata_226)
}