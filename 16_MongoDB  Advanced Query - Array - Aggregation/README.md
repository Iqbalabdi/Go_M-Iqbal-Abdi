# (16) MongoDB : Advanced Query – Array – Aggregation

- [Summary](#Summary)
- [Praktikum](#Praktikum)

## Summary

## Praktikum
### 1. Gabungkan (menampilkan) data buku dari author id 1 dan author id 2.
```js
db.books.aggregate( {$match: {authorID: {$in: [1,2]}} } )
```
Hasil :  
![hasil](./screenshots/1.jpg)  

### 2. Tampilkan daftar buku dan harga author id 1.
```js
db.books.aggregate([
    {$match : {authorID : 1}},
    {$project: {title : 1, price: 1}}
])
```

Hasil :  
![hasil](./screenshots/2.jpg)  

### 3. Tampilan total jumlah halaman buku author id 2.
```js
db.books.aggregate([
    {$group: {_id: "$authorID", totalPages: {$sum: "$stats.page"}}},
    {$match : {_id : 2}}
])
```

Hasil :  
![hasil](./screenshots/3.jpg)  

### 4. Tampilkan semua field books and authors terkait.
```js
 db.books.aggregate([
     {
        $lookup : {
            from : "authors",
            localField: "authorID",
            foreignField: "_id",
            as: "author"
        }
     }
 ])
```
Hasil :  
![hasil](./screenshots/4.jpg)  

### 5. Tampilkan semua field books, authors, dan publishers terkait.
```js
 db.books.aggregate([
     {
        $lookup : {
            from : "authors",
            localField: "authorID",
            foreignField: "_id",
            as: "author"
        }
     },
     {
        $lookup : {
            from : "publishers",
            localField: "publisherID",
            foreignField: "_id",
            as: "publishers"
        }
     }
 ])
```

Hasil :  
![hasil](./screenshots/5.jpg)  

### 7. Tampilkan semua field books, authors, dan publishers terkait.
```js
db.books.aggregate([
     {
        $lookup : {
            from : "authors",
            localField: "authorID",
            foreignField: "_id",
            as: "author"
        }
     },
     {
        $lookup : {
            from : "publishers",
            localField: "publisherID",
            foreignField: "_id",
            as: "publishers"
        }
     },
     {
       $project : {
         title: 1,
         price : 1,
         author: {
           $concat: ["$authors.firstName"," ", "$authors.lastName"]
         },
         publisher: "$publishers.publisherName"
       }
     },
     {
       $sort: {
         price: -1
       }
     }
 ])
```
Hasil :  
 
### 8. Tampilkan data nama buku harga dan publisher, kemudian tampilkan hanya data ke 3 dan ke 4.
```js
 db.books.aggregate([
     {
        $lookup : {
            from : "authors",
            localField: "authorID",
            foreignField: "_id",
            as: "author"
        }
     },
     {
        $lookup : {
            from : "publishers",
            localField: "publisherID",
            foreignField: "_id",
            as: "publishers"
        }
     },
     {
       $project : {
         _id : 1,
         title: 1,
         price : 1,
         publisher: "$publishers.publisherName"
       }
     },
     {
       $skip: 2
     },
     {
       $limit: 2
     }
 ])
```
Hasil :  
![hasil](./screenshots/8.jpg)   
