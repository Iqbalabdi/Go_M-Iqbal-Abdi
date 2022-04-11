# (23) Unit Testing

- [Summary](#Summary)
- [Praktikum](#Praktikum)

## Summary
### Apa itu Software Testing
Software Testing merupakan proses untuk menganalisis software item / unit untuk mendeteksi perbedaan antara kondisi yang sebenarnya dan kondisi yang diinginkan dan untuk mengevaluasi features dari unit

### Level of Testing
dari atas ke bawah :
UI -> Integration -> Unit

### Structure
ada 2 pattern yang biasa digunakan untuk melakukan testing pada unit / software item
- __Centralize__ your test file inside tests folder
- Save test file ___together___ with production file 

### Mocking
Mocking adalah sebuah sebuah cara untuk membuat sebuah ___fake object___ untuk mensimulasikan perilaku dari ___real object___

### Testing in golang
Cara untuk melakukan testing di golang dengan syntax :
```
go test -cover
go test -coverprofile=cover.out && go tool cover -html=cover.out
```

## Praktikum
### 1. Simple Unit Testing
Pada task ini diberikan perintah untuk membuat sebuah function calculator dan melakukan testing terhadap function tersebut. Lalu code coverage dari testing itu harus mencapai 100%

Source Code Testing :  
[source-code](./praktikum/calculator/calculate_test.go)

Berikut merupakan coverage dari testing function calculator : 
![hasil](./screenshots/calculator_test.jpg)  

### 2. RESTful API Testing
Pada task ini diberikan perintah untuk mengimplementasikan unit testing dari seluruh endpoint dari project sebelumnya. Dan coverage yang harus dicapai setidaknya >= 80%

Source Code Testing :  
- User Controller   
  [source-code](./praktikum/RESTful-API/controllers/UserController_test.go)
- Book Controller  
  [source-code](./praktikum/RESTful-API/controllers/BookController_test.go)
  
 Berikut merupakan hasil dari coverage report :
 - HTML file  
  ![hasil](./screenshots/coverage_report.jpg)  
 - CLI  
  ![hasil](./screenshots/coverage_report_2.jpg)
