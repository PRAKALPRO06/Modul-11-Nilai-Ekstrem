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

```

#### Screenshoot Output

#### Deskripsi Program


### 2. Penentuan IPK Tertinggi Mahasiswa dalam Suatu Data Kelas

#### Source Code

```go

```

#### Screenshoot Output

#### Deskripsi Program


## Unguided

### Study Case

Sebuah program digunakan untuk mendata berat anak kelinci yang akan dijual ke pasar. Program ini menggunakan array dengan kapasitas 1000 untuk menampung data berat anak kelinci yang akan dijual.

**Masukan:**

Terdiri dari sekumpulan bilangan, yang mana bilangan pertama adalah bilangan bulat N yang menyatakan banyaknya anak kelinci yang akan ditimbang beratnya. Selanjutnya N bilangan riil berikutnya adalah berat dari anak kelinci yang akan dijual.

**Keluaran:**

Terdiri dari dua buah bilangan riil yang menyatakan berat kelinci terkecil dan terbesar.

#### Source Code

```go


```

#### Screenshoot Output

#### Alur Program


#### Deskripsi Program


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

#### Alur Program


#### Deskripsi Program