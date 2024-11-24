package main

import "fmt"

type berat [1000]float64

func beratTerkecil(b berat, n int) float64 {
    var min float64 = b[0]
    for i := 1; i < n; i++ {
        if min > b[i] {
            min = b[i]
        }
    }

    return min
}

func beratTerbesar(b berat, n int) float64 {
    var maks float64 = b[0]
    for i := 1; i < n; i++ {
        if maks < b[i] {
            maks = b[i]
        }
    }

    return maks
}

func main()  {
    var b berat
    var n int

    fmt.Print("Masukkan jumlah anak kelinci (maks 1000): ")
	fmt.Scan(&n)

    if n < 1 || n > 1000 {
		fmt.Println("Jumlah mahasiswa harus antara 1 dan 2023.")
		return
	}

    fmt.Println("Masukkan berat dari anak kelinci :")
    for i := 0; i < n; i++ {
        fmt.Print("Anak Kelinci ke-", i+1, ": ")
        fmt.Scan(&b[i])
    }

    terkecil := beratTerkecil(b, n)
    terbesar := beratTerbesar(b, n)

    fmt.Printf("Nilai berat kelinci terkecil : %.2f Kg\n", terkecil)
    fmt.Printf("Nilai berat kelinci terbesar : %.2f Kg\n", terbesar)

}