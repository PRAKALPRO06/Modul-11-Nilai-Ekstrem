package main

import "fmt"

func cariBeratKelinci(beratKelinci []float64) (float64, float64) {
    // Penanganan kasus array kosong
    if len(beratKelinci) == 0 {
        return 0, 0
    }
    
    // Inisialisasi nilai awal
    beratTerkecil := beratKelinci[0]
    beratTerbesar := beratKelinci[0]
    
    // Iterasi untuk mencari nilai ekstrim
    for i := 1; i < len(beratKelinci); i++ {
        if beratKelinci[i] < beratTerkecil {
            beratTerkecil = beratKelinci[i]
        }
        if beratKelinci[i] > beratTerbesar {
            beratTerbesar = beratKelinci[i]
        }
    }
    
    return beratTerkecil, beratTerbesar
}

func main() {
    // Deklarasi array untuk menyimpan data
    var jumlahData int
    fmt.Print("Masukkan jumlah kelinci: ")
    fmt.Scan(&jumlahData)
    
    // Membuat slice dengan kapasitas sesuai input
    beratKelinci := make([]float64, jumlahData)
    
    // Input data berat kelinci
    for i := 0; i < jumlahData; i++ {
        fmt.Printf("Masukkan berat kelinci ke-%d (kg): ", i+1)
        fmt.Scan(&beratKelinci[i])
    }
    
    // Mencari berat terkecil dan terbesar
    min, max := cariBeratKelinci(beratKelinci)
    
    // Menampilkan hasil
    fmt.Printf("\nHasil Analisis Berat Kelinci:")
    fmt.Printf("\nBerat kelinci terkecil: %.2f kg", min)
    fmt.Printf("\nBerat kelinci terbesar: %.2f kg\n", max)
}