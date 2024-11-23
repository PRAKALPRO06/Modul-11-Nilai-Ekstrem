<h2 align="center"><strong>LAPORAN PRAKTIKUM</strong></h2>
<h2 align="center"><strong>ALGORITMA DAN PEMROGRAMAN 2</strong></h2>

<br>

<h2 align="center"><strong>MODUL XI</strong></h2>
<h2 align="center"><strong> PENCARIAN NILAI EKSTRIM PADA HIMPUNAN DATA </strong></h2>

<br>

<p align="center">

  <img src="https://github.com/user-attachments/assets/0a03461e-7740-4661-9e83-9925031bd72c" alt="Logo" width="200"/>

</p>

<br>

<p align="center">
  <strong>Disusun Oleh:</strong><br>
  LAILATUR RAHMAH / 2311102177<br>
  S1 IF 11 06
</p>

<br>

<p align="center">
  <strong>Dosen Pengampu:</strong><br>
  Abednego Dwi Septiadi
  
</p>

<br>

<p align="center">
  <strong>PROGRAM STUDI S1 TEKNIK INFORMATIKA</strong><br>
  <strong>FAKULTAS INFORMATIKA</strong><br>
  <strong>TELKOM UNIVERSITY PURWOKERTO</strong><br>
  <strong>2024</strong>
</p>

------

## Features

- [DASAR TEORI](#dasar-teori)
- [GUIDED](#guided)
- [UNGUIDED](#unguided)

## DASAR TEORI


## GUIDED

### 1. Mencari Nilai Terkecil dalam Kumpulan Data Angka

#### Source Code

```go
package main

import "fmt"

// Mendeklarasikan tipe data array arrInt dengan panjang 2023
type arrInt [2023]int

// Fungsi untuk mencari elemen terkecil dalam array
func terkecil(tabInt arrInt, n int) int {
	var min int = tabInt[0]
	var j int = 1
	for j < n {
		if min > tabInt[j] {
			min = tabInt[j]
		}
		j = j + 1
	}
	return min
}

// Fungsi main
func main() {
	var n int
	var tab arrInt

	// Meminta input jumlah elemen array
	fmt.Print("Masukkan jumlah elemen (maks 2023): ")
	fmt.Scan(&n)

	// Validasi input jumlah elemen
	if n < 1 || n > 2023 {
		fmt.Println("Jumlah elemen harus antara 1 dan 2023.")
		return
	}

	// Memasukkan elemen-elemen array
	fmt.Println("Masukkan elemen-elemen array:")
	for i := 0; i < n; i++ {
		fmt.Print("Elemen ke-", i+1, ": ")
		fmt.Scan(&tab[i])
	}

	// Memanggil fungsi terkecil untuk menemukan nilai terkecil
	minVal := terkecil(tab, n)

	// Menampilkan nilai terkecil
	fmt.Println("Nilai terkecil dalam array adalah:", minVal)
}
```

#### Screenshoot Output

![image](https://github.com/user-attachments/assets/03bc6b31-67dc-439b-b846-1923936dd44e)

#### Deskripsi Program

Program di atas bertujuan untuk mencari elemen terkecil dalam sebuah array dengan maksimal 2023 elemen menggunakan bahasa Go. Program dimulai dengan mendeklarasikan tipe data array arrInt untuk menyimpan bilangan integer hingga panjang 2023. Pengguna diminta memasukkan jumlah elemen array (n) dengan validasi agar nilainya berada antara 1 hingga 2023. Selanjutnya, pengguna menginput nilai-nilai elemen array satu per satu. Program kemudian memanfaatkan fungsi terkecil untuk iterasi dan membandingkan setiap elemen dalam array, guna menemukan nilai terkecil. Hasil berupa elemen terkecil dari array tersebut ditampilkan kepada pengguna.

### 2. Penentuan IPK Tertinggi Mahasiswa dalam Suatu Data Kelas

#### Source Code

```go
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
```

#### Screenshoot Output

![image](https://github.com/user-attachments/assets/e73df62e-cdea-43dd-ab62-ff7c844067af)

#### Deskripsi Program

Program ini adalah aplikasi sederhana untuk mengelola data mahasiswa, termasuk nama, NIM, kelas, jurusan, dan IPK. Pengguna diminta untuk memasukkan jumlah mahasiswa, kemudian mengisi data setiap mahasiswa satu per satu. Setelah semua data dimasukkan, program akan mencari dan menampilkan IPK tertinggi di antara seluruh mahasiswa yang telah diinputkan. Program juga memvalidasi jumlah mahasiswa yang diinput untuk memastikan berada dalam rentang yang diperbolehkan (1 hingga 2023). Dengan menggunakan tipe data struct untuk representasi mahasiswa dan array untuk menyimpan data, program ini memberikan solusi sederhana untuk analisis IPK mahasiswa.

## UNGUIDED

### UNGUIDED 1

#### Study Case

Sebuah program digunakan untuk mendata berat anak kelinci yang akan dijual ke pasar. Program ini menggunakan array dengan kapasitas 1000 untuk menampung data berat anak kelinci yang akan dijual.

**Masukan:**

Terdiri dari sekumpulan bilangan, yang mana bilangan pertama adalah bilangan bulat N yang menyatakan banyaknya anak kelinci yang akan ditimbang beratnya. Selanjutnya N bilangan riil berikutnya adalah berat dari anak kelinci yang akan dijual.

**Keluaran:**

Terdiri dari dua buah bilangan riil yang menyatakan berat kelinci terkecil dan terbesar.

#### Source Code

```go
// LAILATUR RAHMAH
// 2311102177
// IF 11 06

package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	for {
		fmt.Println(strings.Repeat("=", 40))
		fmt.Println("    Program Pengolahan Berat Kelinci")
		fmt.Println(strings.Repeat("=", 40))

		// Input: Jumlah kelinci
		var n_2311102177 int
		fmt.Print("\nMasukkan jumlah kelinci: ")
		fmt.Scanln(&n_2311102177)

		// Memastikan jumlah kelinci lebih dari 0
		if n_2311102177 <= 0 {
			fmt.Println("\n[!] Jumlah kelinci harus lebih dari 0. Silakan coba lagi.")
			continue
		}

		// Membuat slice untuk menyimpan berat kelinci
		weights := make([]float64, n_2311102177)
		fmt.Println("\nMasukkan berat masing-masing kelinci:")
		for i := 0; i < n_2311102177; i++ {
			fmt.Printf("  - Berat kelinci %d: ", i+1)
			fmt.Scanln(&weights[i])
		}

		// Mengurutkan berat kelinci dalam urutan menaik
		sort.Float64s(weights)

		// Output: Berat kelinci terkecil dan terbesar
		fmt.Println("\n" + strings.Repeat("-", 40))
		fmt.Println("        Hasil Pengolahan")
		fmt.Println(strings.Repeat("-", 40))
		fmt.Printf(" Berat terkecil : %.2f kg\n", weights[0])
		fmt.Printf(" Berat terbesar : %.2f kg\n", weights[n_2311102177-1])
		fmt.Println("\n Daftar berat kelinci (terurut):")
		for i, weight := range weights {
			fmt.Printf("  %2d. %.2f kg\n", i+1, weight)
		}
		fmt.Println(strings.Repeat("-", 40))

		// Menanyakan apakah pengguna ingin mengulangi
		var again string
		fmt.Print("\n Apakah Anda ingin memasukkan data lagi? (yes/no): ")
		fmt.Scanln(&again)

		// Jika pengguna tidak ingin mengulangi, keluar dari loop
		if strings.ToLower(again) != "ya" {
			fmt.Println("\n Terima kasih telah menggunakan program ini!")
			break
		}
	}
}
```

#### Screenshoot Output

![image](https://github.com/user-attachments/assets/578b5e98-93d0-49eb-b984-59e0d1184f2a)

#### Alur Program

* Inisialisasi Program
  * Program dimulai dengan menampilkan header berisi informasi tentang program.
* Meminta Input Jumlah Kelinci
  * Pengguna diminta memasukkan jumlah kelinci (n).
  * Jika jumlah kelinci yang dimasukkan kurang dari atau sama dengan nol, program meminta pengguna untuk memasukkan jumlah yang valid.
* Memasukkan Berat Kelinci
  * Pengguna diminta untuk memasukkan berat badan masing-masing kelinci sebanyak jumlah yang telah dimasukkan.
* Pengolahan Data Berat Kelinci
  * Program mengurutkan berat kelinci secara menaik menggunakan fungsi bawaan sort.Float64s.
* Menampilkan Hasil Pengolahan, program menampilkan:
  * Berat kelinci terkecil (elemen pertama dalam slice setelah diurutkan).
  * Berat kelinci terbesar (elemen terakhir dalam slice setelah diurutkan).
  * Daftar berat kelinci yang telah diurutkan dalam bentuk daftar.
  * Menanyakan Pengguna untuk Melanjutkan atau Berhenti
* Program menanyakan apakah pengguna ingin mengulangi proses atau keluar.
  * Jika pengguna memilih untuk mengulangi (yes/ya), proses kembali ke langkah 2.
  * Jika pengguna memilih untuk berhenti, program menampilkan pesan penutup dan berhenti.

#### Deskripsi Program

Program ini adalah aplikasi untuk mengolah data berat badan kelinci. Pengguna diminta memasukkan jumlah kelinci yang akan diinputkan dan berat masing-masing kelinci. Program kemudian mengurutkan berat badan kelinci secara menaik, menampilkan berat terkecil, terbesar, serta daftar berat kelinci yang telah terurut. Setelah itu, pengguna diberi pilihan untuk mengulangi proses atau mengakhiri program. Program ini menggunakan slice untuk menyimpan data berat kelinci dan fungsi bawaan sort.Float64s untuk pengurutan, memastikan pengolahan data lebih efisien dan mudah dipahami.

### UNGUIDED 2

#### Study Case

Sebuah program digunakan untuk menentukan tarif ikan yang akan dijual ke pasar. Program ini menggunakan array dengan kapasitas 1000 untuk menampung data berat ikan yang akan dijual.

**Masukan:**

Terdiri dari dua baris:

* Baris pertama: Terdiri dari dua bilangan bulat x dan y.
    * x: Menyatakan banyaknya ikan yang akan dijual.
    * y: Menyatakan banyaknya ikan yang akan dimasukkan ke dalam wadah.
* Baris kedua: Terdiri dari sejumlah x bilangan riil yang menyatakan banyaknya ikan yang akan dijual.

**Keluaran:**

Terdiri dari dua baris:

* Baris pertama: Kumpulan bilangan riil yang menyatakan total berat ikan di setiap wadah (jumlah wadah tergantung pada nilai x dan y, urutan ikan yang dimasukkan ke dalam wadah sesuai urutan pada masukan baris ke-2).
* Baris kedua: Sebuah bilangan riil yang menyatakan berat rata-rata ikan di setiap wadah.

#### Source Code

```go
// LAILATUR RAHMAH
// 2311102177
// IF 11 06

package main

import (
	"fmt"
	"strings"
)

func main() {
	// Input: Masukan
	var x_2311102177, y int
	var ikan []float64

	// Meminta pengguna untuk memasukkan jumlah ikan
	fmt.Println(strings.Repeat("=", 40))
	fmt.Println("    Program Pengolahan Berat Ikan")
	fmt.Println(strings.Repeat("=", 40))

	fmt.Print("\nMasukkan jumlah ikan yang akan dijual (x): ")
	fmt.Scanln(&x_2311102177)
	if x_2311102177 <= 0 {
		fmt.Println("\n[!] Jumlah ikan harus positif.")
		return // Keluar dari program jika jumlah ikan tidak valid
	}

	// Meminta pengguna untuk memasukkan jumlah wadah
	fmt.Print("\nMasukkan jumlah wadah (y): ")
	fmt.Scanln(&y)
	if y <= 0 {
		fmt.Println("\n[!] Jumlah wadah harus positif.")
		return // Keluar dari program jika jumlah wadah tidak valid
	}

	// Meminta pengguna untuk memasukkan berat setiap ikan satu per satu
	fmt.Println("\nMasukkan berat setiap ikan (dalam kg):")
	for i := 0; i < x_2311102177; i++ {
		var berat float64
		fmt.Printf("  - Berat ikan %d: ", i+1)
		_, err := fmt.Scanln(&berat)
		if err != nil {
			fmt.Println("\n[!] Input berat ikan tidak valid.")
			return // Keluar dari program jika ada kesalahan input
		}
		ikan = append(ikan, berat)
	}

	// Hitung total berat ikan di setiap wadah
	wadah := make([]float64, y) // Membuat slice untuk menyimpan total berat di setiap wadah
	for i := 0; i < x_2311102177; i++ {
		wadah[i%y] += ikan[i] // Menambahkan berat ikan ke wadah yang sesuai
	}

	// Hitung rata-rata berat ikan di setiap wadah
	rataRata := 0.0
	for _, v := range wadah {
		rataRata += v // Menjumlahkan total berat di semua wadah
	}
	rataRata /= float64(x_2311102177) // Rata-rata berdasarkan jumlah ikan

	// Output: Keluaran
	fmt.Println("\n" + strings.Repeat("-", 40))
	fmt.Println("     Hasil Pengolahan Berat Ikan")
	fmt.Println(strings.Repeat("-", 40))

	// Menampilkan total berat ikan di setiap wadah
	fmt.Println("\nTotal berat ikan di setiap wadah:")
	for i, v := range wadah {
		fmt.Printf("  Wadah %d: %.2f kg\n", i+1, v) // Menampilkan total berat ikan di setiap wadah
	}

	// Menampilkan berat rata-rata ikan di setiap wadah
	fmt.Println("\nBerat rata-rata ikan di setiap wadah:")
	fmt.Printf("  %.2f kg\n", rataRata) // Menampilkan rata-rata berat ikan

	// Akhir dari program
	fmt.Println("\n" + strings.Repeat("=", 40))
	fmt.Println("Terima kasih telah menggunakan program ini!")
	fmt.Println(strings.Repeat("=", 40))
}
```

#### Screenshoot Output

![image](https://github.com/user-attachments/assets/79059efc-e645-4b32-b45a-5fbb89cc2fb9)

#### Alur Program

* Inisialisasi Program
    * Program menampilkan header untuk menjelaskan tujuan dari aplikasi.
* Input Data Pengguna
    * Pengguna memasukkan jumlah ikan (x).
    * Jika jumlah ikan kurang dari atau sama dengan nol, program menghentikan proses dan memberikan peringatan.
    * Pengguna memasukkan jumlah wadah (y). Jika jumlah wadah kurang dari atau sama dengan nol, program juga menghentikan proses dengan peringatan.
* Input Berat Ikan
    * Pengguna diminta untuk memasukkan berat setiap ikan satu per satu.
    * Jika ada kesalahan input (berat tidak valid), program akan berhenti.
* Proses Distribusi Berat Ikan ke Wadah
    * Berat ikan didistribusikan ke dalam wadah menggunakan indeks modulo (i % y), sehingga ikan dimasukkan secara berurutan ke wadah secara siklis.
* Menghitung Total dan Rata-rata Berat
    * Program menjumlahkan total berat ikan di setiap wadah.
    * Menghitung rata-rata berat ikan per wadah berdasarkan jumlah total ikan.
* Menampilkan Hasil Pengolahan
    * Program menampilkan total berat ikan di setiap wadah.
    * Menampilkan rata-rata berat ikan per wadah.
* Akhir Program
    * Program menampilkan pesan penutup setelah semua hasil ditampilkan.

#### Deskripsi Program

Program ini digunakan untuk mengolah berat ikan yang akan dibagi ke dalam beberapa wadah. Pengguna memasukkan jumlah ikan dan jumlah wadah, lalu memasukkan berat masing-masing ikan. Program menghitung total berat ikan di setiap wadah secara merata menggunakan pendekatan modulo untuk distribusi. Selain itu, program juga menghitung dan menampilkan rata-rata berat ikan per wadah. Hasilnya adalah distribusi berat ikan per wadah serta nilai rata-rata berat ikan.

### UNGUIDED 3

### Study Case

Pos Pelayanan Terpadu (posyandu) sebagai tempat pelayanan kesehatan perlu mencatat data berat balita (dalam kg). Petugas akan memasukkan data tersebut 1  ke dalam array. Dari data yang diperoleh akan dicari berat balita terkecil, terbesar, dan reratanya.   
Buatlah program dengan spesifikasi subprogram sebagai berikut:

```
type arrBalita [100]float64

func hitungMinMax(arrBerat arrBalita, bMin, bMax *float64) {
  /* I.S. Terdefinisi array dinamis arrBerat
     Proses: Menghitung berat minimum dan maksimum dalam array
     F.S. Menampilkan berat minimum dan maksimum balita */
  // ... isi fungsi di sini
}

function rerata(arrBerat arrBalita) real {
  /* menghitung dan mengembalikan rerata berat balita dalam array */
  // ... isi fungsi di sini
}
```

Perhatikan sesi interaksi pada contoh berikut ini (teks bergaris bawah adalah input/read)

```
Masukan banyak data berat balita: 4
Masukan berat balita ke-1: 5.3
Masukan berat balita ke-2: 6.2
Masukan berat balita ke-3: 4.1
Masukan berat balita ke-4: 9.9
Berat balita minimum: 4.10 kg
Berat balita maksimum: 9.90 kg
Rerata berat balita: 6.38 kg
```

#### Source Code

```go
// LAILATUR RAHMAH
// 2311102177
// IF 11 06

package main

import (
	"fmt"
)

// Definisi tipe array untuk berat balita
type arrBalita [100]float64

// Subprogram untuk menghitung berat minimum dan maksimum
func hitungMinMax(arrBerat arrBalita, n int, min, max *float64) {
	*min = arrBerat[0]
	*max = arrBerat[0]
	for i := 1; i < n; i++ {
		if arrBerat[i] < *min {
			*min = arrBerat[i]
		}
		if arrBerat[i] > *max {
			*max = arrBerat[i]
		}
	}
}

// Subprogram untuk menghitung rata-rata berat balita
func rataRata(arrBerat arrBalita, n int) float64 {
	total := 0.0
	for i := 0; i < n; i++ {
		total += arrBerat[i]
	}
	return total / float64(n)
}

func main() {
	var n_2311102177 int
	var berat arrBalita
	var min, max float64

	// Input jumlah data balita
	fmt.Print("Masukkan banyak data berat balita: ")
	fmt.Scan(&n_2311102177)

	// Validasi jumlah data
	if n_2311102177 <= 0 || n_2311102177 > 100 {
		fmt.Println("Jumlah data balita harus antara 1 dan 100.")
		return
	}

	// Input berat balita
	for i := 0; i < n_2311102177; i++ {
		fmt.Printf("Masukkan berat balita ke-%d: ", i+1)
		fmt.Scan(&berat[i])
	}

	// Hitung nilai minimum, maksimum, dan rata-rata
	hitungMinMax(berat, n_2311102177, &min, &max)
	rata := rataRata(berat, n_2311102177)

	// Output hasil
	fmt.Printf("Berat balita minimum: %.2f kg\n", min)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", max)
	fmt.Printf("Rata-rata berat balita: %.2f kg\n", rata)
}
```

#### Screenshoot Output

![image](https://github.com/user-attachments/assets/cc732384-db41-4777-9060-6941740eee3e)

#### Alur Program

* Inisialisasi Program
    * Program dimulai dengan deklarasi variabel untuk menyimpan jumlah balita, data berat balita, nilai minimum, maksimum, dan rata-rata berat balita.
* Input Jumlah Data Balita
    * Pengguna diminta untuk memasukkan jumlah balita (n).
    * Jika jumlah balita tidak valid (kurang dari 1 atau lebih dari 100), program memberikan pesan peringatan dan berhenti.
* Input Berat Balita
    * Pengguna diminta memasukkan berat balita satu per satu.
    * Data berat setiap balita disimpan dalam array berat.
* Menghitung Berat Minimum dan Maksimum
    * Subprogram hitungMinMax digunakan untuk menentukan berat balita terkecil (min) dan terbesar (max) dalam array.
* Menghitung Rata-Rata Berat Balita
    * Subprogram rataRata menghitung rata-rata berat balita berdasarkan jumlah total berat dan jumlah balita.
* Menampilkan Hasil Perhitungan
    * Program menampilkan berat minimum, maksimum, dan rata-rata balita dengan format desimal hingga dua angka di belakang koma.
* Akhir Program
    * Setelah hasil ditampilkan, program berhenti.

#### Deskripsi Program

Program ini dirancang untuk mengolah data berat badan balita. Pengguna diminta untuk memasukkan jumlah balita dan berat badan masing-masing balita. Program kemudian menghitung nilai minimum, maksimum, dan rata-rata berat balita dari data yang diberikan. Hasilnya ditampilkan dalam format yang mudah dipahami, sehingga mempermudah pengguna dalam menganalisis berat badan balita.
