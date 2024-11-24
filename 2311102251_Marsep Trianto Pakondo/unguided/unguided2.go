package main

import "fmt"

type beratIkan [1000]float64

func totalBeratIkan(ikan beratIkan, x, y int) {
	var total float64
	wadah := 1
	fmt.Println("\nTotal berat ikan disetiap wadah : ")

	for i := 1; i <= x; i++ {
		if i % y == 0 {
			total += ikan[i-1]
			fmt.Printf("Wadah ke-%v : %.2f Kg\n", wadah, total)
			wadah++
			total = 0.0
		} else {
			total += ikan[i-1]
			if i == x && i % y != 0 {
				fmt.Printf("Sisanya di wadah ke-%v : %.2f Kg\n", wadah, total)
			}
		}
	}
}

func beratRataIkan(ikan beratIkan, x, y int) {
	var total, rata, bagi float64
	wadah := 1
	fmt.Println("\nRata-rata berat ikan disetiap wadah : ")

	for i := 1; i <= x; i++ {
		if i % y == 0 {
			total += ikan[i-1]
			bagi++
			rata = total / bagi
			fmt.Printf("Wadah ke-%v : %.2f Kg\n", wadah, rata)
			wadah++
			total = 0.0
			rata = 0.0
			bagi = 0.0
		} else {
			total += ikan[i-1]
			bagi++
			if i == x && i % y != 0 {
				rata = total / bagi
				fmt.Printf("Wadah ke-%v : %.2f Kg\n", wadah, rata)
			}
		}
	}
}

func main() {
	var x, y int
	var ikan beratIkan

	fmt.Print("Masukkan banyaknya ikan dijual (maks 1000): ")
	fmt.Scan(&x)
	fmt.Printf("Masukkan banyaknya ikan yang dimasukkan ke dalam wadah (maks sebanyak %v): ", x)
	fmt.Scan(&y)

    if x < 1 || x > 1000 {
		fmt.Println("Jumlah ikan yang dijual harus antara 1 dan 1000.")
		return
	}
    if y < 1 || y > x {
		fmt.Printf("Jumlah ikan yang dimasukkan ke dalam wadah harus antara 1 dan %v.\n", x)
		return
	}

    fmt.Println("\nMasukkan berat ikan :")
    for i := 0; i < x; i++ {
        fmt.Print("Berat ikan ke-", i+1, ": ")
        fmt.Scan(&ikan[i])
    }

	totalBeratIkan(ikan, x, y)
	beratRataIkan(ikan, x, y)
}