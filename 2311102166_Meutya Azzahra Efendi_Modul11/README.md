<h1 align="center">LAPORAN PRAKTIKUM</h1>
<h2 align="center">ALGORITMA DAN PEMROGRAMAN 2</h2>

<h3 align="center">MODUL 10</h3>
<h4 align="center">PENCARIAN NILAI EKSTRIM PADA HIMPUNAN DATA</h4>

<p align="center">
  <img src="telkomuniv.png" alt="Logo Telkom University" width="200">
</p>
<p align="center">
    <strong>Disusun Oleh:</strong><br>
    Meutya Azzahra Efendi / 2311102166<br>
</p>

<p align="center">
    <strong>Dosen Pengampu:</strong><br>
    Abednego Dwi Septiadi
</p>

<p align="center">
    PROGRAM STUDI S1 TEKNIK INFORMATIKA<br>
    FAKULTAS INFORMATIKA<br>
    TELKOM UNIVERSITY PURWOKERTO<br>
    2024
</p>


### I. DASAR TEORI

### 10.1 Ide Pencarian Nilai Max/Min

Pencarian adalah suatu proses yang lazim dilakukan di dalam kehidupan sehari-hari. Contoh pengunaannya dalam kehidupan nyata sangat beragam, misalnya pencarian fike di dalam directory komputer, pencarian suatu teks di dalam sebuah dokumen, pencarian buku pada rak buku, dan contoh lainnya. Pertama pada modul ini aka dipelajari salah satu algoritma pencarian nilai terkecil atau terbesar pada sekumpulan data atau biasa disebut pencarian nilai ekstrim.
Ide algoritma sederhana sekali. Karena data harus diproses secara sekuensial, maka nilai atau indeks ke nilai maksimum dari data yang telah diproses disimpan untuk dibandingkan dengan data berikutnya. Nilai yang berhasil disimpan sampai algoritma tersebut berakhir adalah nilai maksimum yang dicari. Adapun algoritmanya secara umum adalah sebagai berikut:
1. Jadikan data pertama sebagai nilai ekstrim
2. Lakukan validasi nilai ekstrim dari data kedua hingga data terakhir
   - Apabila nilai ekstrim tidak valid, maka update nilai ekstrims tersebut dengan data yang dicek
3. Apabila semua data telah dicek, maka nilai ekstrim yang dimiliki adalah valid

Berikut ini adalah notasi dalam pseudocode dan bahasa Go, misalnya untuk pencarian nilai terbesar atau maksimum:
| Notasi Algoritma | Notasi dalam bahasa Go |
|---|---|
| max ← 1 | max := 0 |
| i ← 2 | i := 1 |
| while i <= n do | for i < n { |
|  if a[i] > a[max] then |  if a[i] > a[max] { |
|   max ← i |   max = i |
|  endif |  } |
|  i ← i + 1 |  i = i + 1 |
| endwhile | } |

### 10.2 Pencarian Nilai Ekstrim pada Array Bertipe Data Dasar

Misalnya terdefinisi sebuah array of integer dengan kapasitas 2023 dan array terisi sejumlah N bilangan bulat, kemudian pencarian nilai terkecil dilakukan pada array tersebut. Perhatikan potongan program dalam bahasa Go berikut ini!

```
type arrInt [2023]int

func terkecil_1(tabInt arrInt, n int) int {
    /* mengembalikan nilai terkecil yang terdapat di dalam tabInt yang berisi n
    bilangan bulat */

    var min int = tabInt[0]  // min berisi data pertama
    var j int = 1

    for j < n {
        if min > tabInt[j] {  // pengecekan apakah nilai minimum valid
            min = tabInt[j]  // update nilai minimum dengan yang valid
        }
        j = j + 1
    }

    return min  // returnkan nilai minimumnya
}
```
Potongan program diatas sedikit berbeda dengan sebelumnya karena penggunaan indeks array pada bahasa Go dimulai dari nol atau "0". Selanjutnya, pada penjelasan di awal bab 3 telah disampaikan bahwa pada pencarian yang terpenting adalah posisi atau indeks dari nilai yang dicari dalam kumpulan data atau array. Oleh karena itu modifikasi pada program di atas dapat dilihat pada potongan program berikut ini!

```
type arrInt [2023]int

func terkecil_2(tabInt arrInt, n int) int {
    /* mengembalikan indeks nilai terkecil yang terdapat di dalam tabInt yang berisi 
    n bilangan bulat */

    var idx int = 0  // idx berisi indeks data pertama
    var j int = 1

    for j < n {
        if tabInt[idx] > tabInt[j] {  // pengecekan apakah nilai minimum valid
            idx = j  // update nilai minimum dengan yang valid
        }
        j = j + 1
    }

    return idx  // returnkan indeks nilai minimumnya
}
```

### 10.3 Pencarian Nilai Ekstrim pada Array Bertipe Data Terstruktur

Pada kasus yang lebih kompleks pencarian ekstrim dapat juga dilakukan, misalnya mencari data mahasiswa dengan nilai terbesar, mencari lagu dengan durasi terlama, mencari pembalap yang memiliki catatan waktu balap tercepat dan sebagainya. Sebagai contoh misalnya terdapat array yang digunakan untuk menyimpan data mahasiswa, kemudian terdapat fungsi IPK yang digunakan untuk mencari data mahasiswa dengan IPK tertinggi.
```
type mahasiswa struct {
    nama, nim, kelas, jurusan string
    ipk float64
}

type arrMhs [2023]mahasiswa

func IPK_1(T arrMhs, n int) float64 {
    /* mengembalikan ipk terkecil yang dimiliki mahasiswa pada array T yang berisi
    n mahasiswa */

    var tertinggi float64 = T[0].ipk
    var j int = 1

    for j < n {
        if tertinggi <= T[j].ipk {
            tertinggi = T[j].ipk
        }
        j = j + 1
    }

    return tertinggi
}
```
Apabila diperhatikan potongan program diatas, maka kita akan memperoleh ipk tertinggi, tetapi kita tidak memperoleh identitas mahasiswa dengan ipk tertinggi tersebut. Maka seperti penjelasan yang sudah diberikan sebelumnya, maka pencarian yang dilakukan bisa mengembalikan indeks mahasiswa dengan ipk tertinggi tersebut. Berikut ini adalah modifikasinya!
```
type mahasiswa struct {
    nama, nim, kelas, jurusan string
    ipk float64
}

type arrMhs [2023]mahasiswa

func IPK_2(T arrMhs, n int) int {
    /* mengembalikan indeks mahasiswa yang memiliki ipk tertinggi pada array T yang 
    berisi n mahasiswa */

    var idx int = 0
    var j int = 1

    for j < n {
        if T[idx].ipk < T[j].ipk {
            idx = j
        }
        j = j + 1
    }

    return idx
}
```
Sehingga melalui algoritma diatas, identitas mahasiswa dapat diperoleh, misalnya T[idx].nama, T[idx].nim, T[idx].kelas, hingga T[idx].jurusan.

### II. GUIDED

# 1. Soal study Case
# Source Code
```go
package main

import "fmt"

// Mendeklarasikan tipe data array arrInt dengan panjang 2023
type arrInt [2023]int

// Fungsi untuk mencari indeks elemen terkecil dalam array
func terkecil(tabInt arrInt, n int) int {
    var idx int = 0  // idx menyimpan indeks elemen terkecil
    var j int = 1
    for j < n {
        if tabInt[idx] > tabInt[j] {
            idx = j  // Simpan indeks j jika elemen di indeks j lebih kecil
        }
        j = j + 1
    }
    return idx
}

// Fungsi main untuk menguji fungsi terkecil
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

    // Memanggil fungsi terkecil untuk menemukan indeks elemen terkecil
    idxMin := terkecil(tab, n)

    // Menampilkan nilai dan indeks terkecil
    fmt.Println("Nilai terkecil dalam array adalah:", tab[idxMin], "pada indeks:", idxMin)
}
```

# Screenshot Output

![image](https://github.com/user-attachments/assets/0c21fd95-c573-4a63-bffd-54d82b280e0a)

# Deskripsi Program
Program ini mencari nilai terkecil dalam array dan indeksnya. Pengguna memasukkan jumlah elemen (maksimal 2023) dan nilai-nilainya. Fungsi `terkecil` digunakan untuk menemukan indeks elemen terkecil, yang kemudian ditampilkan bersama nilainya.

# 2. Soal study Case
# Source Code
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

# Screenshot Output

![image](https://github.com/user-attachments/assets/acef8828-03a8-4328-81e9-615f1e6c277b)

# Deskripsi Program
Program ini nyari **IPK tertinggi** dari data mahasiswa. Isi nama, NIM, kelas, jurusan, sama IPK mereka, terus program bakal ngecek siapa yang paling pintar. Hasilnya? Langsung ditampilkan IPK tertingginya. 

### III. UNGUIDED
# 1. Soal study Case
Sebuah program digunakan untuk mendata berat anak kelinci yang akan dijual ke pasar. Program ini menggunakan array dengan kapasitas 1000 untuk menampung data berat anak kelinci yang akan dijual.

Masukan terdiri dari sekumpulan bilangan yang mana bilangan pertama adalah bilangan bulat N yang menyatakan banyaknya anak kelinci yang akan ditimbang beratnya. Selanjutnya N bilangan rill berikutnya adalah berat dari anak kelinci yang akan dijual.

Keluaran terdirfi dari dua buah bilangan rill yang menyatakan berat kelinci terkecil dan terbesar

# Source Code
```go
// Meutya Azzahra Efendi
// 2311102166
// IF-11-06

package main

import "fmt"

func main() {
	// Kapasitas array
	const kapasitas = 1000
	var beratKelinci [kapasitas]float64

	// Input jumlah kelinci
	var N int
	fmt.Print("Masukkan jumlah anak kelinci: ")
	fmt.Scan(&N)

	// Validasi jumlah kelinci
	if N < 1 || N > kapasitas {
		fmt.Println("Jumlah kelinci harus antara 1 dan 1000.")
		return
	}

	// Input berat kelinci
	for i := 0; i < N; i++ {
		fmt.Printf("Masukkan berat kelinci ke-%d: ", i+1)
		fmt.Scan(&beratKelinci[i])
	}

	// Mencari berat terkecil dan terbesar
	beratTerkecil := beratKelinci[0]
	beratTerbesar := beratKelinci[0]

	for i := 1; i < N; i++ {
		if beratKelinci[i] < beratTerkecil {
			beratTerkecil = beratKelinci[i]
		}
		if beratKelinci[i] > beratTerbesar {
			beratTerbesar = beratKelinci[i]
		}
	}

	fmt.Printf("Berat kelinci terkecil: %.2f\n", beratTerkecil)
	fmt.Printf("Berat kelinci terbesar: %.2f\n", beratTerbesar)
}
```

# Screenshot Output

![image](https://github.com/user-attachments/assets/63db4dde-0114-4d5d-aba5-a8ce0543325d)

# Deskripsi Program
Program ini menentukan berat terkecil dan terbesar dari sejumlah anak kelinci. Pengguna memasukkan jumlah kelinci (maksimal 1000) dan berat masing-masing kelinci. Program memvalidasi input, mencari berat terkecil dan terbesar melalui perulangan, lalu mencetak hasilnya.

# 2. Soal study Case
Sebuah program digunakan untuk tarif ikan yang akan dijual ke pasar. Program ini menggunakan array dengan kapasitas 1000 untuk menampung data berat ikan yang akan dijual.

Masukan terdiri dari dua baris, yang mana baris pertama terdiri dari dua bilangan bulat x dan y. Bilangan x menyatakan banyaknya ikan yang akan dijual, sedangkan y adalah banyaknya ikan yang akan dimasukkan ke dalam wadah. Baris kedua terdiri dari sejumlah x bilangan rill yang menyatakan banyaknya ikan yang akan dijual.

Keluaran terdiri dari dua baris. Baris pertama adalah kumpulan bilangan rill yang menyatakan total berat ikan di setiap wadah (jumlah wadah tergantung pada nilai x dan y, urutan ikan yang dimasukkan ke dalam wadah sesuai urutan pada masukan baris ke-2). Baris kedua adalah sebuah bilangan rill yang menyatakan berat rata-rata aikan di setiap wadah.
# Source Code
```go
// Meutya Azzahra Efendi
// 2311102166
// IF-11-06

package main

import "fmt"

func main() {
	// Kapasitas array
	const kapasitas = 1000
	var beratIkan [kapasitas]float64

	// Input jumlah ikan dan wadah
	var x, y int
	fmt.Print("Masukkan jumlah ikan (x) dan jumlah ikan per wadah (y): ")
	fmt.Scan(&x, &y)

	// Validasi jumlah ikan
	if x < 1 || x > kapasitas || y < 1 {
		fmt.Println("Jumlah ikan harus antara 1 dan 1000, dan jumlah ikan per wadah harus lebih dari 0.")
		return
	}

	// Input berat ikan
	fmt.Println("Masukkan berat ikan:")
	for i := 0; i < x; i++ {
		fmt.Printf("Berat ikan ke-%d: ", i+1)
		fmt.Scan(&beratIkan[i])
	}

	// Menghitung total berat ikan di setiap wadah
	var totalBeratPerWadah []float64
	var jumlahIkanDiWadah int
	var totalBerat float64

	for i := 0; i < x; i++ {
		totalBerat += beratIkan[i]
		jumlahIkanDiWadah++

		// Jika sudah mencapai jumlah ikan per wadah, simpan total berat dan reset
		if jumlahIkanDiWadah == y || i == x-1 {
			totalBeratPerWadah = append(totalBeratPerWadah, totalBerat)
			totalBerat = 0
			jumlahIkanDiWadah = 0
		}
	}

	// Output total berat ikan di setiap wadah
	fmt.Println("Total berat ikan di setiap wadah:")
	for _, total := range totalBeratPerWadah {
		fmt.Printf("%.2f ", total)
	}
	fmt.Println()

	// Menghitung dan menampilkan berat rata-rata ikan di setiap wadah
	var totalRataRata float64
	for _, total := range totalBeratPerWadah {
		totalRataRata += total
	}

	// Menghitung jumlah wadah yang digunakan
	jumlahWadah := len(totalBeratPerWadah)

	if jumlahWadah > 0 {
		rataRata := totalRataRata / float64(jumlahWadah)
		fmt.Printf("Berat rata-rata ikan di setiap wadah: %.2f\n", rataRata)
	} else {
		fmt.Println("Tidak ada wadah yang digunakan.")
	}
}
```
# Screenshot Output

![image](https://github.com/user-attachments/assets/201592ce-e16d-4f21-ba68-f446bc065b68)

# Deskripsi Program
Program ini menghitung total berat ikan dalam setiap wadah serta berat rata-rata ikan di setiap wadah. Pengguna memasukkan jumlah ikan, kapasitas ikan per wadah, dan berat masing-masing ikan. Program memvalidasi input, membagi ikan ke wadah sesuai kapasitas, menghitung total berat setiap wadah, dan menampilkan berat rata-rata. Jika jumlah ikan tidak valid, program memberikan pesan kesalahan.

# 3. Soal study Case
Pos Pelayanan Terpadu (Posyandu) sebagai tempat pelayanan kesehatan perlu mencatat data berat balita (dalam kg). Petugas akan memasukkan data tersebut ke dalam arrat. Dari data yang diperoleh akan dicari berat balita terkecil, terbesar, dan reratanya.

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
    /* menghitung dan mengembalikan rata-rata berat balita dalam array */
    // ... isi fungsi di sini
}
```

Perhatikan sesi interaksi pada contoh berikut ini


| **No** | **Berat Balita (kg)** |  
|--------|-----------------------|  
| 1      | 5.3                   |  
| 2      | 6.2                   |  
| 3      | 4.1                   |  
| 4      | 9.9                   |  

**Hasil Perhitungan:**  
- **Berat Minimum:** 4.10 kg  
- **Berat Maksimum:** 9.90 kg  
- **Rata-rata Berat Balita:** 6.38 kg  

# Source Code
```go
// Meutya Azzahra Efendi
// 2311102166
// IF-11-06

package main

import "fmt"

type arrBalita [100]float64

// Fungsi untuk menghitung berat minimum dan maksimum
func hitungMinMax(arrBerat arrBalita, bMin, bMax *float64, n int) {
	*bMin = arrBerat[0]
	*bMax = arrBerat[0]

	for i := 1; i < n; i++ {
		if arrBerat[i] < *bMin {
			*bMin = arrBerat[i]
		}
		if arrBerat[i] > *bMax {
			*bMax = arrBerat[i]
		}
	}
}

// Fungsi untuk menghitung rata-rata berat balita
func rerata(arrBerat arrBalita, n int) float64 {
	total := 0.0
	for i := 0; i < n; i++ {
		total += arrBerat[i]
	}
	return total / float64(n)
}

func main() {
	var beratBalita arrBalita
	var n int

	// Input banyak data berat balita
	fmt.Print("Masukan banyak data berat balita: ")
	fmt.Scan(&n)

	// Validasi jumlah data
	if n < 1 || n > 100 {
		fmt.Println("Jumlah data harus antara 1 dan 100.")
		return
	}

	// Input berat balita
	for i := 0; i < n; i++ {
		fmt.Printf("Masukan berat balita ke-%d: ", i+1)
		fmt.Scan(&beratBalita[i])
	}

	// Variabel untuk menyimpan berat minimum dan maksimum
	var beratMin, beratMax float64

	// Hitung berat minimum dan maksimum
	hitungMinMax(beratBalita, &beratMin, &beratMax, n)

	// Hitung rata-rata berat balita
	rataRata := rerata(beratBalita, n)

	// Output hasil
	fmt.Printf("Berat balita minimum: %.2f kg\n", beratMin)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", beratMax)
	fmt.Printf("Rata-rata berat balita: %.2f kg\n", rataRata)
}
```
# Screenshot Output

![image](https://github.com/user-attachments/assets/167c6004-8894-4337-8457-7dcdbefc9da7)

# Deskripsi Program
Program ini menghitung berat minimum, maksimum, dan rata-rata dari data berat balita yang dimasukkan.
