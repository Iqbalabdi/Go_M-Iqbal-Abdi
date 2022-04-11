# (24) Clean and Hexagonal Architecture

- [Summary](#Summary)
- [Praktikum](#Praktikum)

## Summary

## Praktikum
### 1. Rewrite
Pada task ini diberikan perintah untuk melakukan rewrite dari kodingan pada `https://github.com/hadihammurabi/belajar-go-echo` dengan menerapkan prinsip clean architecture.

Credit : `https://github.com/bxcodec/go-clean-arch`  

Untuk menerapkan prinsip clean architecture, saya menggunakan referensi diagram berikut :  
![diagram](https://raw.githubusercontent.com/bxcodec/go-clean-arch/master/clean-arch.png)  

Dengan menggunakan diagram tersebut, saya dapat membuat struktur project saya sebagai berikut :  
![structure](./screenshots/structure.jpg) 

- model/user.go
  [code](./screenshots/user_model.jpg) 
  file ini berisi entity user dasn entity ini akan dipakai oleh dua interface yaitu `userRepository` dan `userUseCase`



