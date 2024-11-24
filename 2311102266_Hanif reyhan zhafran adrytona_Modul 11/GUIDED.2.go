// 2311102266_Hanif Reyhan Zhafran Arytona

package main

import "fmt"

type mahasiswa struct {
	nama, nim, kelas, jurusan string
	ipk                       float64
}

type arrMhs [2023]mahasiswa

func ipkTertinggi(T arrMhs, n int) float64 {
	var tertinggi float64 = T[0].ipk
	for j := 1; j < n; j++ {
		if tertinggi < T[j].ipk {
			tertinggi = T[j].ipk
		}
	}
	return tertinggi
}

func main() {
	var n int
	var dataMhs arrMhs

	fmt.Print("Masukkan jumlah mahasiswa (maks 2023): ")
	fmt.Scan(&n)

	if n < 1 || n > 2023 {
		fmt.Println("Jumlah mahasiswa harus antara 1 dan 2023.")
		return
	}

	for i := 0; i < n; i++ {
		fmt.Printf("\nMasukkan data mahasiswa ke-%d\n", i+1)
		fmt.Print("Nama: ")
		fmt.Scan(&dataMhs[i].nama)
		fmt.Print("NIM: ")
		fmt.Scan(&dataMhs[i].nim)
		fmt.Print("Kelas: ")
		fmt.Scan(&dataMhs[i].kelas)
		fmt.Print("Jurusan: ")
		fmt.Scan(&dataMhs[i].jurusan)
		fmt.Print("IPK (0.0 - 4.0): ")
		fmt.Scan(&dataMhs[i].ipk)

		if dataMhs[i].ipk < 0.0 || dataMhs[i].ipk > 4.0 {
			fmt.Println("IPK tidak valid. Harus antara 0.0 dan 4.0.")
			return
		}
	}

	tertinggi := ipkTertinggi(dataMhs, n)
	fmt.Printf("\nIPK tertinggi dari %d mahasiswa adalah: %.2f\n", n, tertinggi)
}
