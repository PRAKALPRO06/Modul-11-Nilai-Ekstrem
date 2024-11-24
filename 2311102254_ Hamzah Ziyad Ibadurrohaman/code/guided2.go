package main

import "fmt"

// Definisi struct mahasiswa dengan atribut nama, nim, kelas, jurusan, dan ipk
type mahasiswa struct {
	nama, nim, kelas, jurusan string
	ipk                       float64
}

// Definisi tipe data array mahasiswa dengan kapasitas maksimal 2023
type arrMhs [2023]mahasiswa

// Fungsi untuk mencari indeks IPK tertinggi dalam array mahasiswa
func indeksIPKTertinggi(T arrMhs, n int) int {
	var idx int = 0 // Inisialisasi indeks IPK tertinggi pada indeks pertama
	for j := 1; j < n; j++ {
		if T[idx].ipk < T[j].ipk {
			idx = j // Update indeks jika ditemukan IPK yang lebih tinggi
		}
	}
	return idx
}

// Fungsi main untuk mengisi data mahasiswa dan mencari IPK tertinggi beserta indeksnya
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

	// Mendapatkan indeks IPK tertinggi
	idxTertinggi := indeksIPKTertinggi(dataMhs, n)

	// Menampilkan data mahasiswa dengan IPK tertinggi
	fmt.Printf("\nMahasiswa dengan IPK tertinggi:\n")
	fmt.Printf("Nama: %s\n", dataMhs[idxTertinggi].nama)
	fmt.Printf("NIM: %s\n", dataMhs[idxTertinggi].nim)
	fmt.Printf("Kelas: %s\n", dataMhs[idxTertinggi].kelas)
	fmt.Printf("Jurusan: %s\n", dataMhs[idxTertinggi].jurusan)
	fmt.Printf("IPK: %.2f\n", dataMhs[idxTertinggi].ipk)
}
