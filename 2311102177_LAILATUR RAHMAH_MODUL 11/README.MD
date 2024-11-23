<h2 align="center"><strong>LAPORAN PRAKTIKUM</strong></h2>
<h2 align="center"><strong>ALGORITMA DAN PEMROGRAMAN 2</strong></h2>

<br>

<h2 align="center"><strong>MODUL V</strong></h2>
<h2 align="center"><strong> REKURSIF </strong></h2>

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

## Unguided

### Study Case

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

* Inisialisasi Program:
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

### Study Case

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


```

#### Screenshoot Output

![image](https://github.com/user-attachments/assets/79059efc-e645-4b32-b45a-5fbb89cc2fb9)

#### Alur Program


#### Deskripsi Program


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


```

#### Screenshoot Output

![image](https://github.com/user-attachments/assets/cc732384-db41-4777-9060-6941740eee3e)

#### Alur Program


#### Deskripsi Program
