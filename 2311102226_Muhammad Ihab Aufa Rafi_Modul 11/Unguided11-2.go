package main

import "fmt"

func main() {
    
    var x_226, y_226 int
    var berat [1000] float64
    var wadah [] float64
    var rerata [] float64
    
    fmt.Print("Masukkan jumlah ikan dan ikan per wadah (x y): ")
    fmt.Scan(&x_226, &y_226)
    
    fmt.Printf("Masukkan berat %d ikan: ", x_226)
    for i := 0; i < x_226; i++ {
        fmt.Scan(&berat[i])
    }
    
    jumlahWadah := (x_226 + y_226 - 1) / y_226
    
    wadah 	= make([]float64, jumlahWadah)
    rerata 	= make([]float64, jumlahWadah)
    
    for i := 0; i < x_226; i++ {
        indeksWadah := i / y_226
        wadah[indeksWadah] += berat[i]
    }
    
    fmt.Println("Total berat ikan di setiap wadah: ")
    for i := 0; i < jumlahWadah; i++ {
        fmt.Printf("%.2f", wadah[i])
        if i < jumlahWadah-1 {
            fmt.Print(" ")
        }
    }
    fmt.Println()
    
    fmt.Println("Rata-rata berat ikan di setiap wadah: ")
    for i := 0; i < jumlahWadah; i++ {
        ikanWadah := y_226
        if i == jumlahWadah-1 && x_226%y_226 != 0 {
            
            ikanWadah = x_226 % y_226
        }
        
        rerata[i] = wadah[i] / float64(ikanWadah)
        fmt.Printf("%.2f", rerata[i])
        if i < jumlahWadah-1 {
            fmt.Print(" ")
        }
    }
    fmt.Println()
}