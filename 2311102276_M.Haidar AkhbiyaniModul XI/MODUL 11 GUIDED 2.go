package main

import "fmt"

// Definisi struct mahasiswa dengan atribut nama, nim, kelas, jurusan, dan ipk
type mahasiswa struct {
	nama, nim, kelas, jurusan string
	ipk                       float64
}

// Definisi tipe data array mahasiswa dengan kapasitas maksimal 2023
type arrMhs [2023]mahasiswa

// Fungsi untuk mencari IPK tertinggi dalam array mahasiswa
func ipk(T arrMhs, n int) float64 {
	var tertinggi float64 = T[0].ipk
	var j int = 1
	for j < n {
		if tertinggi < T[j].ipk {
			tertinggi = T[j].ipk
		}
		j = j + 1
	}
	return tertinggi
}

// Fungsi main untuk mengisi data mahasiswa dan mencari IPK tertinggi
func main() {
	var n int
	var dataMhs arrMhs

	// Meminta input jumlah mahasiswa
	fmt.Print("Masukkan jumlah mahasiswa (maks 2023): ")
	fmt.Scan(&n)

	// Validasi jumlah mahasiswa yang dimasukkan
	if n < 1 || n > 2023 {
		fmt.Println("Jumlah mahasiswa harus antara 1 dan 2023.")
		return
	}

	// Mengisi data mahasiswa
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
		fmt.Print("IPK: ")
		fmt.Scan(&dataMhs[i].ipk)
	}

	// Mencari dan menampilkan IPK tertinggi
	tertinggi := ipk(dataMhs, n)
	fmt.Printf("\nIPK tertinggi dari %d mahasiswa adalah: %.2f\n", n, tertinggi)

}