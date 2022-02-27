# 5. Time Complexity & Space Complexity

## Resume
Dalam materi ini, mempelajari:
1. Time Complexity
2. Space Complexity

### Time Complexity
Time Complexity menyatakan seberapa lama suatu algoritma berjalan saat runtime dari sebuah program. Bisa terlihat sebagai beapa banyak operasi dominan harus melakukan eksekusi nya. Complexity ini dinyatakan dengan Big-O notation. Beberapa Big-O notation yang biasa ditemui adalah :  
1. O(1) : constat
2. O(n) : linear
3. O(log n) : logaritmic
4. O(n^2) : quadratic

Berikut merupakan grafik dari perbandingan beberapa Big-O notation :  
![image](https://user-images.githubusercontent.com/75016595/155878783-4ea77ca9-7451-4809-a04c-72b13423cab4.png)  

### Space Complexity
Space complexity menyatakan seberapa besar alokasi memory yang diperlukan untuk menjalankan program. Pada space complexity, notasi yang digunakan sama dengan time complexity yaitu menggunakan Big-O notation.

## Task
### 1. Menentukan bilangan prima
Pada task ini, dirancang sebuah program untuk menentukan sebuah bilangan prima dengan komplexitas lebih cepat dari O(n) atau O(n/2)

source code :  
![prime-code](./screenshots/1_prime_number_code.jpg) 

output :  
![prime-hasil](./screenshots/1_prime_number_hasil.jpg) 

penjelasan :  
time complexity pada program ini adalah O(sqrt(N)). Karena program ini mengunakan for loop dengan melooping secara kuadratic sebelum nilai loop tersebut melebihi nilai input. Karena travel yang dilakukan pada loop secara quadratic, sehingga waktu yang dibutuhkan untuk menentukan prime number pada program ini berkurang hingga sqrt(N)

### 2. Menentukan nilai exponential
Pada task ini, dirancang sebuah program untuk menghitung nilai exponential dari sebuah bilangan dengan kompleksitas lebih cepat dari O(n)

source code :  
![exponen-code](./screenshots/2_fast_exponen_code.jpg)

output :  
![exponen-hasil](./screenshots/2_fast_exponen_hasil.jpg) 

penjelasan :   
time complexity pada program ini adalah O(log N). Karena program ini melakukan rekursi dengan selalu membagi 2 nilai exponen nya. sehingga travel time yang dilakukan pada rekursi nya meningkat secara log(N) seperti binary search. N disini adalah variabel `expo` pada source code. Karena itu waktu yang dibutuhkan untuk menghitung nilai exponen dari sebuah bilangan secepat log(N)

