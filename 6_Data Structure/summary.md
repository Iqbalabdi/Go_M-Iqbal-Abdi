# 6. Data Structure

## Resume
Dalam materi ini, mempelajari:
1. Array
2. Slice
3. Map

### Array
Array adalah kumpulan data bertipe sama, yang disimpan dalam sebuah variabel. Array memiliki kapasitas yang nilainya ditentukan pada saat pembuatan, menjadikan elemen/data yang disimpan di array tersebut jumlahnya tidak boleh melebihi yang sudah dialokasikan. Default nilai tiap elemen array pada awalnya tergantung dari tipe datanya. Setiap elemen array memiliki indeks berupa angka yang merepresentasikan posisi urutan elemen tersebut. Indeks array dimulai dari 0.

### Slice
Slice adalah reference elemen array. Slice bisa dibuat, atau bisa juga dihasilkan dari manipulasi sebuah array ataupun slice lainnya. Karena merupakan data reference, menjadikan perubahan data di tiap elemen slice akan berdampak pada slice lain yang memiliki alamat memori yang sama.

### Map
Map adalah tipe data asosiatif yang ada di Go, berbentuk key-value pair. Untuk setiap data (atau value) yang disimpan, disiapkan juga key-nya. Key harus unik, karena digunakan sebagai penanda (atau identifier) untuk pengaksesan value yang bersangkutan.

Map terlihat seperti `slice`, hanya saja indeks yang digunakan untuk pengaksesan bisa ditentukan sendiri tipe-nya (indeks tersebut adalah key).

## Task
### 1. Array Merge
Pada task ini, dirancang sebuah program untuk menggabungkan 2 array string. Jika ada elemen yang sama, maka salah satu saja yang ditampilkan

source code :  
![array-merge](./screenshots/1_array_merge_code.jpg) 

output :  
![array-merge](./screenshots/1_array_merge_hasil.jpg) 

penjelasan :  
di awal program ini menggabungkan 2 array `arrayA` dan `arrayB`, lalu dibuat sebuah variabel `check` sebagai map dan `result` untuk menyimpan elemen yang akan dicek. Selanjutnya program ini akan melooping pada array yang telah digabung dan menginput elemen `appendArray` sebagai `key` dan `value = 1` ke variabel `check`. Karena disini menggunakan map, sehingga saat looping tersebut elemen yang duplicate tidak akan masuk ke `check`. Lalu map dari `check` akan diappend ke `result`

### 2. Angka Muncul Sekali
Pada task ini, dirancang sebuah program untuk menemukan elemen apa saja yang muncul nya hanya sekali dari sebuah array

source code :  
![munculsekali](./screenshots/2_munculSekali_code.jpg)

output :  
![munculsekali](./screenshots/2_munculSekali_hasil.jpg)

penjelasan :   
di awal program ini akan melakukan split pada string `angka` agar menjadi sebuah array. Lalu array tersebut akan dikonversi ke int dengan `strconv.Atoi`. Kemudian dilakukan looping untuk men-check apakah ada nilai yang duplicate dengan variabel `counter`. Jika `counter = 1` maka itu artinya variabel tersebut hanya muncul sekali. Lalu direturn nilai tersebut

### 3. Pair With Target Sum
Pada task ini, dirancang sebuah program untuk menemukan index keberapa yang menghasilkan nilain yang sama dengan target jika elemen pada index tersebut ditambah

source code :  
![pairTarget](./screenshots/3_pair_target_code.jpg)

output :  
![pairTarget](./screenshots/3_pair_target_hasil.jpg)

penjelasan :   
Program ini menggunakan variabel `low` dan `high` untuk menentukan index dari setiap elemen yang akan ditambahkan. Jika element yang ditambahkan nilainya belum sama dengan target maka value `low` akan ditambah dan value `high` akan dikurang. Hal ini terus berlanjut hingga nilai dari `low` dan `high` dapat menentukan index yang elemen nya jika ditambahkan nilai nya == target. Karena program ini menggunakan satu `for loop` sehingga compleksitas dari program ini adalah O(n)
